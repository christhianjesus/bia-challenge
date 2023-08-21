package consumption

import (
	"context"
	"testing"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/christhianjesus/bia-challenge/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupConsumptionRepository(t *testing.T) *consumptionRepositoryMock {
	clientMock := mocks.NewSQLClient(t)

	return &consumptionRepositoryMock{
		client: clientMock,
		repo:   NewPostgreSQLConsumptionRepository(clientMock),
	}
}

type consumptionRepositoryMock struct {
	client *mocks.SQLClient
	repo   consumption.ConsumptionRepository
}

func TestGetByMetersIDs_ClientError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)

	crm := setupConsumptionRepository(t)
	crm.client.On("QueryContext", ctx, mock.Anything, mock.Anything, startDate, endDate).Return(nil, assert.AnError)

	consumptions, err := crm.repo.GetByMetersIDsAndDateRange(ctx, metersIDs, startDate, endDate)

	assert.Nil(t, consumptions)
	assert.Error(t, err)
}

func TestGetByMetersIDs_ScanError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)

	crm := setupConsumptionRepository(t)
	rows := mocks.NewSQLRows(t)
	crm.client.On("QueryContext", ctx, mock.Anything, mock.Anything, startDate, endDate).
		Return(rows, nil)
	rows.On("Next").Return(true)
	rows.On("Close").Return(nil)
	rows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(assert.AnError)

	consumptions, err := crm.repo.GetByMetersIDsAndDateRange(ctx, metersIDs, startDate, endDate)

	assert.Nil(t, consumptions)
	assert.Error(t, err)
}

func TestGetByMetersIDs_RowError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)

	crm := setupConsumptionRepository(t)
	rows := mocks.NewSQLRows(t)
	crm.client.On("QueryContext", ctx, mock.Anything, mock.Anything, startDate, endDate).
		Return(rows, nil)
	rows.On("Next").Return(false)
	rows.On("Close").Return(nil)
	rows.On("Err").
		Return(assert.AnError)

	consumptions, err := crm.repo.GetByMetersIDsAndDateRange(ctx, metersIDs, startDate, endDate)

	assert.Nil(t, consumptions)
	assert.Error(t, err)
}

func TestGetByMetersIDs_OK(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}
	startDate := time.Date(2023, time.June, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)

	crm := setupConsumptionRepository(t)
	rows := mocks.NewSQLRows(t)
	crm.client.On("QueryContext", ctx, mock.Anything, mock.Anything, startDate, endDate).
		Return(rows, nil)

	rows.On("Next").Return(true).Twice()
	rows.On("Next").Return(false).Once()
	rows.On("Close").Return(nil)
	rows.On("Err").Return(nil)
	rows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			*args.Get(1).(*int) = 1
		}).Return(nil).Once()
	rows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			*args.Get(1).(*int) = 2
		}).Return(nil).Once()

	consumptions, err := crm.repo.GetByMetersIDsAndDateRange(ctx, metersIDs, startDate, endDate)

	assert.NotNil(t, consumptions)
	assert.NoError(t, err)
	assert.Len(t, consumptions, 2)
	assert.Equal(t, 1, consumptions[0].MeterID())
	assert.Equal(t, 2, consumptions[1].MeterID())
}
