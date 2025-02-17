package mysql

import (
	"interview/internal/entity"

	"gorm.io/gorm"
)

type (
	Cart struct {
		gorm.Model

		Total     float64
		SessionID string
		Status    string
	}

	CartItem struct {
		gorm.Model

		CartID      uint
		ProductName string
		Quantity    int
		Price       float64
	}
)

func mapCartToEntity(cart Cart) entity.Cart {
	return entity.Cart{
		ID:        cart.ID,
		Total:     cart.Total,
		SessionID: cart.SessionID,
		Status:    entity.CartStatus(cart.Status),
	}
}

func mapEntityToCart(cart entity.Cart) Cart {
	return Cart{
		Model: gorm.Model{
			ID: cart.ID,
		},
		SessionID: cart.SessionID,
		Total:     cart.Total,
		Status:    string(cart.Status),
	}
}

func mapCartItemToEntity(cartItem CartItem) entity.CartItem {
	return entity.CartItem{
		ID:          cartItem.ID,
		CartID:      cartItem.CartID,
		ProductName: cartItem.ProductName,
		Quantity:    cartItem.Quantity,
		Price:       cartItem.Price,
	}
}

func mapEntityToCartItem(cartItem entity.CartItem) CartItem {
	return CartItem{
		Model: gorm.Model{
			ID: cartItem.ID,
		},
		CartID:      cartItem.CartID,
		ProductName: cartItem.ProductName,
		Quantity:    cartItem.Quantity,
		Price:       cartItem.Price,
	}
}
