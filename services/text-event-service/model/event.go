package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Event defines the structure for an API event
// swagger:model
type Event struct {
	// The id for the event
	//
	// required: false
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	// The title of event
	//
	// required: true
	// max length: 255
	Title string `json:"title" bson:"title" validate:"required"`
	// The title of event
	//
	// required: true
	// max length: 255
	Owner string `json:"owner" bson:"owner" validate:"required"`
	// The time span of event
	//
	// required: true
	Occurence TimeSpan `json:"date" bson:"date" validate:"required"`
	// The category of event
	//
	// required: true
	// max length: 255
	Category string `json:"category" bson:"category"`
	// The event main content (text message)
	//
	// required: false
	// max length: 10000
	Content string `json:"content" bson:"content"`
}
