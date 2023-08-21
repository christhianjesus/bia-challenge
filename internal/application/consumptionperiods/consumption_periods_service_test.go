package consumptionperiods

import (
	"testing"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/periodstrategy/frequency"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
	"github.com/stretchr/testify/assert"
)

func setupConsumptionCollectionTest() []consumption.Consumption {
	t0 := time.Date(2020, time.Month(7), 11, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2021, time.Month(2), 5, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.Month(1), 23, 0, 0, 0, 0, time.UTC)

	return []consumption.Consumption{
		consumption.NewConsumption("", 1, 1, 3, 4, 1, t0),
		consumption.NewConsumption("", 1, 1, 3, 4, 1, t1),
		consumption.NewConsumption("", 2, 2, 1, 3, 1, t2),
		consumption.NewConsumption("", 2, 2, 4, 4, 1, t3),
	}
}

func setupPeriodsTest() []period.Period {
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	return []period.Period{
		frequency.NewDailyPeriod(t1, t2),
		frequency.NewDailyPeriod(t2, t3),
	}
}

func TestGetConsumptionPeriods(t *testing.T) {
	periods := setupPeriodsTest()
	consumptionCollection := setupConsumptionCollectionTest()
	groupedConsumption := NewConsumptionPeriodsService().GetConsumptionPeriods(consumptionCollection, periods)

	assert.NotNil(t, groupedConsumption)
	assert.Len(t, groupedConsumption, 2)
	assert.Len(t, groupedConsumption[0], 1)
	assert.Len(t, groupedConsumption[1], 1)
}

func TestGetPeriods(t *testing.T) {
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)
	periods, err := NewConsumptionPeriodsService().GetPeriods(startDate, endDate, "daily")

	assert.NotNil(t, periods)
	assert.NoError(t, err)
	assert.Len(t, periods, 10)

	for i, period := range periods {
		assert.Equal(t, i+1, period.StartDate().Day())
		assert.Equal(t, i+2, period.EndDate().Day())
	}
}

func TestGetPeriodsError(t *testing.T) {
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)
	periods, err := NewConsumptionPeriodsService().GetPeriods(startDate, endDate, "non-valid")

	assert.Nil(t, periods)
	assert.Error(t, err)
}
