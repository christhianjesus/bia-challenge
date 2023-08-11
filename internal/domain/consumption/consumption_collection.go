package consumption

import "github.com/christhianjesus/bia-challenge/internal/domain/period"

type ConsumptionCollection []*Consumption

func (cc ConsumptionCollection) GroupByMeterID() map[int]ConsumptionCollection {
	m := make(map[int]ConsumptionCollection)

	for _, consumption := range cc {
		m[consumption.meterID] = append(m[consumption.meterID], consumption)
	}

	return m
}

func (cc ConsumptionCollection) GroupByPeriod(periods []period.Period) []ConsumptionCollection {
	a := make([]ConsumptionCollection, len(periods))

	for _, consumption := range cc {
		for i, period := range periods {
			if period.IsWithin(consumption.date) {
				a[i] = append(a[i], consumption)
				break
			}
		}
	}

	return a
}

func (cc ConsumptionCollection) AccumulatedValues() (active, reactiveInductive, reactiveCapacitive, exported float64) {
	for _, consumption := range cc {
		active += consumption.activeEnergy
		reactiveInductive += consumption.reactiveEnergy
		reactiveCapacitive += consumption.capacitiveReactive
		exported += consumption.solar
	}

	return
}
