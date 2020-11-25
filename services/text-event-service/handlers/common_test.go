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
	TestLogger      *log.Logger
	TestValidator   *utils.Validator
	TestDataservice dataservice.Database
}

func NewCommonMockup() *commonMockup {
	testConfig := utils.ConfigServer{TypeDB: "memory"}
	testFilters := dataservice.NewFilters(&testConfig)
	testDatabase, _ := dataservice.NewInMemory(testFilters)
	return &commonMockup{
		TestLogger:      log.New(os.Stdout, "TEST > ", log.LstdFlags),
		TestValidator:   utils.NewValidator(),
		TestDataservice: testDatabase,
	}
}

func (common *commonMockup) CreateEventsHandler() *handlers.Events {
	return handlers.NewEvents(common.TestLogger, common.TestValidator, common.TestDataservice)
}

func (common *commonMockup) GetDatabaseIds() []primitive.ObjectID {
	events, errors := common.TestDataservice.GetEvents(nil)
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
	common.TestDataservice.AddEvent(newEvent)
	return newEvent.ID
}
