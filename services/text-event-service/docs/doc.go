// Package docs Domesticity Text Event Service API
// NOTE: Types defined here are purely for documentation purposes
//
// Documentation for Text Event Service API used by Domesticity application
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// License: MIT https://opensource.org/licenses/GPL-3.0
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
