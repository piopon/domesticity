package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/piopon/domesticity/services/text-event-service/handlers"
	"github.com/piopon/domesticity/services/text-event-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateEventCorrectlyUpdatesEntryInDb(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	databaseIds := mockupHandler.GetDatabaseIds()
	request, error := http.NewRequest("PUT", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": databaseIds[0].Hex()})
	eventToUpdate := &model.Event{}
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, eventToUpdate)
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	events.UpdateEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but received %d", response.StatusCode)
	}
}

func TestUpdateEventFailsIfEntryCannotBeUpdated(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("PUT", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": primitive.NewObjectID().Hex()})
	eventToUpdate := mockupHandler.CreateEventBadJson()
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, eventToUpdate)
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	events.UpdateEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 but received %d", response.StatusCode)
	}
}

func TestUpdateEventFailsIfEntryCannotBeParsed(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	databaseIds := mockupHandler.GetDatabaseIds()
	request, error := http.NewRequest("PUT", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": databaseIds[0].Hex()})
	eventToUpdate := mockupHandler.CreateEventBadJson()
	ctx := context.WithValue(request.Context(), handlers.KeyEvent{}, eventToUpdate)
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	events.UpdateEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 but received %d", response.StatusCode)
	}
}
