package address

import "context"

type AddressRepository interface {
	GetByMetersIDs(ctx context.Context, metersIDs []int) (map[int]string, error)
}
