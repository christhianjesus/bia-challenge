package frequency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewWeeklyPeriod_WeeklyPeriod(t *testing.T) {
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, time.Month(1), 2, 0, 0, 0, 0, time.UTC)

	period := NewWeeklyPeriod(t1, t2)

	assert.NotNil(t, period)
	assert.IsType(t, &weeklyPeriod{}, period)
}

func TestDescription_WeeklyPeriod(t *testing.T) {
	days := []time.Time{
		time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.July, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.August, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.September, 4, 0, 0, 0, 0, time.UTC),
	}

	dates := make([]*weeklyPeriod, 0, 4)
	for _, d := range days {
		dates = append(dates, &weeklyPeriod{basePeriod{d, d.AddDate(0, 0, 7)}})
	}

	assert.Equal(t, "Jun 1 - Jun 8", dates[0].Describe())
	assert.Equal(t, "Jul 2 - Jul 9", dates[1].Describe())
	assert.Equal(t, "Aug 3 - Aug 10", dates[2].Describe())
	assert.Equal(t, "Sep 4 - Sep 11", dates[3].Describe())
}
