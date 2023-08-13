package consumption

func NewMeterConsumption(meterID int, address string, periods [][]*Consumption) *MeterConsumption {
	active, reactiveInductive, reactiveCapacitive, exported := summarizeValues(periods)

	return &MeterConsumption{
		meterID,
		address,
		active,
		reactiveInductive,
		reactiveCapacitive,
		exported,
	}
}

type MeterConsumption struct {
	meterID            int
	address            string
	active             []float64
	reactiveInductive  []float64
	reactiveCapacitive []float64
	exported           []float64
}

func (c *MeterConsumption) GenerateSerializableResponse() map[string]interface{} {
	return map[string]interface{}{
		"meter_id":            c.meterID,
		"address":             c.address,
		"active":              c.active,
		"reactive_inductive":  c.reactiveInductive,
		"reactive_capacitive": c.reactiveCapacitive,
		"exported":            c.exported,
	}
}

func accumulatedValues(consumptions []*Consumption) (active, reactiveInductive, reactiveCapacitive, exported float64) {
	for _, consumption := range consumptions {
		active += consumption.activeEnergy
		reactiveInductive += consumption.reactiveEnergy
		reactiveCapacitive += consumption.capacitiveReactive
		exported += consumption.solar
	}

	return
}

func summarizeValues(cs [][]*Consumption) (active, reactiveInductive, reactiveCapacitive, exported []float64) {
	for _, period := range cs {
		a, ri, rc, e := accumulatedValues(period)

		active = append(active, a)
		reactiveInductive = append(reactiveInductive, ri)
		reactiveCapacitive = append(reactiveCapacitive, rc)
		exported = append(exported, e)
	}

	return
}
