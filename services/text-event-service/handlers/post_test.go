package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/handlers"
	"github.com/piopon/domesticity/services/text-event-service/model"
)

func TestAddEventCorrectlyAddEntryToDb(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("POST", "/events", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, &model.Event{})
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	events.AddEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but received %d", response.StatusCode)
	}
}

func TestAddEventFailsIfEntryCannotBeAdded(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("POST", "/events", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	id := mockupHandler.GetDatabaseIds()[0]
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, &model.Event{ID: id})
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	events.AddEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 but received %d", response.StatusCode)
	}
}

func TestAddEventFailsIfEntryCannotBeParsed(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("POST", "/events", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, createInvalidEvent())
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	events.AddEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 but received %d", response.StatusCode)
	}
}

func createInvalidEvent() *model.Event {
	return &model.Event{
		Title: "This is my first event",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(-2020, 05, 26, 14, 15, 00, 00, time.Local),
			Stop:  time.Date(-2020, 05, 27, 10, 30, 00, 00, time.Local)},
		Category: "Notes",
		Content:  "Test event number 1",
	}
}
