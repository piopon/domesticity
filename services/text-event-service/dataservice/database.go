package dataservice

import (
	"context"
	"net/url"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Database is an interface representing data base service with all CRUD operations
type Database interface {
	Shutdown(context.Context)
	GetEvents(queryParams url.Values) (*model.Events, error)
	GetEvent(id primitive.ObjectID) (*model.Event, error)
	AddEvent(event *model.Event)
	UpdateEvent(id primitive.ObjectID, event *model.Event) error
	DeleteEvent(id primitive.ObjectID) error
}
