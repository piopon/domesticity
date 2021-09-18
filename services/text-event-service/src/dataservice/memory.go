package dataservice

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InMemory is a test data base service with elements stored in RAM
type InMemory struct {
	filters    *Filters
	eventsList model.Events
}

// NewInMemory is a factory method to create an in memory data base service
func NewInMemory(filters *Filters) (*InMemory, error) {
	return &InMemory{
		filters: filters,
		eventsList: model.Events{
			&model.Event{
				ID:    primitive.NewObjectID(),
				Title: "This is my first event",
				Icon:  "my-icon-1",
				Owner: "Admin",
				Occurence: model.TimeSpan{
					Start: time.Date(2020, 05, 26, 14, 15, 00, 00, time.Local),
					Stop:  time.Date(2020, 05, 27, 10, 30, 00, 00, time.Local)},
				Category: "Notes",
				Content:  "Test event number 1",
			},
		},
	}, nil
}

// Shutdown clears internal events list
func (memory *InMemory) Shutdown(ctx context.Context) {
	memory.eventsList = nil
}

// GetEvents returns all events stored in DB
func (memory *InMemory) GetEvents(queryParams url.Values) (*model.Events, error) {
	internalFilers := url.Values{"internal": {""}}
	filter, err := memory.filters.GetFilters(internalFilers)
	if err != nil {
		return nil, err
	}
	return filter.(func(model.Events, url.Values) (*model.Events, error))(memory.eventsList, queryParams)
}

// GetEvent returns event with specified ID (or error if not found)
func (memory *InMemory) GetEvent(id primitive.ObjectID) (*model.Event, error) {
	index, error := memory.findEvent(id)
	if error != nil {
		return nil, error
	}
	return memory.eventsList[index], nil
}

// AddEvent adds passed event item to DB
func (memory *InMemory) AddEvent(event *model.Event) error {
	_, error := memory.findEvent(event.ID)
	if error == nil {
		return fmt.Errorf("event with specified ID already exists")
	}
	event.ID = primitive.NewObjectID()
	memory.eventsList = append(memory.eventsList, event)
	return nil
}

// UpdateEvent updates an event with specified ID
func (memory *InMemory) UpdateEvent(id primitive.ObjectID, event *model.Event) error {
	index, error := memory.findEvent(id)
	if error != nil {
		return error
	}
	event.ID = id
	memory.eventsList[index] = event
	return nil
}

// DeleteEvent deletes a event with specified ID from the database
func (memory *InMemory) DeleteEvent(id primitive.ObjectID) error {
	index, error := memory.findEvent(id)
	if error != nil {
		return error
	}
	memory.eventsList = append(memory.eventsList[:index], memory.eventsList[index+1:]...)
	return nil
}

func (memory *InMemory) findEvent(id primitive.ObjectID) (int, error) {
	for i, event := range memory.eventsList {
		if event.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("event not found")
}
