package consumption

import (
	"context"
	"fmt"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/application/period"
	"github.com/christhianjesus/bia-challenge/internal/domain/address"
	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
)

type ConsumptionService struct {
	consumptionRepo consumption.ConsumptionRepository
	addressRepo     address.AddressRepository
	periodFact      period.PeriodStrategyFactory
}

func NewConsumptionService(cons consumption.ConsumptionRepository, addr address.AddressRepository, psfa period.PeriodStrategyFactory) *ConsumptionService {
	return &ConsumptionService{consumptionRepo: cons, addressRepo: addr, periodFact: psfa}
}

func (uc *ConsumptionService) GetAccumulatedConsumption(ctx context.Context, meterIDs []int, startDate, endDate time.Time,
	period period.PeriodKind) ([]*consumption.MeterConsumption, error) {
	addresses, err := uc.addressRepo.GetByMeterIDs(ctx, meterIDs)
	if err != nil {
		fmt.Println(err.Error())
	}

	strategy, err := uc.periodFact.CreatePeriodStrategy(period)
	if err != nil {
		return nil, err
	}
	periods := strategy.GeneratePeriods(startDate, endDate)

	consumptions, err := uc.consumptionRepo.GetByMeterIDsAndDateRange(ctx, meterIDs, startDate, endDate)
	if err != nil {
		return nil, err
	}
	consumptionsByMeterID := GroupByMeterIDs(consumptions, meterIDs)

	metersConsumption := make([]*consumption.MeterConsumption, 0, len(consumptionsByMeterID))

	for meterID, consumptions := range consumptionsByMeterID {
		consumptionByPeriod := GroupByPeriods(consumptions, periods)
		meterConsumption := consumption.NewMeterConsumption(meterID, addresses[meterID], consumptionByPeriod)
		metersConsumption = append(metersConsumption, meterConsumption)
	}

	return metersConsumption, nil
}

func (uc *ConsumptionService) GetConsumptionPeriods(ctx context.Context, startDate, endDate time.Time, period period.PeriodKind) ([]string, error) {
	strategy, err := uc.periodFact.CreatePeriodStrategy(period)
	if err != nil {
		return nil, err
	}

	return strategy.GenerateDescriptions(startDate, endDate), nil
}
