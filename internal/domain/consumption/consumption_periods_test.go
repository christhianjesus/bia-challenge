package consumption

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupConsumptionCollectionTest() ConsumptionPeriod {
	t1 := time.Date(2021, time.Month(2), 5, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.Month(1), 23, 0, 0, 0, 0, time.UTC)

	return []Consumption{
		NewConsumption("", 1, 1, 3, 4, 1, t1),
		NewConsumption("", 2, 2, 1, 3, 1, t2),
		NewConsumption("", 2, 2, 4, 4, 1, t3),
	}
}

func TestTotalValues(t *testing.T) {
	consumptionCollection := setupConsumptionCollectionTest()
	active, rInductive, rCapacitive, exported := consumptionCollection.TotalValues()

	assert.Equal(t, active, 5.0)
	assert.Equal(t, rInductive, 8.0)
	assert.Equal(t, rCapacitive, 11.0)
	assert.Equal(t, exported, 3.0)
}

func TestSummarizeValues(t *testing.T) {
	nestedCollection := ConsumptionPeriods{
		setupConsumptionCollectionTest(),
		setupConsumptionCollectionTest(),
	}
	active, rInductive, rCapacitive, exported := nestedCollection.SummarizeValues()

	assert.Equal(t, active, []float64{5, 5})
	assert.Equal(t, rInductive, []float64{8, 8})
	assert.Equal(t, rCapacitive, []float64{11, 11})
	assert.Equal(t, exported, []float64{3, 3})
}
