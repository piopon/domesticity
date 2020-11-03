package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

// DeleteEvent is used to delete event with specified ID stored in DB
//
// swagger:route DELETE /events/{id} events deleteEvent
// Deletes an event from DB by specified ID parameter
// responses:
//  204: responseNoContent
//  400: errorBadQuery
func (events *Events) DeleteEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling DELETE event")
	response.Header().Add("Content-Type", "application/json")
	id := readEventID(request)
	deleteError := model.DeleteEvent(id)
	if deleteError != nil {
		events.logger.Println("Invalid ID in DELETE request")
		response.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericError{"Invalid ID in DELETE request"}, response)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
