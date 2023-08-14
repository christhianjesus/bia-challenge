package strategies

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePeriods_MonthlyPeriodStrategy(t *testing.T) {
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.July, 10, 0, 0, 0, 0, time.UTC)
	periods := (&MonthlyPeriodStrategy{}).GeneratePeriods(startDate, endDate)

	assert.NotNil(t, periods)
	assert.Len(t, periods, 2)

	for i, period := range periods {
		assert.Equal(t, 1, period.StartDate().Day())
		assert.Equal(t, 1, period.EndDate().Day())
		assert.Equal(t, time.Month(6+i), period.StartDate().Month())
		assert.Equal(t, time.Month(7+i), period.EndDate().Month())
	}
}
