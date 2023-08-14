package consumption

import (
	"time"

	appPeriod "github.com/christhianjesus/bia-challenge/internal/application/period"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

type ConsumptionPeriodsService interface {
	GetConsumptionPeriods(consumptions []*consumption.Consumption, periods []period.Period) consumption.ConsumptionPeriods
	GetPeriods(startDate, endDate time.Time, kindPeriod string) ([]period.Period, error)
}

type consumptionPeriodsService struct{}

func NewConsumptionPeriodsService() ConsumptionPeriodsService {
	return &consumptionPeriodsService{}
}

// it is assumed that the consumptions are ordered by date, as well as the periods
func (c *consumptionPeriodsService) GetConsumptionPeriods(consumptions []*consumption.Consumption, periods []period.Period,
) consumption.ConsumptionPeriods {
	groupedConsumptions := make([][]*consumption.Consumption, len(periods))

	// Fast approach to allocate memory using ceil division
	approxSize := (len(consumptions) + len(periods) + 1) / len(periods)
	for i := range periods {
		groupedConsumptions[i] = make([]*consumption.Consumption, 0, approxSize)
	}

	dateIndex := 0
	periodIndex := 0

	for dateIndex < len(consumptions) && periodIndex < len(periods) {
		date := consumptions[dateIndex].Date()
		period := periods[periodIndex]

		if date.Before(period.StartDate()) {
			dateIndex++
		} else if date.After(period.EndDate()) {
			periodIndex++
		} else {
			groupedConsumptions[periodIndex] = append(groupedConsumptions[periodIndex], consumptions[dateIndex])
			dateIndex++
		}
	}

	return consumption.ConsumptionPeriods(groupedConsumptions)
}

func (c *consumptionPeriodsService) GetPeriods(startDate, endDate time.Time, kindPeriod string) ([]period.Period, error) {
	strategy, err := appPeriod.CreatePeriodStrategy(period.KindPeriod(kindPeriod))
	if err != nil {
		return nil, err
	}

	return strategy.GeneratePeriods(startDate, endDate), nil
}
