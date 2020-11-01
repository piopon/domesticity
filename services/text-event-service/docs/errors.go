package docs

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
