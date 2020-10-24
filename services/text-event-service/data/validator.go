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
func (v *Validator) Validate(i interface{}) ValidationErrors {
	errors := v.validate.Struct(i)
	if errors == nil {
		return nil
	}
	var result ValidationErrors
	for _, error := range errors.(validator.ValidationErrors) {
		ve := ValidationError{error.(validator.FieldError)}
		result = append(result, ve)
	}
	return result
}

// ValidationErrors is a collection of ValidationError objects
type ValidationErrors []ValidationError

// Errors returns all error in a string (slice) format
func (vErrors ValidationErrors) Errors() []string {
	errors := []string{}
	for _, error := range vErrors {
		errors = append(errors, error.Error())
	}
	return errors
}

// ValidationError wraps validator FieldError to control exposed format
type ValidationError struct {
	validator.FieldError
}

// Error returns the error string
func (vError ValidationError) Error() string {
	return fmt.Sprintf("'%s' error: '%s' field validation failed on the '%s' tag",
		vError.Namespace(),
		vError.Field(),
		vError.Tag(),
	)
}
