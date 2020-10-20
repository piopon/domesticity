package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/piopon/domesticity/services/text-event-service/data"
)

// Events is a service handler responsible for getting and updating events
type Events struct {
	logger *log.Logger
}

// NewEvents is a factory method to create Events service handler with defined logger
func NewEvents(logger *log.Logger) *Events {
	return &Events{logger}
}

// GetEvents is used to retrieve all currently stored events
func (events *Events) GetEvents(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET events")
	allEvents := data.GetEvents()
	error := allEvents.ToJSON(response)
	if error != nil {
		events.logger.Println("Unable to marshal events data")
	}
}

// AddEvent is used to add new event and store it in DB
func (events *Events) AddEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling POST event")
	event, error := events.parseEvent(request)
	if error != nil {
		http.Error(response, "Bad URL", http.StatusBadRequest)
	}
	data.AddEvent(event)
}

// UpdateEvent is used to update event with specified ID stored in DB
func (events *Events) UpdateEvent(id int, response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling POST event")
	event, error := events.parseEvent(request)
	if error != nil {
		http.Error(response, "Bad URL", http.StatusBadRequest)
	}
	updateError := data.UpdateEvent(id, event)
	if updateError != nil {
		events.logger.Println("Invalid event ID")
		return
	}
}

func (events *Events) parseEvent(request *http.Request) (*data.Event, error) {
	event := &data.Event{}
	error := event.FromJSON(request.Body)
	if error != nil {
		events.logger.Println("Unable to unmarshal events data")
		return nil, error
	}
	return event, nil
}

func (events *Events) parseID(urlPath string) int {
	regex := regexp.MustCompile("/([0-9]+)")
	found := regex.FindAllStringSubmatch(urlPath, -1)
	if len(found) != 1 || len(found[0]) != 2 {
		events.logger.Println("Invalid URI inputted", urlPath)
		return -1
	}
	id, error := strconv.Atoi(found[0][1])
	if error != nil {
		events.logger.Println("Invalid URI inputted", urlPath)
		return -1
	}
	return id
}
