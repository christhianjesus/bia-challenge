package periodstrategy

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

type PeriodStrategy interface {
	GeneratePeriods(startDate, endDate time.Time) []period.Period
}
