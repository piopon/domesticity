package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validator contains validate rules
type Validator struct {
	validate *validator.Validate
}

// NewValidator is a factory method to create a new Validator type
func NewValidator() *Validator {
	validate := validator.New()
	validate.RegisterValidation("date-time", validateDateTime)
	return &Validator{validate}
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

// validateDateTime is added because go-validator and JSON decoded format cannot be matched
// problem:
//	format "2006-01-02T15:04:05+07:00" = not accepted by validator, decoded by JSON.Decode()
//	format "2006-01-02T15:04:05Z07:00" = accepted by validator, JSON.Decode() fails
func validateDateTime(field validator.FieldLevel) bool {
	dateRegex := `(\d+)-(0[1-9]|1[012])-(0[1-9]|[12]\d|3[01])`
	timeRegex := `([01]\d|2[0-3]):([0-5]\d):([0-5]\d|60)(\.\d+)?`
	zoneRegex := `(([Zz])|([\+|\-]([01]\d|2[0-3])))`
	regex := regexp.MustCompile(dateRegex + `\s` + timeRegex + `\s` + zoneRegex)
	datetime := regex.FindAllString(fmt.Sprintln(field.Field().Interface()), -1)
	if len(datetime) == 1 {
		return true
	}
	return false
}
