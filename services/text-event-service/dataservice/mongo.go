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
	client         *mongo.Client
	nameDatabase   string
	nameCollection string
}

// NewMongoDB is a factory method to create a Mongo DB service
func NewMongoDB(config *utils.ConfigMongo) *MongoDB {
	mongoClient, error := initMongoClient(config.Scheme + config.IP + ":" + config.Port)
	if error != nil {
		panic("Cannot initialize MongoDB client. " + error.Error())
	}
	return &MongoDB{mongoClient, "event-service", "events"}
}

func initMongoClient(URI string) (*mongo.Client, error) {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	collection := mongo.client.Database(mongo.nameDatabase).Collection(mongo.nameCollection)
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, error := collection.Find(context, bson.M{})
	if error != nil {
		return nil, error
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
	collection := mongo.client.Database(mongo.nameDatabase).Collection(mongo.nameCollection)
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	error := collection.FindOne(context, bson.M{"_id": id}).Decode(&event)
	if error != nil {
		return nil, error
	}
	return &event, nil
}

// AddEvent adds passed event item to DB
func (mongo MongoDB) AddEvent(event *model.Event) error {
	collection := mongo.client.Database(mongo.nameDatabase).Collection(mongo.nameCollection)
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newID, error := collection.InsertOne(context, event)
	if error == nil {
		event.ID = newID.InsertedID.(primitive.ObjectID)
	}
	return error
}

// UpdateEvent updates an event with specified ID
func (mongo MongoDB) UpdateEvent(id primitive.ObjectID, event *model.Event) error {
	collection := mongo.client.Database(mongo.nameDatabase).Collection(mongo.nameCollection)
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, error := collection.UpdateOne(context, bson.M{"_id": id}, bson.M{"$set": event})
	if error == nil {
		event.ID = id
	}
	return error
}

// DeleteEvent deletes a event with specified ID from the database
func (mongo MongoDB) DeleteEvent(id primitive.ObjectID) error {
	collection := mongo.client.Database(mongo.nameDatabase).Collection(mongo.nameCollection)
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, error := collection.DeleteOne(context, bson.M{"_id": id})
	return error
}
