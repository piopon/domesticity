package docs

import "github.com/piopon/domesticity/services/text-event-service/model"

// ResponseEvents with all or filtered events from DB (depends on query params)
// swagger:response responseEvents
type responseEvents struct {
	// All / filtered events from DB
	// in: body
	Body model.Events
}

// ResponseEvent with single event from DB (depends on ID parameter)
// swagger:response responseEvent
type responseEvent struct {
	// Single event from DB
	// in: body
	Body model.Event
}

// ResponseNoContent is returned when no specific response is needed
// swagger:response responseNoContent
type responseNoContent struct {
}
