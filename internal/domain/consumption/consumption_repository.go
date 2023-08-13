package consumption

import (
	"context"
	"time"
)

type ConsumptionRepository interface {
	GetByMeterIDsAndDateRange(ctx context.Context, meterIDs []int, startDate, endDate time.Time) ([]*Consumption, error)
}
