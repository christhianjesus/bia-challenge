package consumption

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/periodstrategy/frequency"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumptionperiods"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
	"github.com/christhianjesus/bia-challenge/internal/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupConsumptionHandler(t *testing.T, metersIDs, starDate, endDate, KindPeriod string) *consumptionHandlerMock {
	q := make(url.Values, 4)
	q.Set("meters_ids", metersIDs)
	q.Set("start_date", starDate)
	q.Set("end_date", endDate)
	q.Set("kind_period", KindPeriod)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/consumption?"+q.Encode(), nil)
	rec := httptest.NewRecorder()

	return &consumptionHandlerMock{
		rec:  rec,
		ctx:  e.NewContext(req, rec),
		asm:  mocks.NewAddressService(t),
		csm:  mocks.NewConsumptionService(t),
		cpsm: mocks.NewConsumptionPeriodsService(t),
	}
}

type consumptionHandlerMock struct {
	rec  *httptest.ResponseRecorder
	ctx  echo.Context
	asm  *mocks.AddressService
	csm  *mocks.ConsumptionService
	cpsm *mocks.ConsumptionPeriodsService
}

func TestGetAccumulatedConsumption_BindError(t *testing.T) {
	ch := setupConsumptionHandler(t, "a", "2023-06-01", "2023-06-10", "daily")

	h := NewConsumptionHandler(ch.csm, ch.cpsm, ch.asm).(*consumptionHandler)
	err := h.GetAccumulatedConsumption(ch.ctx)
	he := err.(*echo.HTTPError)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, he.Code)
	assert.Equal(t, "strconv.ParseInt: parsing \"a\": invalid syntax", he.Message)
}

func TestGetAccumulatedConsumption_StarDateError(t *testing.T) {
	ch := setupConsumptionHandler(t, "1", "20230601", "2023-06-10", "daily")

	h := NewConsumptionHandler(ch.csm, ch.cpsm, ch.asm).(*consumptionHandler)
	err := h.GetAccumulatedConsumption(ch.ctx)
	he := err.(*echo.HTTPError)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, he.Code)
	assert.Equal(t, "Invalid start_date format", he.Message)
}

func TestGetAccumulatedConsumption_EndDateError(t *testing.T) {
	ch := setupConsumptionHandler(t, "1", "2023-06-01", "20230610", "daily")

	h := NewConsumptionHandler(ch.csm, ch.cpsm, ch.asm).(*consumptionHandler)
	err := h.GetAccumulatedConsumption(ch.ctx)
	he := err.(*echo.HTTPError)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, he.Code)
	assert.Equal(t, "Invalid end_date format", he.Message)
}

func TestGetAccumulatedConsumption_PeriodKindError(t *testing.T) {
	ch := setupConsumptionHandler(t, "1", "2023-06-01", "2023-06-10", "daily")

	startDate, _ := time.Parse("2006-01-02", "2023-06-01")
	endDate, _ := time.Parse("2006-01-02", "2023-06-10")
	ch.cpsm.On("GetPeriods", startDate, endDate, "daily").Return(nil, assert.AnError)

	h := NewConsumptionHandler(ch.csm, ch.cpsm, ch.asm).(*consumptionHandler)
	err := h.GetAccumulatedConsumption(ch.ctx)
	he := err.(*echo.HTTPError)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, he.Code)
	assert.Equal(t, "assert.AnError general error for testing", he.Message)
}

func TestGetAccumulatedConsumption_AddressError(t *testing.T) {
	ch := setupConsumptionHandler(t, "1", "2023-06-01", "2023-06-03", "daily")

	startDate, _ := time.Parse("2006-01-02", "2023-06-01")
	endDate, _ := time.Parse("2006-01-02", "2023-06-03")
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, time.Month(1), 2, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2021, time.Month(1), 3, 0, 0, 0, 0, time.UTC)
	ch.cpsm.On("GetPeriods", startDate, endDate, "daily").Return([]period.Period{
		frequency.NewDailyPeriod(t1, t2),
		frequency.NewDailyPeriod(t2, t3),
	}, nil)
	ctx := context.TODO()
	ch.asm.On("GetByMetersIDs", ctx, []int{1}).Return(nil, assert.AnError)

	h := NewConsumptionHandler(ch.csm, ch.cpsm, ch.asm).(*consumptionHandler)
	err := h.GetAccumulatedConsumption(ch.ctx)
	he := err.(*echo.HTTPError)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
	assert.Equal(t, "assert.AnError general error for testing", he.Message)
}

