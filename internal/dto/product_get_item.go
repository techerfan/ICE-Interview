package dto

type (
	ProductGetItemRequest struct {
		ProductName string
	}

	ProductGetItemResponse struct {
		ProductName string
		Price       float64
	}
)
