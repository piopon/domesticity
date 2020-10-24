package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/piopon/domesticity/services/text-event-service/data"
)

// DeleteEvent is used to delete event with specified ID stored in DB
func (events *Events) DeleteEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling DELETE event")
	id, error := strconv.Atoi(mux.Vars(request)["id"])
	if error != nil {
		http.Error(response, "Bad URL", http.StatusBadRequest)
	}
	deleteError := data.DeleteEvent(id)
	if deleteError != nil {
		events.logger.Println("Invalid event ID")
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
