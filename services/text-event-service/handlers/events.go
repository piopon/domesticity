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

func (events *Events) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if http.MethodGet == request.Method {
		events.getEvents(response, request)
		return
	}
	if http.MethodPost == request.Method {
		events.addEvent(response, request)
		return
	}
	if http.MethodPut == request.Method {
		id := events.parseID(request.URL.Path)
		events.updateEvent(id, response, request)
		return
	}
	response.WriteHeader(http.StatusMethodNotAllowed)
}

func (events *Events) getEvents(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling GET events")
	allEvents := data.GetEvents()
	error := allEvents.ToJSON(response)
	if error != nil {
		events.logger.Println("Unable to marshal events data")
	}
}

func (events *Events) addEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling POST event")
	event := &data.Event{}
	error := event.FromJSON(request.Body)
	if error != nil {
		events.logger.Println("Unable to unmarshal events data")
	}
	data.AddEvent(event)
}

func (events *Events) updateEvent(id int, response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling POST event")
	event := &data.Event{}
	error := event.FromJSON(request.Body)
	if error != nil {
		events.logger.Println("Unable to unmarshal events data")
	}
	data.UpdateEvent(id, event)
}

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
