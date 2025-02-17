package dto

type (
	AddItemToCartRequest struct {
		SessionID string
		Product   string
		Quantity  int
	}

	AddItemToCartResponse struct {
	}
)
