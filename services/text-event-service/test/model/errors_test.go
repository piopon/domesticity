package model_test

import (
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/model"
)

func TestValidationErrorsToStringRetursOkResult(t *testing.T) {
	testErrors := model.ValidationError{
		[]string{
			"error 1",
			"error 2",
			"error 3",
		},
	}
	errorStr := testErrors.ToString("* ")
	expectedString := "* error 1\n* error 2\n* error 3"
	if expectedString != errorStr {
		t.Error("Validation errors string is not correct.")
	}
}
