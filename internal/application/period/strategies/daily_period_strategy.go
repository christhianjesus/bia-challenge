package strategies

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/period"
)

type DailyPeriodStrategy struct{}

func (d *DailyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []*period.Period {
	days := int(endDate.Sub(startDate).Hours()/24) + 1
	periods := make([]*period.Period, 0, days)

	currentDate := startDate
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 0, 1)
		periods = append(periods, period.NewPeriod(currentDate, nextDate))
		currentDate = nextDate
	}

	return periods
}

func (d *DailyPeriodStrategy) GenerateDescriptions(startDate, endDate time.Time) []string {
	days := int(endDate.Sub(startDate).Hours()/24) + 1
	descriptions := make([]string, 0, days)

	currentDate := startDate
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 0, 1)
		descriptions = append(descriptions, currentDate.Format("Jan 2"))
		currentDate = nextDate
	}

	return descriptions
}
