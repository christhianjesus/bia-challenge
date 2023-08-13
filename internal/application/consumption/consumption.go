package consumption

import (
	"github.com/christhianjesus/bia-challenge/internal/application/period"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
)

func GroupByMetersIDs(consumptions []*consumption.Consumption, metersIDs []int) map[int][]*consumption.Consumption {
	meterConsumptionsCounter := make(map[int]int, len(metersIDs))
	for _, consumption := range consumptions {
		meterConsumptionsCounter[consumption.MeterID()] += 1
	}

	groupedConsumptions := make(map[int][]*consumption.Consumption, len(metersIDs))
	for _, meterID := range metersIDs {
		groupedConsumptions[meterID] = make([]*consumption.Consumption, 0, meterConsumptionsCounter[meterID])
	}

	for _, consumption := range consumptions {
		groupedConsumptions[consumption.MeterID()] = append(groupedConsumptions[consumption.MeterID()], consumption)
	}

	return groupedConsumptions
}

// it is assumed that the consumptions are ordered by date, as well as the periods
func GroupByPeriods(consumptions []*consumption.Consumption, periods []*period.Period) [][]*consumption.Consumption {
	groupedConsumptions := make([][]*consumption.Consumption, len(periods))

	// ceil division
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
		} else if date.After(period.EndDate().AddDate(0, 0, 1)) {
			periodIndex++
		} else {
			groupedConsumptions[periodIndex] = append(groupedConsumptions[periodIndex], consumptions[dateIndex])
			dateIndex++
		}
	}

	return groupedConsumptions
}
