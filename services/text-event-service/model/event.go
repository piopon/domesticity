package model

import (
	"fmt"
	"net/url"
	"time"
)

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

// GetEvents returns all events stored in DB
func GetEvents(queryParams url.Values) (*Events, error) {
	return eventList.Filter(queryParams)
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
	eventList = append(eventList[:index], eventList[index+1:]...)
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
		ID:    1,
		Title: "This is my first event",
		Owner: "Admin",
		Occurence: TimeSpan{
			Start: time.Date(2020, 01, 02, 12, 12, 00, 00, time.Local),
			Stop:  time.Date(2020, 01, 02, 12, 32, 00, 00, time.Local)},
		Category: "Notes",
		Content:  "Test event number 1",
	},
	&Event{
		ID:    2,
		Title: "2nd event",
		Owner: "Admin",
		Occurence: TimeSpan{
			Start: time.Date(2020, 07, 13, 12, 12, 00, 00, time.Local),
			Stop:  time.Date(2020, 07, 13, 12, 32, 00, 00, time.Local)},
		Category: "Stuff",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	},
}
