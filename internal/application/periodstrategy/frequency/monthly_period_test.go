package frequency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMonthlyPeriod_MonthlyPeriod(t *testing.T) {
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, time.Month(1), 2, 0, 0, 0, 0, time.UTC)

	period := NewMonthlyPeriod(t1, t2)

	assert.NotNil(t, period)
	assert.IsType(t, &monthlyPeriod{}, period)
}

func TestDescription_MonthlyPeriod(t *testing.T) {
	days := []time.Time{
		time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.July, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.August, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, time.September, 4, 0, 0, 0, 0, time.UTC),
	}

	dates := make([]*monthlyPeriod, 0, 4)
	for _, d := range days {
		dates = append(dates, &monthlyPeriod{basePeriod{d, d}})
	}

	assert.Equal(t, "Jun 2023", dates[0].Describe())
	assert.Equal(t, "Jul 2023", dates[1].Describe())
	assert.Equal(t, "Aug 2023", dates[2].Describe())
	assert.Equal(t, "Sep 2023", dates[3].Describe())
}
