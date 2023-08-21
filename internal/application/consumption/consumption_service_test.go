package consumption

import (
	"context"
	"testing"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func setupConsumptionService(t *testing.T) *consumptionServiceMock {
	repoMock := mocks.NewConsumptionRepository(t)

	return &consumptionServiceMock{
		repo: repoMock,
		srv:  NewConsumptionService(repoMock),
	}
}

type consumptionServiceMock struct {
	repo *mocks.ConsumptionRepository
	srv  consumption.ConsumptionService
}

func TestGetGroupedByMetersIDs_ApiError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)

	csm := setupConsumptionService(t)
	csm.repo.On("GetByMetersIDsAndDateRange", ctx, metersIDs, startDate, endDate).Return(nil, assert.AnError)

	groupedConsumption, err := csm.srv.GetGroupedByMetersIDs(ctx, metersIDs, startDate, endDate)

	assert.Nil(t, groupedConsumption)
	assert.Error(t, err)
}

func TestGetGroupedByMetersIDs_ApiOK(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)
	consumptionReponse := []consumption.Consumption{
		consumption.NewConsumption("", 1, 0, 0, 0, 0, startDate.AddDate(0, 0, 1)),
		consumption.NewConsumption("", 1, 0, 0, 0, 0, startDate.AddDate(0, 0, 3)),
		consumption.NewConsumption("", 2, 0, 0, 0, 0, startDate.AddDate(0, 0, 1)),
	}

	csm := setupConsumptionService(t)
	csm.repo.On("GetByMetersIDsAndDateRange", ctx, metersIDs, startDate, endDate).Return(consumptionReponse, nil)

	groupedConsumption, err := csm.srv.GetGroupedByMetersIDs(ctx, metersIDs, startDate, endDate)

	assert.NotNil(t, groupedConsumption)
	assert.NoError(t, err)
	assert.Len(t, groupedConsumption, 2)
	assert.Len(t, groupedConsumption[1], 2)
	assert.Len(t, groupedConsumption[2], 1)
}
