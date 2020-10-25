package data

import (
	"fmt"
)

// Event defines the structure for an API event
type Event struct {
	ID       int    `json:"id"`
	Title    string `json:"title" validate:"required"`
	Owner    string `json:"owner" validate:"required"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

// Events is a type definition for slice of Event pointers
type Events []*Event

// GetEvents returns all events stored in DB
func GetEvents() Events {
	return eventList
}

// GetEventByID returns event with specified ID (or error if not found)
func GetEventByID(id int) (*Event, error) {
	index, error := findEvent(id)
	if error != nil {
		return nil, error
	}
	return eventList[index], nil
}

// AddEvent adds passed event item to DB
func AddEvent(event *Event) {
	event.ID = getNextID()
	eventList = append(eventList, event)
}

// UpdateEvent updates an event with specified ID
func UpdateEvent(id int, event *Event) error {
	index, error := findEvent(id)
	if error != nil {
		return error
	}
	event.ID = id
	eventList[index] = event
	return nil
}

// DeleteEvent deletes a event with specified ID from the database
func DeleteEvent(id int) error {
	index, error := findEvent(id)
	if error != nil {
		return error
	}
	eventList = append(eventList[:index], eventList[index+1])
	return nil
}

func findEvent(id int) (int, error) {
	for i, event := range eventList {
		if event.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Event not found")
}

func getNextID() int {
	return eventList[len(eventList)-1].ID + 1
}

var eventList = Events{
	&Event{
		ID:       1,
		Title:    "This is my first event",
		Owner:    "Admin",
		Category: "Notes",
		Content:  "Test event number 1",
	},
	&Event{
		ID:       2,
		Title:    "2nd event",
		Owner:    "Admin",
		Category: "Stuff",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	},
}
