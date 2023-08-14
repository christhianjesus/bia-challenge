package consumption

import (
	"net/http"
	"time"

	appConsumption "github.com/christhianjesus/bia-challenge/internal/application/consumption"
	"github.com/christhianjesus/bia-challenge/internal/domain/address"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

type consumptionHandler struct {
	cs  consumption.ConsumptionService
	cps appConsumption.ConsumptionPeriodsService
	as  address.AddressService
}

type consumptionSearchParams struct {
	MetersIDs  []int  `query:"meters_ids"`
	StartDate  string `query:"start_date"`
	EndDate    string `query:"end_date"`
	KindPeriod string `query:"kind_period"`
}

type accumulatedConsumption struct {
	MeterID            int       `json:"meter_id"`
	Address            string    `json:"address"`
	Active             []float64 `json:"active"`
	ReactiveInductive  []float64 `json:"reactive_inductive"`
	ReactiveCapacitive []float64 `json:"reactive_capacitive"`
	Exported           []float64 `json:"exported"`
}

func NewConsumptionHandler(cs consumption.ConsumptionService, cps appConsumption.ConsumptionPeriodsService, as address.AddressService) infrastructure.Handler {
	return &consumptionHandler{cs, cps, as}
}

func (h *consumptionHandler) RegisterRoutes(router *echo.Group) {
	router.GET("/consumption", h.GetAccumulatedConsumption)
}

func (h *consumptionHandler) GetAccumulatedConsumption(c echo.Context) error {
	ctx := c.Request().Context()

	params := &consumptionSearchParams{}
	if err := c.Bind(params); err != nil {
		return err
	}

	startDate, err := time.Parse("2006-01-02", params.StartDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid start_date format")
	}

	endDate, err := time.Parse("2006-01-02", params.EndDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid end_date format")
	}

	periods, err := h.cps.GetPeriods(startDate, endDate, params.KindPeriod)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	addresses, err := h.as.GetByMetersIDs(ctx, params.MetersIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	groupedConsumptions, err := h.cs.GetGroupedByMetersIDs(ctx, params.MetersIDs, startDate, endDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	dataGraph := make([]*accumulatedConsumption, 0, len(params.MetersIDs))

	for _, meterID := range params.MetersIDs {
		consumptions := groupedConsumptions[meterID]
		consumptionPeriods := h.cps.GetConsumptionPeriods(consumptions, periods)
		active, rInductive, rCapacitive, exported := consumptionPeriods.SummarizeValues()
		dataGraph = append(dataGraph, &accumulatedConsumption{
			MeterID:            meterID,
			Address:            addresses[meterID],
			Active:             active,
			ReactiveInductive:  rInductive,
			ReactiveCapacitive: rCapacitive,
			Exported:           exported,
		})
	}

	periodsRanges := make([]string, 0, len(periods))
	for _, period := range periods {
		periodsRanges = append(periodsRanges, period.Describe())
	}

	response := map[string]interface{}{
		"period":     periodsRanges,
		"data_graph": dataGraph,
	}

	return c.JSON(http.StatusOK, response)
}
