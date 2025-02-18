package cartvalidator

import (
	"context"
	"interview/internal/entity"
)

type CacheRepository interface {
	GetProduct(ctx context.Context, productName string) (float64, bool)
}

type Repository interface {
	FindCartItemByID(ctx context.Context, id uint) (entity.CartItem, bool, error)
	FindCartByID(ctx context.Context, cartID uint) (entity.Cart, bool, error)
}

type Validator struct {
	cache CacheRepository
	repo  Repository
}

func New(cache CacheRepository, repo Repository) Validator {
	return Validator{
		cache: cache,
		repo:  repo,
	}
}
