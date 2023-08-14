package consumption

import (
	"context"
	"time"
)

type ConsumptionService interface {
	GetGroupedByMetersIDs(ctx context.Context, metersIDs []int, startDate, endDate time.Time) (map[int][]*Consumption, error)
}
