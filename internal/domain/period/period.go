package period

import "time"

type Period interface {
	StartDate() time.Time
	EndDate() time.Time
	Describe() string
}
