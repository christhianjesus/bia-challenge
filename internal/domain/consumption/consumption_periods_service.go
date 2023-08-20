package consumption

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

type ConsumptionPeriodsService interface {
	GetConsumptionPeriods(consumptions []Consumption, periods []period.Period) ConsumptionPeriods
	GetPeriods(startDate, endDate time.Time, kindPeriod string) ([]period.Period, error)
}
