package consumption

import (
	"net/http"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

type consumptionHandler struct {
	service consumption.ConsumptionService
}

type consumptionSearchParams struct {
	MetersIDs  []int  `query:"meters_ids"`
	StartDate  string `query:"start_date"`
	EndDate    string `query:"end_date"`
	KindPeriod string `query:"kind_period"`
}

func NewConsumptionHandler(service consumption.ConsumptionService) infrastructure.Handler {
	return &consumptionHandler{service}
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

	consumptionPeriods, err := h.service.GetConsumptionPeriods(ctx, startDate, endDate, params.KindPeriod)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	accumulatedConsumption, err := h.service.GetAccumulatedConsumption(ctx, params.MetersIDs, startDate, endDate, params.KindPeriod)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dataGraph := []map[string]interface{}{}
	for _, meterConsumption := range accumulatedConsumption {
		dataGraph = append(dataGraph, meterConsumption.GenerateSerializableResponse())
	}

	response := map[string]interface{}{
		"period":     consumptionPeriods,
		"data_graph": dataGraph,
	}

	return c.JSON(http.StatusOK, response)
}
