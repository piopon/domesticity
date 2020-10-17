package handlers

import (
	"encoding/json"
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
	allEvents := data.GetEvents()
	data, error := json.Marshal(allEvents)
	if error != nil {
		events.logger.Println("Unable to marshal events data")
	}
	response.Write(data)
}
