package dataservice

import (
	"context"
	"net/url"

	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Database is an interface representing data base service with all CRUD operations
type Database interface {
	Shutdown(context.Context)
	GetEvents(queryParams url.Values) (*model.Events, error)
	GetEvent(id primitive.ObjectID) (*model.Event, error)
	AddEvent(event *model.Event) error
	UpdateEvent(id primitive.ObjectID, event *model.Event) error
	DeleteEvent(id primitive.ObjectID) error
}

// NewDatabase is a factory method for creating database service according to configuration
func NewDatabase(config *utils.Config) (Database, error) {
	filters := NewFilters(&config.Server)
	if config.Server.TypeDB == "mongo" {
		return NewMongoDB(&config.MongoDB, filters)
	}
	return NewInMemory(filters)
}
