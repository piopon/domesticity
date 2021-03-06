package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
)

// GetEvents is used to retrieve currently stored events
//
// swagger:route GET /events events getEvents
// Returns a list of currently stored events (all if no query params is used or filtered otherwise)
// responses:
//  200: responseGetEvents
//  400: errorBadQuery
//  500: errorInternal
func (events *Events) GetEvents(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET all/filtered events")
	response.Header().Add("Content-Type", "application/json")
	allEvents, error := events.database.GetEvents(request.URL.Query())
	if error != nil {
		events.logger.Println("Unable to receive all events from database:", error.Error())
		response.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericError{Message: "Unable to receive all events from database: " + error.Error()}, response)
		return
	}
	jsonError := utils.ToJSON(allEvents, response)
	if jsonError != nil {
		events.logger.Println("Unable to marshal events data in GetEvents handler")
		response.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericError{Message: "Cannot send JSON response in GET request"}, response)
		return
	}
}

// GetEvent is used to retrieve stored events with specified ID
//
// swagger:route GET /event/{id} events getEvent
// Returns a event with provided id
// responses:
//  200: responseGetEvent
//  404: errorNotFound
//  500: errorInternal
func (events *Events) GetEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET single event")
	response.Header().Add("Content-Type", "application/json")
	id := readEventID(request)
	event, error := events.database.GetEvent(id)
	if error != nil {
		events.logger.Println("Unable to receive a single event from database:", error.Error())
		response.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&model.GenericError{Message: "Unable to receive a single event from database: " + error.Error()}, response)
		return
	}
	jsonError := utils.ToJSON(event, response)
	if jsonError != nil {
		events.logger.Println("Unable to marshal events data in GetEvent handler")
		response.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericError{Message: "Cannot send JSON response in GET request"}, response)
		return
	}
}
