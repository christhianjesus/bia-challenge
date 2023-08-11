package consumption

type ConsumptionSummary []ConsumptionCollection

func (cs ConsumptionSummary) SummarizeValues() (active, reactiveInductive, reactiveCapacitive, exported []float64) {
	for _, period := range cs {
		a, ri, rc, e := period.AccumulatedValues()

		active = append(active, a)
		reactiveInductive = append(reactiveInductive, ri)
		reactiveCapacitive = append(reactiveCapacitive, rc)
		exported = append(exported, e)
	}

	return
}
