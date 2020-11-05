package docs

import "github.com/piopon/domesticity/services/text-event-service/model"

// Response with all or filtered events from DB (filter depends on query params)
// swagger:response responseGetEvents
type responseGetEvents struct {
	// All / filtered events from DB
	// in: body
	Body model.Events
}

// Response with single event from DB (selected by ID parameter in path param)
// swagger:response responseGetEvent
type responseGetEvent struct {
	// Single event from DB (selected by ID parameter)
	// in: body
	Body model.Event
}

// Response with currently created event which was added to DB
// swagger:response responsePostEvent
type responsePostEvent struct {
	// Currently added event to DB
	// in: body
	Body model.Event
}

// Response with currently updated event stored in DB (selected by ID parameter in path param)
// swagger:response responsePutEvent
type responsePutEvent struct {
	// Currently updated event in DB (selected by ID parameter)
	// in: body
	Body model.Event
}

// Response with no specific content in body (status indicates success)
// swagger:response responseNoContent
type responseNoContent struct {
}

// Response with a HTML documentation file in body
// swagger:response responseDocumentation
type responseDocumentation struct {
	// A text/html content with documentation
	// in: body
	HTML string
}

// Response with a YAML swagger configuration file in body
// swagger:response responseSwagger
type responseSwagger struct {
	// A text/ content with swagger.yaml configuration file
	// in: body
	YAML byte
}
