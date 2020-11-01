// Package docs Domesticity Text Event Service API
//
// Documentation for Text Event Service API used by Domesticity application
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// License: GPL-3.0 https://opensource.org/licenses/GPL-3.0
// Contact: Piotr Ponikowski <piopon.github@gmail.com>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Security:
// - basic
//
// swagger:meta
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

// Error message returned as a string when bad request was invoked (400)
// swagger:response errorBadQuery
type errorBadQuery struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body string
}

// Error message returned as a string when desired item is not found (404)
// swagger:response errorNotFound
type errorNotFound struct {
	// Error related to non-existent data to fetch
	// in: body
	Body string
}

// Error message returned as a string when internal server error occurs (500)
// swagger:response errorInternal
type errorInternal struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body string
}

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
