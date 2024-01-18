package validatorutil

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// PrettyValidate validates the struct and returns a pretty error message
func PrettyValidate(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		var errStr string

		for _, err := range err.(validator.ValidationErrors) {
			if errStr != "" {
				errStr += ", "
			}
			errStr += fmt.Sprintf("%s: %s", err.Field(), err.Tag())
		}

		return errors.New(errStr)
	}

	return nil
}
