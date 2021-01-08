package dataservice_test

import (
	"net/url"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewMongoDbReturnsErrorWhenNoConnection(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	mongo, error := dataservice.NewMongoDB(&config, nil)
	if error == nil {
		t.Error("MongoDB factory method not returned any error when connection not present")
	}
	if mongo == nil {
		t.Error("MongoDB factory method returned nil object")
	}
}

func TestShutdownPanicsWhenNoConnection(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	mongo, _ := dataservice.NewMongoDB(&config, nil)
	assert.Panics(t, func() { mongo.Shutdown(nil) }, "The code did not panic")
}

func TestGetEventsPanicsWhenNoConnection(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	filters := dataservice.NewFilters(&utils.ConfigServer{TypeDB: "mongo"})
	mongo, _ := dataservice.NewMongoDB(&config, filters)
	queryParams := url.Values{}
	assert.Panics(t, func() { mongo.GetEvents(queryParams) }, "The code did not panic")
}

func TestGetEventsFailsWhenOptionsAreIncorrect(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	filters := dataservice.NewFilters(&utils.ConfigServer{TypeDB: "mongo"})
	mongo, _ := dataservice.NewMongoDB(&config, filters)
	queryParams := url.Values{
		"none":  []string{"zero"},
		"owner": []string{"Admin"},
	}
	events, error := mongo.GetEvents(queryParams)
	if error == nil {
		t.Error("Get events should return error if query params is incorrect")
	}
	if events != nil {
		t.Error("Get events should return events = nil but didn't")
	}
}

func TestGetEventsFailsWhenFiltersAreIncorrect(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	filters := dataservice.NewFilters(&utils.ConfigServer{TypeDB: "mongo"})
	mongo, _ := dataservice.NewMongoDB(&config, filters)
	queryParams := url.Values{
		"limit": []string{"10"},
		"none":  []string{"zero"},
	}
	events, error := mongo.GetEvents(queryParams)
	if error == nil {
		t.Error("Get events should return error if query params is incorrect")
	}
	if events != nil {
		t.Error("Get events should return events = nil but didn't")
	}
}

func TestGetEventPanicsWhenNoConnection(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	filters := dataservice.NewFilters(&utils.ConfigServer{TypeDB: "mongo"})
	mongo, _ := dataservice.NewMongoDB(&config, filters)
	assert.Panics(t, func() { mongo.GetEvent(primitive.NilObjectID) }, "The code did not panic")
}

func TestAddEventPanicsWhenNoConnection(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	filters := dataservice.NewFilters(&utils.ConfigServer{TypeDB: "mongo"})
	mongo, _ := dataservice.NewMongoDB(&config, filters)
	newEvent := model.Event{}
	assert.Panics(t, func() { mongo.AddEvent(&newEvent) }, "The code did not panic")
}

func TestUpdateEventPanicsWhenNoConnection(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	filters := dataservice.NewFilters(&utils.ConfigServer{TypeDB: "mongo"})
	mongo, _ := dataservice.NewMongoDB(&config, filters)
	newEvent := model.Event{}
	assert.Panics(t, func() { mongo.UpdateEvent(primitive.NilObjectID, &newEvent) }, "The code did not panic")
}

func TestDeleteEventPanicsWhenNoConnection(t *testing.T) {
	config := utils.ConfigMongo{IP: "localhost", Port: "27017", Scheme: "mongodb://"}
	filters := dataservice.NewFilters(&utils.ConfigServer{TypeDB: "mongo"})
	mongo, _ := dataservice.NewMongoDB(&config, filters)
	assert.Panics(t, func() { mongo.DeleteEvent(primitive.NilObjectID) }, "The code did not panic")
}
