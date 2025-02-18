package cart

import (
	"context"
	"interview/internal/dto"
	"interview/internal/entity"
	"interview/internal/pkg/richerror"
)

func (s *Service) AddItemToCart(ctx context.Context, req dto.AddItemToCartRequest) error {
	const op = "calculatorservice.AddItemToCart"

	var isCartNew bool

	cartEntity, exist, err := s.repo.FindOpenCartBySessionID(ctx, req.SessionID)
	if err != nil {
		// TODO: c.Redirect(302, "/")
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if !exist {
		isCartNew = true
		cartEntity = entity.Cart{
			SessionID: req.SessionID,
			Status:    entity.CartOpen,
		}

		cartEntity, err = s.repo.CreateCart(ctx, cartEntity)
		if err != nil {
			return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
	}

	// TODO: this must be validated in the validator layer
	// item, ok := itemPriceMapping[req.Product]
	// if !ok {
	// 	// TODO: add the following line to the controller
	// 	// c.Redirect(302, "/?error=invalid item name")
	// 	return richerror.New(op).WithMessage("invalid item name").WithKind(richerror.KindInvalid)
	// }

	item, err := s.productService.GetProduct(ctx, dto.ProductGetItemRequest{ProductName: req.Product})
	if err != nil {
		// Since we have validated the request and at this point we are sure
		// that the product exist, if any error occurs, it is unexpected
		return richerror.New(op).WithMessage("invalid item name").WithKind(richerror.KindUnexpected)
	}

	var cartItemEntity entity.CartItem
	if isCartNew {
		cartItemEntity = entity.CartItem{
			CartID:      cartEntity.ID,
			ProductName: req.Product,
			Quantity:    req.Quantity,
			Price:       item.Price * float64(req.Quantity),
		}

		cartItemEntity, err = s.repo.CreateCartItem(ctx, cartItemEntity)
		if err != nil {
			return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
	} else {
		cartItemEntity, exist, err = s.repo.FindCartItemByProduct(ctx, cartEntity.ID, req.Product)

		if err != nil {
			// TODO: handle it in the controller
			// c.Redirect(302, "/")
			return richerror.New(op).WithMessage("invalid item name").WithKind(richerror.KindUnexpected)
		}

		if !exist {
			cartItemEntity = entity.CartItem{
				CartID:      cartEntity.ID,
				ProductName: req.Product,
				Quantity:    req.Quantity,
				Price:       item.Price * float64(req.Quantity),
			}

			cartItemEntity, err = s.repo.CreateCartItem(ctx, cartItemEntity)
			if err != nil {
				return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
			}

		} else {
			cartItemEntity.Quantity += int(req.Quantity)
			cartItemEntity.Price += item.Price * float64(req.Quantity)

			err = s.repo.UpdateCartItem(ctx, cartItemEntity)
			if err != nil {
				return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
			}
		}
	}

	return nil
}
