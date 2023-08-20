package strategies

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePeriods_WeeklyPeriodStrategy(t *testing.T) {
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 26, 0, 0, 0, 0, time.UTC)
	periods := (&WeeklyPeriodStrategy{}).GeneratePeriods(startDate, endDate)

	assert.NotNil(t, periods)
	assert.Len(t, periods, 5)

	expectedValues := map[int]time.Month{
		0: time.May,
		1: time.June,
		2: time.June,
		3: time.June,
		4: time.June,
		5: time.July,
	}

	for i, period := range periods {
		assert.Equal(t, time.Monday, period.StartDate().Weekday())
		assert.Equal(t, time.Monday, period.EndDate().Weekday())
		assert.Equal(t, expectedValues[i], period.StartDate().Month())
		assert.Equal(t, expectedValues[i+1], period.EndDate().Month())
	}
}
