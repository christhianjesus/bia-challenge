package strategies

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/period"
)

type MonthlyPeriodStrategy struct{}

func (m *MonthlyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []*period.Period {
	months := int(endDate.Sub(startDate).Hours()/24/28) + 2
	periods := make([]*period.Period, 0, months)

	year, month, _ := startDate.Date()
	currentDate := time.Date(year, month, 1, 0, 0, 0, 0, startDate.Location())
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 1, 0)
		periods = append(periods, period.NewPeriod(currentDate, nextDate))
		currentDate = nextDate
	}

	return periods
}

func (m *MonthlyPeriodStrategy) GenerateDescriptions(startDate, endDate time.Time) []string {
	months := int(endDate.Sub(startDate).Hours()/24/28) + 2
	descriptions := make([]string, 0, months)

	year, month, _ := startDate.Date()
	currentDate := time.Date(year, month, 1, 0, 0, 0, 0, startDate.Location())
	for !currentDate.After(endDate) {
		nextDate := currentDate.AddDate(0, 1, 0)
		descriptions = append(descriptions, currentDate.Format("Jan 2006"))
		currentDate = nextDate
	}

	return descriptions
}
