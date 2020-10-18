package handlers

import (
	"log"
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/data"
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
	if http.MethodGet == request.Method {
		events.getEvents(response, request)
		return
	}
	if http.MethodPost == request.Method {
		events.addEvent(response, request)
		return
	}
	response.WriteHeader(http.StatusMethodNotAllowed)
}

func (events *Events) getEvents(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET events")
	allEvents := data.GetEvents()
	error := allEvents.ToJSON(response)
	if error != nil {
		events.logger.Println("Unable to marshal events data")
	}
}

func (events *Events) addEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling POST event")
}
