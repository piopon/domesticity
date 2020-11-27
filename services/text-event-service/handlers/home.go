package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Home is a service handler used after visiting main service URL
type Home struct {
	logger *log.Logger
}

// HomeContent is a struct containing a content for index/home page
type HomeContent struct {
	Name string
}

// NewHome is a factory method to create Home service handler
func NewHome(logger *log.Logger) *Home {
	return &Home{logger}
}

// GetIndex is used to serve main page of service
func (home *Home) GetIndex(response http.ResponseWriter, request *http.Request) {
	content := HomeContent{"test template htmla"}
	template, parseError := template.ParseFiles("templates/index.html")
	if parseError != nil {
		fmt.Println("Got error while parsing template: " + parseError.Error())
		return
	}
	executeError := template.Execute(response, content)
	if executeError != nil {
		fmt.Println("Got error while executing template: " + executeError.Error())
	}
}
