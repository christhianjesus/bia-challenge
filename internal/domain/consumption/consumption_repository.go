package consumption

import (
	"context"
	"time"
)

type ConsumptionRepository interface {
	GetByMetersIDsAndDateRange(ctx context.Context, metersIDs []int, startDate, endDate time.Time) ([]*Consumption, error)
}
