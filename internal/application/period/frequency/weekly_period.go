package frequency

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

func NewWeeklyPeriod(startDate, endDate time.Time) period.Period {
	return &weeklyPeriod{basePeriod{
		startDate,
		endDate,
	}}
}

type weeklyPeriod struct {
	basePeriod
}

func (p *weeklyPeriod) Describe() string {
	return p.startDate.Format("Jan 2") + " - " + p.endDate.Format("Jan 2")
}
