package product

import (
	"context"
	"fmt"
	"interview/internal/dto"
)

func (s *Service) GetProduct(ctx context.Context, req dto.ProductGetItemRequest) (dto.ProductGetItemResponse, error) {
	price, ok := s.repo.GetProduct(ctx, req.ProductName)
	if !ok {
		return dto.ProductGetItemResponse{}, fmt.Errorf("the specified product does not exist")
	}

	return dto.ProductGetItemResponse{
		ProductName: req.ProductName,
		Price:       price,
	}, nil
}
