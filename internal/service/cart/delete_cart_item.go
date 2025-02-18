package cart

import (
	"context"
	"interview/internal/dto"
	"interview/internal/entity"
	"interview/internal/pkg/richerror"
)

func (s *Service) DeleteCartItem(ctx context.Context, req dto.DeleteCartItemRequest) error {
	const op = "calculatorservice.DeleteCartItem"

	var cartEntity entity.Cart
	var cartExist bool
	var err error
	cartEntity, cartExist, err = s.repo.FindOpenCartBySessionID(ctx, req.SessionID)
	if !cartExist && err != nil {
		return richerror.New(op).WithKind(richerror.KindNotFound).WithErr(err).
			WithMessage("cart does not exist")
	} else if err != nil {
		return richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	if cartEntity.Status == entity.CartClosed {
		// TODO:
		// c.Redirect(302, "/")
		return richerror.New(op).WithKind(richerror.KindInvalid).WithMessage("cart is closeds")
	}

	// TODO: item existense must be checked using the validator
	// var cartItemEntity entity.CartItem

	// result = db.Where(" ID  = ?", cartItemID).First(&cartItemEntity)
	// if result.Error != nil {
	// 	c.Redirect(302, "/")
	// 	return
	// }

	if err := s.repo.DeleteCartItemByID(ctx, req.CartItemID); err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
