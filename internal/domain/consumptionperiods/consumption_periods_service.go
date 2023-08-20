package consumptionperiods

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

type ConsumptionPeriodsService interface {
	GetConsumptionPeriods(consumptions []consumption.Consumption, periods []period.Period) ConsumptionPeriods
	GetPeriods(startDate, endDate time.Time, kindPeriod string) ([]period.Period, error)
}