func TestGetAccumulatedConsumption_ConsumptionsError(t *testing.T) {
	ch := setupConsumptionHandler(t, "1", "2023-06-01", "2023-06-03", "daily")

	startDate, _ := time.Parse("2006-01-02", "2023-06-01")
	endDate, _ := time.Parse("2006-01-02", "2023-06-03")
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, time.Month(1), 2, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2021, time.Month(1), 3, 0, 0, 0, 0, time.UTC)
	ch.cpsm.On("GetPeriods", startDate, endDate, "daily").Return([]period.Period{
		frequency.NewDailyPeriod(t1, t2),
		frequency.NewDailyPeriod(t2, t3),
	}, nil)
	ctx := context.TODO()
	ch.asm.On("GetByMetersIDs", ctx, []int{1}).Return(map[int]string{
		1: "Dirección Mock 1",
	}, nil)
	ch.csm.On("GetGroupedByMetersIDs", ctx, []int{1}, startDate, endDate).Return(nil, assert.AnError)

	h := NewConsumptionHandler(ch.csm, ch.cpsm, ch.asm).(*consumptionHandler)
	err := h.GetAccumulatedConsumption(ch.ctx)
	he := err.(*echo.HTTPError)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
	assert.Equal(t, "assert.AnError general error for testing", he.Message)
}

func TestGetAccumulatedConsumption_OK(t *testing.T) {
	ch := setupConsumptionHandler(t, "1", "2023-06-01", "2023-06-03", "daily")

	startDate, _ := time.Parse("2006-01-02", "2023-06-01")
	endDate, _ := time.Parse("2006-01-02", "2023-06-03")
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, time.Month(1), 2, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2021, time.Month(1), 3, 0, 0, 0, 0, time.UTC)
	periods := []period.Period{
		frequency.NewDailyPeriod(t1, t2),
		frequency.NewDailyPeriod(t2, t3),
	}
	ch.cpsm.On("GetPeriods", startDate, endDate, "daily").Return(periods, nil)
	ctx := context.TODO()
	ch.asm.On("GetByMetersIDs", ctx, []int{1}).Return(map[int]string{
		1: "Dirección Mock 1",
	}, nil)
	consumptions := []consumption.Consumption{
		consumption.NewConsumption("", 1, 1, 3, 4, 1, t1),
		consumption.NewConsumption("", 2, 2, 1, 3, 1, t2),
		consumption.NewConsumption("", 2, 2, 4, 4, 1, t3),
	}
	ch.csm.On("GetGroupedByMetersIDs", ctx, []int{1}, startDate, endDate).Return(map[int][]consumption.Consumption{
		1: consumptions,
	}, nil)
	consumptionPeriods := consumptionperiods.ConsumptionPeriods{
		{
			consumption.NewConsumption("", 1, 1, 3, 4, 1, t1),
		}, {
			consumption.NewConsumption("", 2, 2, 1, 3, 1, t2),
		}, {
			consumption.NewConsumption("", 2, 2, 4, 4, 1, t3),
		},
	}
	ch.cpsm.On("GetConsumptionPeriods", consumptions, periods).Return(consumptionPeriods, nil)

	h := NewConsumptionHandler(ch.csm, ch.cpsm, ch.asm).(*consumptionHandler)
	err := h.GetAccumulatedConsumption(ch.ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, ch.rec.Code)
	assert.JSONEq(t, `{"data_graph":[{"meter_id":1,"address":"Dirección Mock 1","active":[1,2,2],"reactive_inductive":[3,1,4],"reactive_capacitive":[4,3,4],"exported":[1,1,1]}],"period":["Jan 1","Jan 2"]}`, ch.rec.Body.String())
}
