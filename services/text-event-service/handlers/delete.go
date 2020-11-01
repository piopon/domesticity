package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

// DeleteEvent is used to delete event with specified ID stored in DB
//
// swagger:route DELETE /events/{id} events deleteEvent
// Deletes an event from DB by specified ID parameter
// responses:
//  204: noContentResponse
//  400: errorBadQuery
func (events *Events) DeleteEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling DELETE event")
	id := readEventID(request)
	deleteError := model.DeleteEvent(id)
	if deleteError != nil {
		http.Error(response, "Invalid ID in DELETE request", http.StatusBadRequest)
		events.logger.Println("Invalid event ID")
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
