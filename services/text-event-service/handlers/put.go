package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

// UpdateEvent is used to update event with specified ID stored in DB
func (events *Events) UpdateEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling PUT event")
	id := readEventID(request)
	event := request.Context().Value(KeyEvent{}).(model.Event)
	updateError := model.UpdateEvent(id, &event)
	if updateError != nil {
		http.Error(response, "Invalid ID in PUT request", http.StatusBadRequest)
		events.logger.Println("Invalid event ID")
		return
	}
}
