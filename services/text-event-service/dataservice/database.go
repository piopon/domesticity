package dataservice

import (
	"net/url"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

// Database is an interface representing data base service with all CRUD operations
type Database interface {
	GetEvents(queryParams url.Values) (*model.Events, error)
	GetEvent(id int) (*model.Event, error)
	AddEvent(event *model.Event)
	UpdateEvent(id int, event *model.Event) error
	DeleteEvent(id int) error
}
