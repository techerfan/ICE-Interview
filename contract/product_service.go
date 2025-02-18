package contract

import (
	"context"
	"interview/internal/dto"
)

type ProductService interface {
	GetProduct(context.Context, dto.ProductGetItemRequest) (dto.ProductGetItemResponse, error)
}
