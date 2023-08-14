package consumption

type ConsumptionPeriod []*Consumption

func (cp ConsumptionPeriod) TotalValues() (active, rInductive, rCapacitive, exported float64) {
	for _, consumption := range cp {
		active += consumption.activeEnergy
		rInductive += consumption.reactiveEnergy
		rCapacitive += consumption.capacitiveReactive
		exported += consumption.solar
	}

	return
}

type ConsumptionPeriods [][]*Consumption

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
