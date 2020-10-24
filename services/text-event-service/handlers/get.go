package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/data"
)

// GetEvents is used to retrieve all currently stored events
func (events *Events) GetEvents(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET events")
	allEvents := data.GetEvents()
	error := data.ToJSON(allEvents, response)
	if error != nil {
		events.logger.Println("Unable to marshal events data")
	}
}
