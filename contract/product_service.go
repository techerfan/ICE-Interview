package contract

import (
	"context"
	"interview/internal/dto"
)

//go:generate mockgen -source=./product_service.go -destination=../internal/mocks/productservice_mock/productservice.go -package=productservicemock .

type ProductService interface {
	GetProduct(context.Context, dto.ProductGetItemRequest) (dto.ProductGetItemResponse, error)
}
