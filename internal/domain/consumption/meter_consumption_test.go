package consumption

/*
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

// ConsumptionHandler
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

func TestNewMeterConsumption(t *testing.T) {
	meterID := 1
	address := "Mock address"
	consumptionCollection := [][]*Consumption{
		setupConsumptionCollectionTest(),
		setupConsumptionCollectionTest(),
	}

	consumptionResponse := NewMeterConsumption(meterID, address, consumptionCollection)

	assert.NotNil(t, consumptionResponse)
	assert.Equal(t, consumptionResponse.meterID, meterID)
	assert.Equal(t, consumptionResponse.address, address)
	assert.Equal(t, consumptionResponse.active, []float64{5, 5})
	assert.Equal(t, consumptionResponse.reactiveInductive, []float64{8, 8})
	assert.Equal(t, consumptionResponse.reactiveCapacitive, []float64{11, 11})
	assert.Equal(t, consumptionResponse.exported, []float64{3, 3})
}
*/
