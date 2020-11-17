package dataservice

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"github.com/piopon/domesticity/services/text-event-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB is a data base service with elements are stored with use of MongoDB
type MongoDB struct {
	client   *mongo.Client
	filters  *Filters
	document *mongo.Collection
	timeouts *utils.ConfigMongoTimeout
}

// NewMongoDB is a factory method to create a Mongo DB service
func NewMongoDB(config *utils.ConfigMongo, filters *Filters) *MongoDB {
	mongoClient, error := initMongoClient(config)
	if error != nil {
		panic("Cannot initialize MongoDB client. " + error.Error())
	}
	return &MongoDB{
		client:   mongoClient,
		filters:  filters,
		document: mongoClient.Database(config.Database.Name).Collection(config.Database.Collection),
		timeouts: &config.Timeout,
	}
}

func initMongoClient(config *utils.ConfigMongo) (*mongo.Client, error) {
	URI := config.Scheme + config.IP + ":" + config.Port
	context, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout.Connection)*time.Second)
	defer cancel()
	client, error := mongo.Connect(context, options.Client().ApplyURI(URI))
	if error != nil {
		return nil, fmt.Errorf("Error while creating a MongoDB client: "+URI, error.Error())
	}
	if error := client.Ping(context, readpref.Primary()); error != nil {
		return nil, fmt.Errorf("Error while connecting to MongoDB server: "+URI, error.Error())
	}
	return client, nil
}

// Shutdown closes active database connection
func (mongo MongoDB) Shutdown(ctx context.Context) {
	if error := mongo.client.Disconnect(ctx); error != nil {
		panic("Cannot disconnect from MongoDB client: " + error.Error())
	}
}

// GetEvents returns all events stored in DB
func (mongo MongoDB) GetEvents(queryParams url.Values) (*model.Events, error) {
	var events model.Events
	context, cancel := context.WithTimeout(context.Background(), time.Duration(mongo.timeouts.Get)*time.Second)
	defer cancel()
	findOptions, error := mongo.filters.GetOptions(queryParams)
	if error != nil {
		return nil, fmt.Errorf("Cannot get query params: " + error.Error())
	}
	findFilters, error := mongo.filters.GetFilters(queryParams)
	if error != nil {
		return nil, fmt.Errorf("Cannot get query params: " + error.Error())
	}
	cursor, error := mongo.document.Find(context, findFilters, findOptions)
	if error != nil {
		return nil, fmt.Errorf("Cannot find elements: " + error.Error())
	}
	defer cursor.Close(context)
	for cursor.Next(context) {
		var event model.Event
		cursor.Decode(&event)
		events = append(events, &event)
	}
	if error := cursor.Err(); error != nil {
		return nil, error
	}
	return &events, nil
}

// GetEvent returns event with specified ID (or error if not found)
func (mongo MongoDB) GetEvent(id primitive.ObjectID) (*model.Event, error) {
	var event model.Event
	context, cancel := context.WithTimeout(context.Background(), time.Duration(mongo.timeouts.Get)*time.Second)
	defer cancel()
	error := mongo.document.FindOne(context, bson.M{"_id": id}).Decode(&event)
	if error != nil {
		return nil, error
	}
	return &event, nil
}

// AddEvent adds passed event item to DB
func (mongo MongoDB) AddEvent(event *model.Event) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(mongo.timeouts.Post)*time.Second)
	defer cancel()
	newID, error := mongo.document.InsertOne(context, event)
	if error == nil {
		event.ID = newID.InsertedID.(primitive.ObjectID)
	}
	return error
}

// UpdateEvent updates an event with specified ID
func (mongo MongoDB) UpdateEvent(id primitive.ObjectID, event *model.Event) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(mongo.timeouts.Put)*time.Second)
	defer cancel()
	_, error := mongo.document.UpdateOne(context, bson.M{"_id": id}, bson.M{"$set": event})
	if error == nil {
		event.ID = id
	}
	return error
}

// DeleteEvent deletes a event with specified ID from the database
func (mongo MongoDB) DeleteEvent(id primitive.ObjectID) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(mongo.timeouts.Delete)*time.Second)
	defer cancel()
	_, error := mongo.document.DeleteOne(context, bson.M{"_id": id})
	return error
}
