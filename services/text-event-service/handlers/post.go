package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

// AddEvent is used to add new event and store it in DB
//
// swagger:route POST /events events addEvent
// Creates a new event and adds it to DB
// responses:
//  200: responsePostEvent
//  400: errorBadQuery
//  422: errorValidation
//  500: errorInternal
func (events *Events) AddEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling POST event")
	response.Header().Add("Content-Type", "application/json")
	event := request.Context().Value(KeyEvent{}).(*model.Event)
	error := events.database.AddEvent(event)
	if error != nil {
		events.logger.Println("Cannot add new event to database:", error.Error())
		response.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericError{"Cannot add new event to database: " + error.Error()}, response)
		return
	}
	jsonError := utils.ToJSON(event, response)
	if jsonError != nil {
		events.logger.Println("Unable to marshal events data in AddEvent handler")
		response.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericError{"Cannot send JSON response in POST request"}, response)
		return
	}
}
