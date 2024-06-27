package helpers

import (
	"github.com/go-playground/validator"
)

func Validator(data interface{}) error {
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		return err
	}

	return err
}
