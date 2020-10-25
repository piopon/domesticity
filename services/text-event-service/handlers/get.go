package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

// GetAllEvents is used to retrieve all currently stored events
func (events *Events) GetAllEvents(response http.ResponseWriter, request *http.Request) {
	allEvents := model.GetEvents()
	error := utils.ToJSON(allEvents, response)
	if error != nil {
		http.Error(response, "Cannot send JSON response in GET request", http.StatusInternalServerError)
		events.logger.Println("Unable to marshal events data")
		return
	}
}

// GetSingleEvent is used to retrieve stored events with specified ID
func (events *Events) GetSingleEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET single event")
	id := readEventID(request)
	event, error := model.GetEventByID(id)
	if error != nil {
		http.Error(response, "Cannot find event with specified ID in GET request", http.StatusNotFound)
		events.logger.Println("Unable to find event with specified id:", id)
	}
	jsonError := utils.ToJSON(event, response)
	if jsonError != nil {
		http.Error(response, "Cannot send JSON response in GET request", http.StatusInternalServerError)
		events.logger.Println("Unable to marshal events data")
		return
	}
}
