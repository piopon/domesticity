package model_test

import (
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/model"
)

func TestHealthToStringRetursOkResult(t *testing.T) {
	testErrors := model.Health{"ALL RIGHTY!!!"}
	healthStr := testErrors.ToString()
	expectedString := "ALL RIGHTY!!!"
	if expectedString != healthStr {
		t.Error("Health string is not correct.")
	}
}
