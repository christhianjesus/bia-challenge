package periodstrategy

import (
	"testing"

	"github.com/christhianjesus/bia-challenge/internal/application/periodstrategy/strategies"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
	"github.com/stretchr/testify/assert"
)

func TestCreateMonthlyPeriodStrategy(t *testing.T) {
	monthlyPeriodStrategy, err := CreatePeriodStrategy(period.Monthly)

	assert.NoError(t, err)
	assert.IsType(t, &strategies.MonthlyPeriodStrategy{}, monthlyPeriodStrategy)
}

func TestCreateWeeklyPeriodStrategy(t *testing.T) {
	weeklyPeriodStrategy, err := CreatePeriodStrategy(period.Weekly)

	assert.NoError(t, err)
	assert.IsType(t, &strategies.WeeklyPeriodStrategy{}, weeklyPeriodStrategy)
}

func TestCreateDailyPeriodStrategy(t *testing.T) {
	dailyPeriodStrategy, err := CreatePeriodStrategy(period.Daily)

	assert.NoError(t, err)
	assert.IsType(t, &strategies.DailyPeriodStrategy{}, dailyPeriodStrategy)
}

func TestCreateInvalidPeriodStrategy(t *testing.T) {
	invalidPeriodStrategy, err := CreatePeriodStrategy("none")

	assert.Error(t, err)
	assert.Nil(t, invalidPeriodStrategy)
}
