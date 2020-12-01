package docs

import "github.com/piopon/domesticity/services/text-event-service/model"

// Error with JSON message returned in body when bad request was invoked
// swagger:response errorBadQuery
type errorBadQuery struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body model.GenericError
}

// Error with JSON message returned in body when desired item is not found
// swagger:response errorNotFound
type errorNotFound struct {
	// Error related to non-existent data to fetch
	// in: body
	Body model.GenericError
}

// Error with JSON message returned in body when internal server error occurs
// swagger:response errorInternal
type errorInternal struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body model.GenericError
}

// Error with JSON messages returned in body when validation error occurs
// swagger:response errorValidation
type errorValidation struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body model.ValidationError
}

// Error with no specific content returned when no swagger.yaml is found
// swagger:response errorSwagger
type errorSwagger struct {
}
