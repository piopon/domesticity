package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/handlers"
	"github.com/piopon/domesticity/services/text-event-service/src/model"
)

func TestAddEventCorrectlyAddEntryToDb(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("POST", "/events", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	eventToAdd := &model.Event{}
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, eventToAdd)
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
	eventToAdd := &model.Event{ID: mockupHandler.GetDatabaseIds()[0]}
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, eventToAdd)
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
	eventToAdd := mockupHandler.CreateEventBadJson()
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, eventToAdd)
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	events.AddEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 but received %d", response.StatusCode)
	}
}
