package consumption

import (
	"testing"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/period"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/stretchr/testify/assert"
)

func setupConsumptionCollectionTest() []*consumption.Consumption {
	t1 := time.Date(2021, time.Month(2), 5, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.Month(1), 23, 0, 0, 0, 0, time.UTC)

	return []*consumption.Consumption{
		consumption.NewConsumption("", 1, 1, 3, 4, 1, t1),
		consumption.NewConsumption("", 2, 2, 1, 3, 1, t2),
		consumption.NewConsumption("", 2, 2, 4, 4, 1, t3),
	}
}

func setupPeriodsTest() []*period.Period {
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	return []*period.Period{
		period.NewPeriod(t1, t2),
		period.NewPeriod(t2, t3),
	}
}

func TestGroupByMeterID(t *testing.T) {
	consumptionCollection := setupConsumptionCollectionTest()
	groupedConsumption := GroupByMeterIDs(consumptionCollection, []int{1, 2})

	assert.NotNil(t, groupedConsumption)
	assert.Len(t, groupedConsumption, 2)
	assert.NotNil(t, groupedConsumption[1])
	assert.NotNil(t, groupedConsumption[2])
}

func TestGroupByPeriod(t *testing.T) {
	periods := setupPeriodsTest()
	consumptionCollection := setupConsumptionCollectionTest()
	groupedConsumption := GroupByPeriods(consumptionCollection, periods)

	assert.NotNil(t, groupedConsumption)
	assert.Len(t, groupedConsumption, 2)
	assert.Len(t, groupedConsumption[0], 1)
	assert.Len(t, groupedConsumption[1], 1)
}
