package docs

import (
	"time"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

// swagger:parameters getEvent updateEvent deleteEvent
type paramID struct {
	// The id of the event for which the operation relates<br>
	// NOTE: The type is primitive.ObjectID which is the BSON ObjectID type
	// in: path
	// required: true
	ID []int `json:"id"`
}

// swagger:parameters getEvents
type paramGetEventsFilter struct {
	// The number of items to return.
	// in: query
	// minimum: 0
	Limit int `json:"limit"`
	// The number of items to skip before starting to collect the result set.
	// in: query
	// minimum: 0
	Offset int `json:"offset"`
	// The string which has to be included (not exact match) in searched event title
	// in: query
	Title string `json:"title"`
	// The string which has to match exact searched event owner
	// in: query
	Owner string `json:"owner"`
	// The start date of seached event in format YYYY-MM-DD
	// in: query
	DayStart time.Time `json:"dayStart"`
	// The end date of seached event in format YYYY-MM-DD
	// in: query
	DayStop time.Time `json:"dayStop"`
	// The string which has to match exact searched event category
	// in: query
	Category string `json:"category"`
	// The string which has to be included (not exact match) in searched event content
	// in: query
	Content string `json:"content"`
}

// swagger:parameters addEvent updateEvent
type paramEvent struct {
	// Event data structure to update or create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body model.Event
}
