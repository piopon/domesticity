package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/piopon/domesticity/services/text-event-service/data"
)

// UpdateEvent is used to update event with specified ID stored in DB
func (events *Events) UpdateEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling PUT event")
	id, error := strconv.Atoi(mux.Vars(request)["id"])
	if error != nil {
		http.Error(response, "Bad URL", http.StatusBadRequest)
	}
	event := request.Context().Value(KeyEvent{}).(data.Event)
	updateError := data.UpdateEvent(id, &event)
	if updateError != nil {
		events.logger.Println("Invalid event ID")
		return
	}
}
