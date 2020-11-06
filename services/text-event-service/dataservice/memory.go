package dataservice

import (
	"fmt"
	"net/url"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

// InMemory is a test data base service with elements stored in RAM
type InMemory struct{}

// NewInMemory is a factory method to create an in memory data base service
func NewInMemory() *InMemory {
	return &InMemory{}
}

// GetEvents returns all events stored in DB
func (memory InMemory) GetEvents(queryParams url.Values) (*model.Events, error) {
	return eventList.Filter(queryParams)
}

// GetEvent returns event with specified ID (or error if not found)
func (memory InMemory) GetEvent(id int) (*model.Event, error) {
	index, error := memory.findEvent(id)
	if error != nil {
		return nil, error
	}
	return eventList[index], nil
}

// AddEvent adds passed event item to DB
func (memory InMemory) AddEvent(event *model.Event) {
	event.ID = memory.getNextID()
	eventList = append(eventList, event)
}

// UpdateEvent updates an event with specified ID
func (memory InMemory) UpdateEvent(id int, event *model.Event) error {
	index, error := memory.findEvent(id)
	if error != nil {
		return error
	}
	event.ID = id
	eventList[index] = event
	return nil
}

// DeleteEvent deletes a event with specified ID from the database
func (memory InMemory) DeleteEvent(id int) error {
	index, error := memory.findEvent(id)
	if error != nil {
		return error
	}
	eventList = append(eventList[:index], eventList[index+1:]...)
	return nil
}

func (memory InMemory) findEvent(id int) (int, error) {
	for i, event := range eventList {
		if event.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Event not found")
}

func (memory InMemory) getNextID() int {
	return eventList[len(eventList)-1].ID + 1
}

var eventList = model.Events{
	&model.Event{
		ID:    1,
		Title: "This is my first event",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(2020, 01, 02, 12, 12, 00, 00, time.Local),
			Stop:  time.Date(2020, 01, 02, 12, 32, 00, 00, time.Local)},
		Category: "Notes",
		Content:  "Test event number 1",
	},
	&model.Event{
		ID:    2,
		Title: "2nd event",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(2020, 07, 13, 12, 12, 00, 00, time.Local),
			Stop:  time.Date(2020, 07, 13, 12, 32, 00, 00, time.Local)},
		Category: "Stuff",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	},
}
