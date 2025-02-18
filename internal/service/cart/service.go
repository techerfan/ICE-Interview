package cart

import (
	"context"
	"interview/contract"
	"interview/internal/entity"
)

type Repository interface {
	// Cart functionalities
	CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error)
	UpdateCart(ctx context.Context, cart entity.Cart) error
	FindOpenCartBySessionID(ctx context.Context, sessionID string) (entity.Cart, bool, error)

	// Cart item functionalities
	CreateCartItem(ctx context.Context, cartItem entity.CartItem) (entity.CartItem, error)
	UpdateCartItem(ctx context.Context, cartItem entity.CartItem) error
	FindCartItemByID(ctx context.Context, id uint) (entity.CartItem, bool, error)
	FindCartItemByProduct(ctx context.Context, cartID uint, product string) (entity.CartItem, bool, error)
	FindCartItemsByCartID(ctx context.Context, cartID uint) ([]entity.CartItem, error)
	DeleteCartItemByID(ctx context.Context, id uint) error
}

type Service struct {
	repo           Repository
	productService contract.ProductService
}

func New(repo Repository, productService contract.ProductService) *Service {
	return &Service{
		repo:           repo,
		productService: productService,
	}
}
