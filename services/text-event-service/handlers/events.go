package handlers

import (
	"log"
	"net/http"
)

// Events is a service handler responsible for getting and updating events
type Events struct {
	logger *log.Logger
}

// NewEvents is a factory method to create Events service handler with defined logger
func NewEvents(logger *log.Logger) *Events {
	return &Events{logger}
}

func (events *Events) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
