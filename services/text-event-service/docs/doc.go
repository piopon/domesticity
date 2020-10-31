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
