package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/piopon/domesticity/services/text-event-service/data"
)

// Events is a service handler responsible for getting and updating events
type Events struct {
	logger *log.Logger
}

// KeyEvent is a key used for add and get the Event object in the context
type KeyEvent struct{}

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
	event := request.Context().Value(KeyEvent{}).(data.Event)
	data.AddEvent(&event)
}

// UpdateEvent is used to update event with specified ID stored in DB
func (events *Events) UpdateEvent(response http.ResponseWriter, request *http.Request) {
	events.logger.Println("Handling PUT event")
	id, error := strconv.Atoi(mux.Vars(request)["id"])
	if error != nil {
		http.Error(response, "Bad URL", http.StatusBadRequest)
	}
	event := request.Context().Value(KeyEvent{}).(data.Event)
	updateError := data.UpdateEvent(id, &event)
	if updateError != nil {
		events.logger.Println("Invalid event ID")
		return
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

// ValidationMiddleware is used to parse and validate Event from request
func (events *Events) ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		event := data.Event{}
		error := event.FromJSON(request.Body)
		if error != nil {
			events.logger.Println("Unable to unmarshal events data")
			return
		}
		// add event to the context and call next handler (other middleware or final handler)
		ctx := context.WithValue(request.Context(), KeyEvent{}, event)
		request = request.WithContext(ctx)
		next.ServeHTTP(response, request)
	})
}
