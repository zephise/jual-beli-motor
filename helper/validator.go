package helper

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}
