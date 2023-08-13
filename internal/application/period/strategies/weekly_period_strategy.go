package strategies

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/period"
)

type WeeklyPeriodStrategy struct{}

func (w *WeeklyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []*period.Period {
	weeks := int(endDate.Sub(startDate).Hours()/24/7) + 2
	periods := make([]*period.Period, 0, weeks)

	daysToSubtract := int(startDate.Weekday()+6) % 7
	currentDate := startDate.AddDate(0, 0, -daysToSubtract)
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 0, 7)
		periods = append(periods, period.NewPeriod(currentDate, nextDate))
		currentDate = nextDate
	}

	return periods
}

func (w *WeeklyPeriodStrategy) GenerateDescriptions(startDate, endDate time.Time) []string {
	weeks := int(endDate.Sub(startDate).Hours()/24/7) + 2
	descriptions := make([]string, 0, weeks)

	daysToSubtract := int(startDate.Weekday()+6) % 7
	currentDate := startDate.AddDate(0, 0, -daysToSubtract)
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 0, 7)
		descriptions = append(descriptions, currentDate.Format("Jan 2")+" - "+nextDate.Format("Jan 2"))
		currentDate = nextDate
	}

	return descriptions
}
