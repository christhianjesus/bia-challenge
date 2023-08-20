package strategies

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePeriods_DailyPeriodStrategy(t *testing.T) {
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)
	periods := (&DailyPeriodStrategy{}).GeneratePeriods(startDate, endDate)

	assert.NotNil(t, periods)
	assert.Len(t, periods, 10)

	for i, period := range periods {
		assert.Equal(t, i+1, period.StartDate().Day())
		assert.Equal(t, i+2, period.EndDate().Day())
	}
}
