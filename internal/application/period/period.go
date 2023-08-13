package period

import (
	"time"
)

func NewPeriod(startDate, endDate time.Time) *Period {
	return &Period{
		startDate,
		endDate,
	}
}

type Period struct {
	startDate time.Time
	endDate   time.Time
}

func (p *Period) StartDate() time.Time {
	return p.startDate
}

func (p *Period) EndDate() time.Time {
	return p.endDate
}
