package dto

type (
	DeleteCartItemRequest struct {
		SessionID  string
		CartItemID uint
	}

	DeleteCartItemResponse struct {
	}
)
