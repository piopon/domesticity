package dataservice

import (
	"net/url"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

// MongoDB is a data base service with elements are stored with use of MongoDB
type MongoDB struct{}

// NewMongoDB is a factory method to create a Mongo DB service
func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

// GetEvents returns all events stored in DB
func (mongo MongoDB) GetEvents(queryParams url.Values) (*model.Events, error) {
	return nil, nil
}

// GetEvent returns event with specified ID (or error if not found)
func (mongo MongoDB) GetEvent(id int) (*model.Event, error) {
	return nil, nil
}

// AddEvent adds passed event item to DB
func (mongo MongoDB) AddEvent(event *model.Event) {
	return
}

// UpdateEvent updates an event with specified ID
func (mongo MongoDB) UpdateEvent(id int, event *model.Event) error {
	return nil
}

// DeleteEvent deletes a event with specified ID from the database
func (mongo MongoDB) DeleteEvent(id int) error {
	return nil
}
