package period

import "time"

type PeriodStrategy interface {
	GeneratePeriods(startDate, endDate time.Time) []Period
	GenerateDescriptions(periods []Period) []string
}

type MonthlyPeriodStrategy struct{}

func (m *MonthlyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []Period {
	return nil
}

func (m *MonthlyPeriodStrategy) GenerateDescriptions(periods []Period) []string {
	return nil
}

type WeeklyPeriodStrategy struct{}

func (w *WeeklyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []Period {
	return nil
}

func (w *WeeklyPeriodStrategy) GenerateDescriptions(periods []Period) []string {
	return nil
}

type DailyPeriodStrategy struct{}

func (d *DailyPeriodStrategy) GeneratePeriods(startDate, endDate time.Time) []Period {
	return nil
}

func (d *DailyPeriodStrategy) GenerateDescriptions(periods []Period) []string {
	return nil
}
