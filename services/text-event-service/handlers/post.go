package handlers

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

// AddEvent is used to add new event and store it in DB
func (events *Events) AddEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling POST event")
	event := request.Context().Value(KeyEvent{}).(*model.Event)
	model.AddEvent(event)
}
