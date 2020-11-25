package handlers_test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/handlers"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

func TestNewEventsReturnsEventsHandlerObject(t *testing.T) {
	events := handlers.NewEvents(nil, nil, nil)
	if events == nil {
		t.Error("New events should return event object but returned nil")
	}
}

func TestValidationMiddlewarePassesWhenObjectIsValid(t *testing.T) {
	validation, testHandler := createValidationHandler()
	if validation == nil {
		t.Error("Validation middleware should return validation handler but returned nil")
	}
	responseCode, error := verifyValidationHandler(validation, testHandler.GetCorrectBody())
	if error != nil {
		t.Error(error.Error())
	}
	if responseCode != http.StatusOK {
		t.Errorf("Expected status ok but received %d", responseCode)
	}
}

func TestValidationMiddlewareFailsWhenUnmarshallError(t *testing.T) {
	validation, testHandler := createValidationHandler()
	if validation == nil {
		t.Error("Validation middleware should return validation handler but returned nil")
	}
	responseCode, error := verifyValidationHandler(validation, testHandler.GetParseNokBody())
	if error != nil {
		t.Error(error.Error())
	}
	if responseCode != http.StatusBadRequest {
		t.Errorf("Expected status 404 but received %d", responseCode)
	}
}

func TestValidationMiddlewareFailsWhenValidationError(t *testing.T) {
	validation, testHandler := createValidationHandler()
	if validation == nil {
		t.Error("Validation middleware should return validation handler but returned nil")
	}
	responseCode, error := verifyValidationHandler(validation, testHandler.GetValidationNokBody())
	if error != nil {
		t.Error(error.Error())
	}
	if responseCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422 but received %d", responseCode)
	}
}

func createValidationHandler() (http.Handler, testHandler) {
	testLogger := log.New(os.Stdout, "> ", log.LstdFlags)
	events := handlers.NewEvents(testLogger, utils.NewValidator(), nil)
	handler := testHandler{"title", "category", "owner", "content"}
	validation := events.ValidationMiddleware(&handler)
	return validation, handler
}

func verifyValidationHandler(handler http.Handler, bodyOption io.Reader) (int, error) {
	request, error := http.NewRequest("PUT", "/test", bodyOption)
	if error != nil {
		return -1, fmt.Errorf("Could not create a request: %s", error.Error())
	}
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	return response.StatusCode, nil
}

type testHandler struct {
	Title    string
	Category string
	Owner    string
	Content  string
}

func (handler *testHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	fmt.Print("TEST HANDLER")
}

func (handler *testHandler) GetCorrectBody() *strings.Reader {
	return strings.NewReader("{\"title\": \"" + handler.Title + "\"," +
		"\"category\": \"" + handler.Category + "\"," +
		"\"owner\": \"" + handler.Owner + "\"," +
		"\"content\": \"" + handler.Content + "\"," +
		"\"date\": {\"start\": \"2020-11-11T00:00:00+00:00\", \"stop\": \"2021-11-11T00:00:01+00:00\"}}")
}

func (handler *testHandler) GetParseNokBody() *strings.Reader {
	return strings.NewReader("{\"title\": \"" + handler.Title + "\"," +
		"\"category\": \"" + handler.Category + "\"," +
		"\"owner\": \"" + handler.Owner + "\"," +
		"\"content\": \"" + handler.Content + "\"," +
		"\"date\": {\"start\": \"2020-11-1100:00:00+00:00\", \"stop\": \"2021-11-11T00:00:01+00:00\"}}")
}

func (handler *testHandler) GetValidationNokBody() *strings.Reader {
	return strings.NewReader("{\"summary\": \"" + handler.Title + "\"," +
		"\"category\": \"" + handler.Category + "\"," +
		"\"test\": \"" + handler.Owner + "\"," +
		"\"content\": \"" + handler.Content + "\"," +
		"\"date\": {\"start\": \"2020-11-11T00:00:00+00:00\", \"stop\": \"2021-11-11T00:00:01+00:00\"}}")
}
