package mysql

import (
	"context"
	"errors"
	"interview/internal/entity"

	"gorm.io/gorm"
)

func (m *Mysql) CreateCartItem(ctx context.Context, cartItem entity.CartItem) (entity.CartItem, error) {
	var model = mapEntityToCartItem(cartItem)
	if err := m.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entity.CartItem{}, err
	}

	return mapCartItemToEntity(model), nil
}

func (m *Mysql) UpdateCartItem(ctx context.Context, cartItem entity.CartItem) error {
	var model CartItem

	if err := m.db.WithContext(ctx).Where("id = ?", cartItem.ID).First(&model).Error; err != nil {
		return err
	}

	model.CartID = cartItem.CartID
	model.Price = cartItem.Price
	model.ProductName = cartItem.ProductName
	model.Quantity = cartItem.Quantity

	return m.db.WithContext(ctx).Save(&model).Error
}

func (m *Mysql) FindCartItemByID(ctx context.Context, id uint) (entity.CartItem, bool, error) {
	var cartItem CartItem

	if err := m.db.WithContext(ctx).Where("id = ?", id).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.CartItem{}, false, nil
		}
		return entity.CartItem{}, false, err
	}

	return mapCartItemToEntity(cartItem), true, nil
}

func (m *Mysql) FindCartItemByProduct(ctx context.Context, cartID uint, product string) (entity.CartItem, bool, error) {
	var cartItem CartItem

	if err := m.db.WithContext(ctx).Where("cart_id = ? AND product_name = ?", cartID, product).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.CartItem{}, false, nil
		}
		return entity.CartItem{}, false, err
	}

	return mapCartItemToEntity(cartItem), true, nil
}

func (m *Mysql) FindCartItemsByCartID(ctx context.Context, cartID uint) ([]entity.CartItem, error) {
	var cartItems []CartItem

	if err := m.db.WithContext(ctx).Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return nil, err
	}

	items := make([]entity.CartItem, 0, len(cartItems))
	for _, item := range cartItems {
		items = append(items, mapCartItemToEntity(item))
	}

	return items, nil
}

func (m *Mysql) DeleteCartItemByID(ctx context.Context, id uint) error {
	return m.db.WithContext(ctx).Where("id = ?", id).Delete(&CartItem{}).Error
}
