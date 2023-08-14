package address

import (
	"context"

	"github.com/christhianjesus/bia-challenge/internal/domain/address"
)

type addressService struct {
	repo address.AddressRepository
}

func NewAddressService(repo address.AddressRepository) address.AddressService {
	return &addressService{repo}
}

func (as *addressService) GetByMetersIDs(ctx context.Context, metersIDs []int) (map[int]string, error) {
	return as.repo.GetByMetersIDs(ctx, metersIDs)
}
