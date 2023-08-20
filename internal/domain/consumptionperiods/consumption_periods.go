package consumptionperiods

import "github.com/christhianjesus/bia-challenge/internal/domain/consumption"

type ConsumptionPeriod []consumption.Consumption

func (cp ConsumptionPeriod) TotalValues() (active, rInductive, rCapacitive, exported float64) {
	for _, consumption := range cp {
		active += consumption.ActiveEnergy()
		rInductive += consumption.ReactiveEnergy()
		rCapacitive += consumption.CapacitiveReactive()
		exported += consumption.Solar()
	}

	return
}

type ConsumptionPeriods [][]consumption.Consumption

func (cp ConsumptionPeriods) SummarizeValues() (active, rInductive, rCapacitive, exported []float64) {
	active = make([]float64, 0, len(cp))
	rInductive = make([]float64, 0, len(cp))
	rCapacitive = make([]float64, 0, len(cp))
	exported = make([]float64, 0, len(cp))

	for _, consumptionPeriod := range cp {
		ac, ri, rc, ex := ConsumptionPeriod(consumptionPeriod).TotalValues()

		active = append(active, ac)
		rInductive = append(rInductive, ri)
		rCapacitive = append(rCapacitive, rc)
		exported = append(exported, ex)
	}

	return
}
