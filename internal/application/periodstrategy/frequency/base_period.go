package frequency

import (
	"time"
)

type basePeriod struct {
	startDate time.Time
	endDate   time.Time
}

func (p *basePeriod) StartDate() time.Time {
	return p.startDate
}

func (p *basePeriod) EndDate() time.Time {
	return p.endDate
}
