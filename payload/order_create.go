package payload

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

//nolint:tagliatelle
type OrderItem struct {
	ProductID int64 `json:"productId" mod:"trim" validate:"required,min=1"`
	Quantity  int   `json:"quantity" mod:"trim" validate:"required,min=1"`
}

//nolint:tagliatelle
type CreateOrderRequest struct {
	CouponCode string      `json:"couponCode" mod:"trim" validate:"lte=128"`
	Items      []OrderItem `json:"items" mod:"trim,ucase" validate:"required,min=1,dive"`
}

func (c *CreateOrderRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				return fmt.Errorf("field: '%s', tag: '%s', value: '%v', param: '%s'",
					e.Field(), e.Tag(), e.Value(), e.Param())
			}
		}
	}

	return nil
}
