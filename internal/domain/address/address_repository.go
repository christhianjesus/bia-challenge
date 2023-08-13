package address

import "context"

type AddressRepository interface {
	GetByMeterIDs(ctx context.Context, meterIDs []int) (map[int]string, error)
}
