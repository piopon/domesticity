package handlers_test

import (
	"log"
	"os"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/handlers"
	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commonMockup struct {
	testLogger      *log.Logger
	testConfig      *utils.Config
	testValidator   *utils.Validator
	testDataservice dataservice.Database
}

func NewCommonMockup() *commonMockup {
	testConfig := utils.ConfigServer{TypeDB: "memory"}
	testFilters := dataservice.NewFilters(&testConfig)
	testDatabase, _ := dataservice.NewInMemory(testFilters)
	return &commonMockup{
		testLogger:      log.New(os.Stdout, "TEST > ", log.LstdFlags),
		testConfig:      utils.NewConfig("../../resources"),
		testValidator:   utils.NewValidator(),
		testDataservice: testDatabase,
	}
}

func (common *commonMockup) CreateEventsHandler() *handlers.Events {
	return handlers.NewEvents(common.testLogger, common.testValidator, common.testDataservice)
}

func (common *commonMockup) CreateHomeHandler(template string) *handlers.Home {
	return handlers.NewHome(template, common.testLogger, common.testConfig)
}

func (common *commonMockup) CreateEventBadJson() *model.Event {
	return &model.Event{
		Title: "This is my first event",
		Icon:  "my-icon-1",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(-2020, 05, 26, 14, 15, 00, 00, time.Local),
			Stop:  time.Date(-2020, 05, 27, 10, 30, 00, 00, time.Local)},
		Category: "Notes",
		Content:  "Test event number 1",
	}
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
	newEvent := common.CreateEventBadJson()
	common.testDataservice.AddEvent(newEvent)
	return newEvent.ID
}
