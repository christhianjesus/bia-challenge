package consumption

import (
	"context"
	"time"
)

type ConsumptionService interface {
	GetAccumulatedConsumption(ctx context.Context, metersIDs []int, startDate, endDate time.Time, period string) ([]*MeterConsumption, error)
	GetConsumptionPeriods(ctx context.Context, startDate, endDate time.Time, period string) ([]string, error)
}
