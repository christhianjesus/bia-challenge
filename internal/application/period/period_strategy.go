package period

import "time"

type PeriodStrategy interface {
	GeneratePeriods(startDate, endDate time.Time) []*Period
	GenerateDescriptions(startDate, endDate time.Time) []string
}
