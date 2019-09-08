package external

import (
	"context"
)

type Locator interface {
	GetCarsNearPoint(ctx context.Context, lat, lng float64) ([]Car, error)
}

type Timer interface {
	GetETAs(ctx context.Context, cars []Car) ([]int, error)
}
