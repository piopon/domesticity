package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/data"
)

// DeleteEvent is used to delete event with specified ID stored in DB
func (events *Events) DeleteEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling DELETE event")
	id := readEventID(request)
	deleteError := data.DeleteEvent(id)
	if deleteError != nil {
		http.Error(response, "Invalid ID in DELETE request", http.StatusBadRequest)
		events.logger.Println("Invalid event ID")
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
