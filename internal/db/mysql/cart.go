package mysql

import (
	"context"
	"errors"
	"interview/internal/entity"

	"gorm.io/gorm"
)

func (m *Mysql) CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error) {
	var model = mapEntityToCart(cart)
	if err := m.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entity.Cart{}, err
	}

	return mapCartToEntity(model), nil
}

func (m *Mysql) UpdateCart(ctx context.Context, cart entity.Cart) error {
	var model = mapEntityToCart(cart)

	return m.db.WithContext(ctx).Save(&model).Error
}

func (m *Mysql) FindCartByID(ctx context.Context, cartID uint) (entity.Cart, bool, error) {
	var cart Cart
	if err := m.db.WithContext(ctx).Where("id = ?", cartID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Cart{}, false, nil
		}
		return entity.Cart{}, false, err
	}

	return mapCartToEntity(cart), true, nil
}

func (m *Mysql) FindOpenCartBySessionID(ctx context.Context, sessionID string) (entity.Cart, bool, error) {
	var cart Cart

	if err := m.db.WithContext(ctx).Where("status = ? AND session_id = ?", entity.CartOpen, sessionID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Cart{}, false, nil
		}
		return entity.Cart{}, false, err
	}

	return mapCartToEntity(cart), true, nil
}
