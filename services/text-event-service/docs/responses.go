package docs

import "github.com/piopon/domesticity/services/text-event-service/model"

// ResponseGetEvents returns all or filtered events from DB (depends on query params)
// swagger:response responseGetEvents
type responseGetEvents struct {
	// All / filtered events from DB
	// in: body
	Body model.Events
}

// ResponseGetEvent return single event from DB (depends on ID parameter)
// swagger:response responseGetEvent
type responseGetEvent struct {
	// Single event from DB (selected by ID parameter)
	// in: body
	Body model.Event
}

// ResponseNoContent is returned when no specific response is needed
// swagger:response responseNoContent
type responseNoContent struct {
}
