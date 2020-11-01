package docs

import "github.com/piopon/domesticity/services/text-event-service/model"

// EventsResponse with all or filtered events from DB (depends on query params)
// swagger:response eventsResponse
type eventsResponse struct {
	// All / filtered events from DB
	// in: body
	Body model.Events
}

// EventResponse with single event from DB (depends on ID parameter)
// swagger:response eventResponse
type eventResponse struct {
	// Single event from DB
	// in: body
	Body model.Event
}

// NoContentResponse is returned when no specific response is needed
// swagger:response noContentResponse
type noContentResponse struct {
}
