package strategies

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/period/frequency"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

type DailyPeriodStrategy struct{}

func (d *DailyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []period.Period {
	days := int(endDate.Sub(startDate).Hours()/24) + 1
	periods := make([]period.Period, 0, days)

	currentDate := startDate
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 0, 1)
		periods = append(periods, frequency.NewDailyPeriod(currentDate, nextDate))
		currentDate = nextDate
	}

	return periods
}
