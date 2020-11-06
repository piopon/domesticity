package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/model"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

// Events is a service handler responsible for getting and updating events
type Events struct {
	logger    *log.Logger
	validator *utils.Validator
	database  dataservice.Database
}

// KeyEvent is a key used for add and get the Event object in the context
type KeyEvent struct{}

// NewEvents is a factory method to create Events service handler with defined logger
func NewEvents(logger *log.Logger, validator *utils.Validator, database dataservice.Database) *Events {
	return &Events{logger, validator, database}
}

// ValidationMiddleware is used to parse and validate Event from request
func (events *Events) ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		event := &model.Event{}
		error := utils.FromJSON(event, request.Body)
		if error != nil {
			events.logger.Println("Unable to unmarshal events data")
			response.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&model.GenericError{"Cannot deserialize inputted Event json"}, response)
			return
		}
		validationErrors := events.validator.Validate(event)
		if validationErrors != nil {
			events.logger.Println("Unable to validate events data")
			response.WriteHeader(http.StatusUnprocessableEntity)
			utils.ToJSON(&model.ValidationError{validationErrors.Errors()}, response)
			return
		}
		// add event to the context and call next handler (other middleware or final handler)
		ctx := context.WithValue(request.Context(), KeyEvent{}, event)
		request = request.WithContext(ctx)
		next.ServeHTTP(response, request)
	})
}
