package handlers_test

import (
	"log"
	"os"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/handlers"
	"github.com/piopon/domesticity/services/text-event-service/model"
	"github.com/piopon/domesticity/services/text-event-service/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commonMockup struct {
	testLogger      *log.Logger
	testValidator   *utils.Validator
	testDataservice dataservice.Database
}

func NewCommonMockup() *commonMockup {
	testConfig := utils.ConfigServer{TypeDB: "memory"}
	testFilters := dataservice.NewFilters(&testConfig)
	testDatabase, _ := dataservice.NewInMemory(testFilters)
	return &commonMockup{
		testLogger:      log.New(os.Stdout, "TEST > ", log.LstdFlags),
		testValidator:   utils.NewValidator(),
		testDataservice: testDatabase,
	}
}

func (common *commonMockup) CreateEventsHandler() *handlers.Events {
	return handlers.NewEvents(common.testLogger, common.testValidator, common.testDataservice)
}

func (common *commonMockup) GetDatabaseIds() []primitive.ObjectID {
	events, errors := common.testDataservice.GetEvents(nil)
	if errors != nil {
		return nil
	}
	result := []primitive.ObjectID{}
	for _, event := range *events {
		result = append(result, event.ID)
	}
	return result
}

func (common *commonMockup) AddBadEventToDB() primitive.ObjectID {
	newEvent := &model.Event{
		Title: "This is my first event",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(-2020, 05, 26, 14, 15, 00, 00, time.Local),
			Stop:  time.Date(-2020, 05, 27, 10, 30, 00, 00, time.Local)},
		Category: "Notes",
		Content:  "Test event number 1",
	}
	common.testDataservice.AddEvent(newEvent)
	return newEvent.ID
}
