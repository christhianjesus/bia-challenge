package address

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/christhianjesus/bia-challenge/internal/domain/address"
	"github.com/christhianjesus/bia-challenge/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupAddressRepository(t *testing.T) *addressRepositoryMock {
	clientMock := mocks.NewHTTPClient(t)

	return &addressRepositoryMock{
		client: clientMock,
		repo:   NewMSAddressRepository(clientMock),
	}
}

type addressRepositoryMock struct {
	client *mocks.HTTPClient
	repo   address.AddressRepository
}

func TestGetByMetersIDs_RequestError(t *testing.T) {
	ctx := context.Context(nil)
	metersIDs := []int{1, 2}

	arm := setupAddressRepository(t)

	addresses, err := arm.repo.GetByMetersIDs(ctx, metersIDs)

	assert.Nil(t, addresses)
	assert.Error(t, err)
}

func TestGetByMetersIDs_ClientError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}

	arm := setupAddressRepository(t)
	arm.client.On("Do", mock.Anything).Return(nil, assert.AnError)

	addresses, err := arm.repo.GetByMetersIDs(ctx, metersIDs)

	assert.Nil(t, addresses)
	assert.Error(t, err)
}

func TestGetByMetersIDs_ReadError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}

	arm := setupAddressRepository(t)
	arm.client.On("Do", mock.Anything).
		Return(&http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(&FailRead{}),
		}, nil)

	addresses, err := arm.repo.GetByMetersIDs(ctx, metersIDs)

	assert.Nil(t, addresses)
	assert.Error(t, err)
}

func TestGetByMetersIDs_UnmarshalError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}

	arm := setupAddressRepository(t)
	arm.client.On("Do", mock.Anything).
		Return(&http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(``)),
		}, nil)

	addresses, err := arm.repo.GetByMetersIDs(ctx, metersIDs)

	assert.Nil(t, addresses)
	assert.Error(t, err)
}

func TestGetByMetersIDs_OK(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2, 3}

	arm := setupAddressRepository(t)
	arm.client.On("Do", mock.Anything).
		Return(&http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{"addresses": {"1": "Dirección Mock 1", "2": "Dirección Mock 2", "3": "Dirección Mock 3"}}`)),
		}, nil)

	addresses, err := arm.repo.GetByMetersIDs(ctx, metersIDs)

	assert.NotNil(t, addresses)
	assert.NoError(t, err)

	assert.Len(t, addresses, 3)
	assert.Equal(t, "Dirección Mock 1", addresses[1])
	assert.Equal(t, "Dirección Mock 2", addresses[2])
	assert.Equal(t, "Dirección Mock 3", addresses[3])
}

type FailRead struct{}

func (*FailRead) Read(p []byte) (n int, err error) {
	return 0, errors.New("any error")
}
