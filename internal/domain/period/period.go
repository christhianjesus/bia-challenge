package period

import "time"

type Period struct {
	startDate time.Time
	endDate   time.Time
}

func (p Period) IsWithin(date time.Time) bool {
	return true
}
