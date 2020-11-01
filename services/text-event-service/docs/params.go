package docs

import "github.com/piopon/domesticity/services/text-event-service/model"

// swagger:parameters getEvent updateEvent deleteEvent
type paramID struct {
	// The id of the event for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters addEvent updateEvent
type paramEvent struct {
	// Event data structure to update or create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body model.Event
}
