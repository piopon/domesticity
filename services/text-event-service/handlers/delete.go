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
	deleteError := events.database.DeleteEvent(id)
	if deleteError != nil {
		events.logger.Println("Cannot delete event with specified id:", deleteError.Error())
		response.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericError{"Cannot delete event with specified id: " + deleteError.Error()}, response)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
