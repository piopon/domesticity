package model

// Event defines the structure for an API event
// swagger:model
type Event struct {
	// The id for the event
	//
	// required: false
	// min: 1
	ID int `json:"id"`
	// The title of event
	//
	// required: true
	// max length: 255
	Title string `json:"title" validate:"required"`
	// The title of event
	//
	// required: true
	// max length: 255
	Owner string `json:"owner" validate:"required"`
	// The time span of event
	//
	// required: true
	Occurence TimeSpan `json:"date"  validate:"required"`
	// The category of event
	//
	// required: true
	// max length: 255
	Category string `json:"category"`
	// The event main content (text message)
	//
	// required: false
	// max length: 10000
	Content string `json:"content"`
}
