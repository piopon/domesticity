package model_test

import (
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

func TestEventToStringRetursOkResult(t *testing.T) {
	testEvent := model.Event{
		Title:    "TestTitle",
		Owner:    "TestOwner",
		Category: "TestCategory",
	}
	errorStr := testEvent.ToString()
	expectedString := "Event: \"TestTitle\" created by: TestOwner [category: TestCategory]"
	if expectedString != errorStr {
		t.Error("Event returned incorrect string result.")
	}
}

func TestEventToStringRetursCorrectResultWithEmptyFields(t *testing.T) {
	testEvent := model.Event{
		Title:    "",
		Owner:    "",
		Category: "",
	}
	errorStr := testEvent.ToString()
	expectedString := "Event: \"NIL\" created by: NIL [category: NIL]"
	if expectedString != errorStr {
		t.Error("Event returned incorrect string result.")
	}
}
