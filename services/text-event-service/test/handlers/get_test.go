package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetEventsCorrectlyReceivesEventsFromDb(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("GET", "/events", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	recorder := httptest.NewRecorder()
	events.GetEvents(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but received %d", response.StatusCode)
	}
}

func TestGetEventsFailsIfEventsCannotBeRead(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("GET", "/events", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	query := request.URL.Query()
	query.Add("limit", "1.1")
	request.URL.RawQuery = query.Encode()
	recorder := httptest.NewRecorder()
	events.GetEvents(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 404 but received %d", response.StatusCode)
	}
}

func TestGetEventsFailsIfEventCannotBeParsed(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	mockupHandler.AddBadEventToDB()
	request, error := http.NewRequest("GET", "/events", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	recorder := httptest.NewRecorder()
	events.GetEvents(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 but received %d", response.StatusCode)
	}
}

func TestGetEventCorrectlyReceivesEventFromDb(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	databaseIds := mockupHandler.GetDatabaseIds()
	request, error := http.NewRequest("GET", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": databaseIds[0].Hex()})
	recorder := httptest.NewRecorder()
	events.GetEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but received %d", response.StatusCode)
	}
}

func TestGetEventFailsIfEventCannotBeParsed(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	badEventId := mockupHandler.AddBadEventToDB()
	request, error := http.NewRequest("GET", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": badEventId.Hex()})
	recorder := httptest.NewRecorder()
	events.GetEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 but received %d", response.StatusCode)
	}
}

func TestGetEventFailsIfIdDoesNotExist(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	request, error := http.NewRequest("GET", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": primitive.NewObjectID().Hex()})
	recorder := httptest.NewRecorder()
	events.GetEvent(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404 but received %d", response.StatusCode)
	}
}

func TestGetEventPanicsIfIdIsIncorrect(t *testing.T) {
	mockupHandler := NewCommonMockup()
	events := mockupHandler.CreateEventsHandler()
	databaseIds := mockupHandler.GetDatabaseIds()
	request, error := http.NewRequest("GET", "/events/{id}", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	request = mux.SetURLVars(request, map[string]string{"id": databaseIds[0].String()})
	recorder := httptest.NewRecorder()
	assert.Panics(t, func() { events.GetEvent(recorder, request) }, "The code did not panic")
}
