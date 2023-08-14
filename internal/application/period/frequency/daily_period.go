package frequency

import (
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

func NewDailyPeriod(startDate, endDate time.Time) period.Period {
	return &dailyPeriod{basePeriod{
		startDate,
		endDate,
	}}
}

type dailyPeriod struct {
	basePeriod
}

func (p *dailyPeriod) Describe() string {
	return p.startDate.Format("Jan 2")
}
