package consumption

import (
	"context"
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

func (uc *ConsumptionService) GetAccumulatedConsumption(ctx context.Context, metersIDs []int, startDate, endDate time.Time,
	kindPeriod string) ([]*consumption.MeterConsumption, error) {
	/* 	addresses, err := uc.addressRepo.GetByMetersIDs(ctx, metersIDs)
	   	if err != nil {
	   		fmt.Println(err.Error())
	   	} */

	strategy, err := uc.periodFact.CreatePeriodStrategy(period.PeriodKind(kindPeriod))
	if err != nil {
		return nil, err
	}
	periods := strategy.GeneratePeriods(startDate, endDate)

	consumptions, err := uc.consumptionRepo.GetByMetersIDsAndDateRange(ctx, metersIDs, startDate, endDate)
	if err != nil {
		return nil, err
	}
	consumptionsByMeterID := GroupByMetersIDs(consumptions, metersIDs)

	metersConsumption := make([]*consumption.MeterConsumption, 0, len(consumptionsByMeterID))

	for meterID, consumptions := range consumptionsByMeterID {
		consumptionByPeriod := GroupByPeriods(consumptions, periods)
		meterConsumption := consumption.NewMeterConsumption(meterID, "addresses[meterID]", consumptionByPeriod)
		metersConsumption = append(metersConsumption, meterConsumption)
	}

	return metersConsumption, nil
}

func (uc *ConsumptionService) GetConsumptionPeriods(ctx context.Context, startDate, endDate time.Time, kindPeriod string) ([]string, error) {
	strategy, err := uc.periodFact.CreatePeriodStrategy(period.PeriodKind(kindPeriod))
	if err != nil {
		return nil, err
	}

	return strategy.GenerateDescriptions(startDate, endDate), nil
}
