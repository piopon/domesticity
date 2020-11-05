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

// ResponsePostEvent returns an event which was added to DB
// swagger:response responsePostEvent
type responsePostEvent struct {
	// Currently added event to DB
	// in: body
	Body model.Event
}

// ResponsePutEvent returns an event which was updated in DB (depends on ID parameter)
// swagger:response responsePutEvent
type responsePutEvent struct {
	// Currently updated event in DB (selected by ID parameter)
	// in: body
	Body model.Event
}

// ResponseNoContent is returned when no specific response is needed
// swagger:response responseNoContent
type responseNoContent struct {
}

// ResponseDocumentation returns a HTML file with documentation
// swagger:response responseDocumentation
type responseDocumentation struct {
	// A text/html content with documentation
	// in: body
	HTML string
}

// ResponseSwagger returns a YAML file with swagger configuration
// swagger:response responseSwagger
type responseSwagger struct {
	// A text/ content with swagger.yaml configuration file
	// in: body
	YAML byte
}
