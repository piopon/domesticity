package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/data"
)

// GetAllEvents is used to retrieve all currently stored events
func (events *Events) GetAllEvents(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET all events")
	allEvents := data.GetEvents()
	error := data.ToJSON(allEvents, response)
	if error != nil {
		http.Error(response, "Cannot send JSON response in GET request", http.StatusInternalServerError)
		events.logger.Println("Unable to marshal events data")
		return
	}
}
