package utils_test

import (
	"strings"
	"testing"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/utils"
)

func TestNewValidatorCreatesCorrectValidator(t *testing.T) {
	validator := utils.NewValidator()
	if validator == nil {
		t.Errorf("Could not create validator object")
	}
}

func TestValidatorCorrectlyValidatesFineObject(t *testing.T) {
	validator := utils.NewValidator()
	testStruct := validationStruct{
		ID:    1,
		Name:  "test",
		Price: 9.99,
		Date:  time.Now().String(),
	}
	errors := validator.Validate(testStruct)
	if errors != nil {
		t.Errorf("Validate returned errors but should not: %s", strings.Join(errors.Errors(), ", "))
	}
}

func TestValidatorFailsOnEmptyRequiredString(t *testing.T) {
	validator := utils.NewValidator()
	testStruct := validationStruct{
		ID:    1,
		Name:  "",
		Price: 9.99,
		Date:  time.Now().String(),
	}
	errors := validator.Validate(testStruct)
	if errors == nil {
		t.Errorf("Validate should return error on empty required string.")
	}
	t.Logf("Validate correctly returned error: %s", strings.Join(errors.Errors(), ", "))
}

func TestValidatorFailsOnIncomatibleDateFormat(t *testing.T) {
	validator := utils.NewValidator()
	testStruct := validationStruct{
		ID:    1,
		Name:  "test",
		Price: 9.99,
		Date:  "2006-01-02T15:04:05Z",
	}
	errors := validator.Validate(testStruct)
	if errors == nil {
		t.Errorf("Validate should return error on badly formatted time.")
	}
	t.Logf("Validate correctly returned error: %s", strings.Join(errors.Errors(), ", "))
}

type validationStruct struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Date  string  `json:"date" validate:"required,date-time"`
}
