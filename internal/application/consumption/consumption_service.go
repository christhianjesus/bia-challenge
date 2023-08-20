package consumption

import (
	"context"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
)

type consumptionService struct {
	repo consumption.ConsumptionRepository
}

func NewConsumptionService(repo consumption.ConsumptionRepository) consumption.ConsumptionService {
	return &consumptionService{repo}
}

func (cs *consumptionService) GetGroupedByMetersIDs(ctx context.Context, metersIDs []int, startDate, endDate time.Time) (map[int][]consumption.Consumption, error) {
	consumptions, err := cs.repo.GetByMetersIDsAndDateRange(ctx, metersIDs, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Fast approach to allocate memory by counting elements in each bucket
	meterConsumptionsTotal := make(map[int]int, len(metersIDs))
	for _, consumption := range consumptions {
		meterConsumptionsTotal[consumption.MeterID()] += 1
	}

	groupedConsumptions := make(map[int][]consumption.Consumption, len(metersIDs))
	for _, meterID := range metersIDs {
		// preallocate memory needed in each bucket
		groupedConsumptions[meterID] = make([]consumption.Consumption, 0, meterConsumptionsTotal[meterID])
	}

	for _, consumption := range consumptions {
		groupedConsumptions[consumption.MeterID()] = append(groupedConsumptions[consumption.MeterID()], consumption)
	}

	return groupedConsumptions, nil
}
