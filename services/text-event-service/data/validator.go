package data

import "github.com/go-playground/validator"

// Validator contains validate rules
type Validator struct {
	validate *validator.Validate
}

// NewValidator is a factory method to create a new Validator type
func NewValidator() *Validator {
	return &Validator{validator.New()}
}
