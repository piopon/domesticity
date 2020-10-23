package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// Validator contains validate rules
type Validator struct {
	validate *validator.Validate
}

// NewValidator is a factory method to create a new Validator type
func NewValidator() *Validator {
	return &Validator{validator.New()}
}

// Validate is a Validator method used to inspect inputted interface
func (validator *Validator) Validate(i interface{}) error {
	return validator.validate.Struct(i)
}

// ValidationError wraps validator FieldError to control exposed format
type ValidationError struct {
	validator.FieldError
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("'%s' error: '%s' field validation failed on the '%s' tag",
		validationError.Namespace(),
		validationError.Field(),
		validationError.Tag(),
	)
}
