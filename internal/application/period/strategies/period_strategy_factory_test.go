package strategies

import (
	"testing"

	"github.com/christhianjesus/bia-challenge/internal/application/period"
	"github.com/stretchr/testify/assert"
)

func TestCreateMonthlyPeriodStrategy(t *testing.T) {
	monthlyPeriodStrategy, err := (&PeriodStrategyFactory{}).CreatePeriodStrategy(period.Monthly)

	assert.NoError(t, err)
	assert.IsType(t, &MonthlyPeriodStrategy{}, monthlyPeriodStrategy)
}

func TestCreateWeeklyPeriodStrategy(t *testing.T) {
	weeklyPeriodStrategy, err := (&PeriodStrategyFactory{}).CreatePeriodStrategy(period.Weekly)

	assert.NoError(t, err)
	assert.IsType(t, &WeeklyPeriodStrategy{}, weeklyPeriodStrategy)
}

func TestCreateDailyPeriodStrategy(t *testing.T) {
	dailyPeriodStrategy, err := (&PeriodStrategyFactory{}).CreatePeriodStrategy(period.Daily)

	assert.NoError(t, err)
	assert.IsType(t, &DailyPeriodStrategy{}, dailyPeriodStrategy)
}

func TestCreateInvalidPeriodStrategy(t *testing.T) {
	invalidPeriodStrategy, err := (&PeriodStrategyFactory{}).CreatePeriodStrategy("none")

	assert.Error(t, err)
	assert.Nil(t, invalidPeriodStrategy)
}
