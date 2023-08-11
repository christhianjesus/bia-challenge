package consumption

import (
	"context"
	"time"
)

func NewConsumption(id, meterID int, activeEnergy, reactiveEnergy, capacitiveReactive, solar float64, date time.Time) *Consumption {
	return &Consumption{
		id,
		meterID,
		activeEnergy,
		reactiveEnergy,
		capacitiveReactive,
		solar,
		date,
	}
}

type Consumption struct {
	id                 int
	meterID            int
	activeEnergy       float64
	reactiveEnergy     float64
	capacitiveReactive float64
	solar              float64
	date               time.Time
}

type ConsumptionRepository interface {
	GetByMeterIDsAndDateRange(ctx context.Context, meterIDs []int, startDate, endDate time.Time) ([]*Consumption, error)
}

type ConsumptionService interface {
	GetWeeklyConsumptions(ctx context.Context, meterIDs []int, referenceDate time.Time) (*ConsumptionResponse, error)
	GetMonthlyConsumptions(ctx context.Context, meterIDs []int, referenceDate time.Time) (*ConsumptionResponse, error)
	GetDailyConsumptions(ctx context.Context, meterIDs []int, referenceDate time.Time) (*ConsumptionResponse, error)
}
