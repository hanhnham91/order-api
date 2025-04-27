package payload

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AuthFirebaseRequest struct {
	TokenID string `json:"token_id" mod:"trim" validate:"required"`
}

func (c *AuthFirebaseRequest) Validate() error {
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
