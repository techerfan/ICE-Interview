package cart

import (
	"context"
	"interview/internal/dto"
	"interview/internal/entity"
	"interview/internal/pkg/richerror"
)

func (s *Service) GetCartData(ctx context.Context, req dto.GetCartDataRequest) (dto.GetCartDataResponse, error) {
	const op = "calculatorservice.GetCartData"

	var cartEntity entity.Cart
	var cartExist bool
	var err error
	cartEntity, cartExist, err = s.repo.FindOpenCartBySessionID(ctx, req.SessionID)
	if !cartExist && err == nil {
		return dto.GetCartDataResponse{}, richerror.New(op).WithKind(richerror.KindNotFound).WithMessage("cart does not exist")
	} else if err != nil {
		return dto.GetCartDataResponse{}, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	var cartItems []entity.CartItem
	cartItems, err = s.repo.FindCartItemsByCartID(ctx, cartEntity.ID)
	if err != nil {
		return dto.GetCartDataResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return dto.GetCartDataResponse{
		Items: cartItems,
	}, nil
}
