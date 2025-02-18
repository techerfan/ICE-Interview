package product

import (
	"context"
	"time"
)

var defaultItemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

type CachRepository interface {
	SetProduct(ctx context.Context, productName string, cost float64, expTime time.Duration) error
	GetProduct(ctx context.Context, productName string) (float64, bool)
	DeleteProduct(ctx context.Context, productName string) error
}

type Service struct {
	repo CachRepository
}

func New(repo CachRepository) *Service {

	for product, price := range defaultItemPriceMapping {
		_, ok := repo.GetProduct(context.Background(), product)
		if !ok {
			if err := repo.SetProduct(context.Background(), product, price, 0); err != nil {
				// Default products must exist in the cache store. Otherwise, our app
				// cannot do as it supposed to
				panic(err)
			}
		}
	}

	return &Service{
		repo: repo,
	}
}
