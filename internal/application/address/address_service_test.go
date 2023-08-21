package address

import (
	"context"
	"testing"

	"github.com/christhianjesus/bia-challenge/internal/domain/address"
	"github.com/christhianjesus/bia-challenge/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func setupAddressService(t *testing.T) *addressServiceMock {
	repoMock := mocks.NewAddressRepository(t)

	return &addressServiceMock{
		repo: repoMock,
		srv:  NewAddressService(repoMock),
	}
}

type addressServiceMock struct {
	repo *mocks.AddressRepository
	srv  address.AddressService
}

func TestGetByMetersIDs_ApiError(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}

	asm := setupAddressService(t)
	asm.repo.On("GetByMetersIDs", ctx, metersIDs).Return(nil, assert.AnError)

	addresses, err := asm.srv.GetByMetersIDs(ctx, metersIDs)

	assert.Nil(t, addresses)
	assert.Error(t, err)
}

func TestGetByMetersIDs_ApiOK(t *testing.T) {
	ctx := context.TODO()
	metersIDs := []int{1, 2}
	addressResponse := map[int]string{
		1: "Dirección Mock 1",
		2: "Dirección Mock 2",
	}

	asm := setupAddressService(t)
	asm.repo.On("GetByMetersIDs", ctx, metersIDs).Return(addressResponse, nil)

	addresses, err := asm.srv.GetByMetersIDs(ctx, metersIDs)

	assert.NotNil(t, addresses)
	assert.NoError(t, err)
	assert.Len(t, addresses, 2)
	assert.Equal(t, "Dirección Mock 1", addresses[1])
	assert.Equal(t, "Dirección Mock 2", addresses[2])
}
