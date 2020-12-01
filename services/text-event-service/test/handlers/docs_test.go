package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/handlers"
)

func TestNewDocsReturnsDocHandlerObject(t *testing.T) {
	docHandler := handlers.NewDocs("")
	if docHandler == nil {
		t.Error("New docs returner nil object and it should not.")
	}
}

func TestGetDocumentationReturnsCorrectResult(t *testing.T) {
	docHandler := handlers.NewDocs("../resources/swagger.yaml")
	recorder := httptest.NewRecorder()
	request, error := http.NewRequest("GET", "/docs", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	docHandler.GetDocumentation(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but received %v", response.StatusCode)
	}
	contentType := response.Header.Get("content-type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("Expected HTML content but received %v", contentType)
	}
}

func TestGetSwaggerReturnsCorrectResult(t *testing.T) {
	docHandler := handlers.NewDocs("../resources/swagger.yaml")
	recorder := httptest.NewRecorder()
	request, error := http.NewRequest("GET", "/docs_test.go", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	docHandler.GetSwagger(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but received %v", response.StatusCode)
	}
	contentType := response.Header.Get("content-type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected plain text content but received %v", contentType)
	}
}
