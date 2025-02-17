package dto

import "interview/internal/entity"

type (
	GetCartDataRequest struct {
		SessionID string
	}

	GetCartDataResponse struct {
		Items []entity.CartItem
	}
)
