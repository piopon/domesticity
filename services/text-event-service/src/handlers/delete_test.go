package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteEventCorrectlyRemovedEventFromDb(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	databaseIds := mockupHandler.GetDatabaseIds()
	request, error := http.NewRequest("DELETE", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": databaseIds[0].Hex()})
	recorder := httptest.NewRecorder()
	events.DeleteEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status 204 but received %d", response.StatusCode)
	}
}

func TestDeleteEventFailsIfIdDoesNotExist(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("DELETE", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": primitive.NewObjectID().Hex()})
	recorder := httptest.NewRecorder()
	events.DeleteEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 404 but received %d", response.StatusCode)
	}
}

func TestDeleteEventPanicsIfIdIsIncorrect(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	databaseIds := mockupHandler.GetDatabaseIds()
	request, error := http.NewRequest("DELETE", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": databaseIds[0].String()})
	recorder := httptest.NewRecorder()
	assert.Panics(t, func() { events.DeleteEvent(recorder, request) }, "The code did not panic")
}
