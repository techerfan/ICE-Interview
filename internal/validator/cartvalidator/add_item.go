package cartvalidator

import (
	"context"
	"fmt"
	"interview/internal/dto"

	validation "github.com/go-ozzo/ozzo-validation"
)

func (v Validator) ValidateAddItem(req dto.AddItemToCartRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Quantity,
			validation.Required,
			validation.Min(1)),

		validation.Field(&req.Product,
			validation.Required,
			validation.By(v.checkProductExist),
		),
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

func (v Validator) checkProductExist(value interface{}) error {
	product := value.(string)

	_, ok := v.cache.GetProduct(context.Background(), product)
	if !ok {
		return fmt.Errorf("the specified product does not exist")
	}

	return nil
}
