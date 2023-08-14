package strategies

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/period/frequency"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

type WeeklyPeriodStrategy struct{}

func (w *WeeklyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []period.Period {
	weeks := int(endDate.Sub(startDate).Hours()/24/7) + 2
	periods := make([]period.Period, 0, weeks)

	daysToSubtract := int(startDate.Weekday()+6) % 7
	currentDate := startDate.AddDate(0, 0, -daysToSubtract)
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 0, 7)
		periods = append(periods, frequency.NewWeeklyPeriod(currentDate, nextDate))
		currentDate = nextDate
	}

	return periods
}
