package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

// UpdateEvent is used to update event with specified ID stored in DB
//
// swagger:route PUT /events/{id} events updateEvent
// Updates an event in DB by specified ID parameter
// responses:
//  200: responseEvent
//  400: errorBadQuery
//  422: errorValidation
//  500: errorInternal
func (events *Events) UpdateEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling PUT event")
	response.Header().Add("Content-Type", "application/json")
	id := readEventID(request)
	event := request.Context().Value(KeyEvent{}).(*model.Event)
	updateError := model.UpdateEvent(id, event)
	if updateError != nil {
		events.logger.Println("Invalid ID in PUT request")
		response.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericError{"Invalid ID in PUT request"}, response)
		return
	}
	jsonError := utils.ToJSON(event, response)
	if jsonError != nil {
		events.logger.Println("Unable to marshal events data in UpdateEvent handler")
		response.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericError{"Cannot send JSON response in PUT request"}, response)
		return
	}
}
