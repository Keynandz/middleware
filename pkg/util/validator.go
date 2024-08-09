package util

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidation struct {
	Validator *validator.Validate
}

func (cv *CustomValidation) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
