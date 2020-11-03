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
//  204: responseNoContent
//  400: errorBadQuery
//  422: errorValidation
func (events *Events) UpdateEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling PUT event")
	response.Header().Add("Content-Type", "application/json")
	id := readEventID(request)
	event := request.Context().Value(KeyEvent{}).(model.Event)
	updateError := model.UpdateEvent(id, &event)
	if updateError != nil {
		events.logger.Println("Invalid ID in PUT request")
		response.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericError{"Invalid ID in PUT request"}, response)
		return
	}
}
