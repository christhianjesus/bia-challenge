package consumption

func NewConsumptionResponse(meterID int, address string, periods []ConsumptionCollection) *ConsumptionResponse {
	active, reactiveInductive, reactiveCapacitive, exported := ConsumptionSummary(periods).SummarizeValues()

	return &ConsumptionResponse{
		meterID,
		address,
		active,
		reactiveInductive,
		reactiveCapacitive,
		exported,
	}
}

type ConsumptionResponse struct {
	meterID            int
	address            string
	active             []float64
	reactiveInductive  []float64
	reactiveCapacitive []float64
	exported           []float64
}

func (c *ConsumptionResponse) GenerateSerializableResponse() map[string]interface{} {
	return map[string]interface{}{
		"meter_id":            c.meterID,
		"address":             c.address,
		"active":              c.active,
		"reactive_inductive":  c.reactiveInductive,
		"reactive_capacitive": c.reactiveCapacitive,
		"exported":            c.exported,
	}
}
