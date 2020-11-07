package dataservice

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB is a data base service with elements are stored with use of MongoDB
type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB is a factory method to create a Mongo DB service
func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func initMongoClient(URI string) (*mongo.Client, error) {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, error := mongo.Connect(context, options.Client().ApplyURI(URI))
	if error != nil {
		return nil, fmt.Errorf("error while creating a MongoDB client [" + URI + "]")
	}
	defer func() {
		if error := client.Disconnect(context); error != nil {
			panic(error)
		}
	}()
	if error := client.Ping(context, readpref.Primary()); error != nil {
		return nil, fmt.Errorf("error while connecting to MongoDB server [" + URI + "]")
	}
	return client, nil
}

// GetEvents returns all events stored in DB
func (mongo MongoDB) GetEvents(queryParams url.Values) (*model.Events, error) {
	return nil, nil
}

// GetEvent returns event with specified ID (or error if not found)
func (mongo MongoDB) GetEvent(id int) (*model.Event, error) {
	return nil, nil
}

// AddEvent adds passed event item to DB
func (mongo MongoDB) AddEvent(event *model.Event) {
	return
}

// UpdateEvent updates an event with specified ID
func (mongo MongoDB) UpdateEvent(id int, event *model.Event) error {
	return nil
}

// DeleteEvent deletes a event with specified ID from the database
func (mongo MongoDB) DeleteEvent(id int) error {
	return nil
}
