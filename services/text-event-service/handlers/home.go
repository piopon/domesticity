package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/utils"
)

// Home is a service handler used after visiting main service URL
type Home struct {
	logger *log.Logger
	config *utils.Config
}

// NewHome is a factory method to create Home service handler
func NewHome(logger *log.Logger, config *utils.Config) *Home {
	return &Home{logger, config}
}

// GetIndex is used to serve main page of service
func (home *Home) GetIndex(response http.ResponseWriter, request *http.Request) {
	template, parseError := template.ParseFiles("templates/index.html")
	if parseError != nil {
		fmt.Println("Got error while parsing template: " + parseError.Error())
		return
	}
	executeError := template.Execute(response, home.config)
	if executeError != nil {
		fmt.Println("Got error while executing template: " + executeError.Error())
	}
}
