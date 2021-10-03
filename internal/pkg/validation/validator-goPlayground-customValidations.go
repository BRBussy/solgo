package validation

import "gopkg.in/go-playground/validator.v9"

type CustomValidation struct {
	Tag        string
	CustomFunc func(fl validator.FieldLevel) bool
}

func registerCustomValidations(validator *validator.Validate, customValidations []CustomValidation) (err error) {
	for _, x := range customValidations {
		if err := validator.RegisterValidation(x.Tag, x.CustomFunc); err != nil {
			return err
		}
	}

	return nil
}
