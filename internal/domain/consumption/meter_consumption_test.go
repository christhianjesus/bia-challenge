package consumption

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupConsumptionCollectionTest() []*Consumption {
	t1 := time.Date(2021, time.Month(2), 5, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.Month(1), 23, 0, 0, 0, 0, time.UTC)

	return []*Consumption{
		NewConsumption("", 1, 1, 3, 4, 1, t1),
		NewConsumption("", 2, 2, 1, 3, 1, t2),
		NewConsumption("", 2, 2, 4, 4, 1, t3),
	}
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

func TestAccumulatedValues(t *testing.T) {
	consumptionCollection := setupConsumptionCollectionTest()
	active, reactiveInductive, reactiveCapacitive, exported := accumulatedValues(consumptionCollection)

	assert.Equal(t, active, 5.0)
	assert.Equal(t, reactiveInductive, 8.0)
	assert.Equal(t, reactiveCapacitive, 11.0)
	assert.Equal(t, exported, 3.0)
}

func TestSummarizeValues(t *testing.T) {
	consumptionSummary := [][]*Consumption{
		setupConsumptionCollectionTest(),
		setupConsumptionCollectionTest(),
	}
	active, reactiveInductive, reactiveCapacitive, exported := summarizeValues(consumptionSummary)

	assert.Equal(t, active, []float64{5, 5})
	assert.Equal(t, reactiveInductive, []float64{8, 8})
	assert.Equal(t, reactiveCapacitive, []float64{11, 11})
	assert.Equal(t, exported, []float64{3, 3})
}
