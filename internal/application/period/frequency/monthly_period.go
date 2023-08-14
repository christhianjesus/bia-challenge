package frequency

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

func NewMonthlyPeriod(startDate, endDate time.Time) period.Period {
	return &monthlyPeriod{basePeriod{
		startDate,
		endDate,
	}}
}

type monthlyPeriod struct {
	basePeriod
}

func (p *monthlyPeriod) Describe() string {
	return p.startDate.Format("Jan 2006")
}
