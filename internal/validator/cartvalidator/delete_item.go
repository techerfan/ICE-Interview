package cartvalidator

import (
	"context"
	"fmt"
	"interview/internal/dto"
	"interview/internal/entity"

	validation "github.com/go-ozzo/ozzo-validation"
)

func (v Validator) ValidateDeleteItem(req dto.DeleteCartItemRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.CartItemID,
			validation.Required,
			validation.By(v.checkCartItemExist),
			validation.By(v.checkCartIsOpen)),
	); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, err
	}

	return nil, nil
}

func (v Validator) checkCartItemExist(value interface{}) error {
	cartItemID := value.(uint)

	_, exist, err := v.repo.FindCartItemByID(context.Background(), cartItemID)
	if !exist && err == nil {
		return fmt.Errorf("cart item does not exist")
	} else if err != nil {
		return fmt.Errorf("cannot find the cart item due to an internal error")
	}

	return nil
}

func (v Validator) checkCartIsOpen(value interface{}) error {
	cartItemID := value.(uint)

	cartItem, exist, err := v.repo.FindCartItemByID(context.Background(), cartItemID)
	if !exist && err == nil {
		return fmt.Errorf("cart item does not exist")
	} else if err != nil {
		return fmt.Errorf("cannot find the cart item due to an internal error")
	}

	cart, exist, err := v.repo.FindCartByID(context.Background(), cartItem.CartID)
	if !exist && err == nil {
		return fmt.Errorf("cart does not exist")
	} else if err != nil {
		return fmt.Errorf("cannot find the cart due to an internal error")
	}

	if cart.Status == entity.CartClosed {
		return fmt.Errorf("cannot remove an item from a closed cart")
	}

	return nil
}
