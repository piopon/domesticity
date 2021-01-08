package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewHomeCreatesCorrectHomeHandler(t *testing.T) {
	home := NewCommonMockup().CreateHomeHandler("template.html")
	if home == nil {
		t.Error("New home should return Home object but returned nil")
	}
}

func TestGetIndexGeneratesCorrectResponse(t *testing.T) {
	mockupHandler := NewCommonMockup()
	home := mockupHandler.CreateHomeHandler("../../resources/index.html")
	request, error := http.NewRequest("GET", "/", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	recorder := httptest.NewRecorder()
	home.GetIndex(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but received %d", response.StatusCode)
	}
}

func TestGetIndexFailsWhenTemplateCannotBeFound(t *testing.T) {
	mockupHandler := NewCommonMockup()
	home := mockupHandler.CreateHomeHandler("template.html")
	request, error := http.NewRequest("GET", "/", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	recorder := httptest.NewRecorder()
	home.GetIndex(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 but received %d", response.StatusCode)
	}
}

func TestGetIndexFailsWhenTemplateExecutionFails(t *testing.T) {
	fileName := "test.html"
	error := createFile(fileName, badTemplate())
	if error != nil {
		t.Errorf("Could not create a temp file: %s", error.Error())
	}
	mockupHandler := NewCommonMockup()
	home := mockupHandler.CreateHomeHandler(fileName)
	request, error := http.NewRequest("GET", "/", nil)
	if error != nil {
		t.Errorf("Could not create a request: %s", error.Error())
	}
	recorder := httptest.NewRecorder()
	home.GetIndex(recorder, request)
	response := recorder.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but received %d", response.StatusCode)
	}
	err := os.Remove(fileName)
	if err != nil {
		t.Errorf("Cannot remove test file: %s", err.Error())
	}
}

func createFile(name string, contents string) error {
	file, error := os.Create(name)
	if error != nil {
		return error
	}
	_, error = file.WriteString(contents)
	if error != nil {
		file.Close()
		return error
	}
	error = file.Close()
	if error != nil {
		return error
	}
	return nil
}

func badTemplate() string {
	head := "<head>\n<title>{{ .Config.Name }}</title>\n</head>\n"
	body := "<body>\n<h1>{{ .Config.Nam }}</h1>\n</body>\n"
	return "<html>\n" + head + body + "</html>"
}
