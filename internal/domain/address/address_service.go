package address

import "context"

type AddressService interface {
	GetByMetersIDs(ctx context.Context, metersIDs []int) (map[int]string, error)
}
