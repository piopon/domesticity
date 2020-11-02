package docs

import "github.com/piopon/domesticity/services/text-event-service/model"

// Error message returned as a string when bad request was invoked (400)
// swagger:response errorBadQuery
type errorBadQuery struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body model.GenericError
}

// Error message returned as a string when desired item is not found (404)
// swagger:response errorNotFound
type errorNotFound struct {
	// Error related to non-existent data to fetch
	// in: body
	Body model.GenericError
}

// Error message returned as a string when internal server error occurs (500)
// swagger:response errorInternal
type errorInternal struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body model.GenericError
}

// Error message returned as a string when validation error occurs (422)
// swagger:response errorValidation
type errorValidation struct {
	// Error related to bad query parameters (bad filter or value)
	// in: body
	Body model.ValidationError
}
