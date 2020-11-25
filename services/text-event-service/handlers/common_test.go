package handlers_test

import (
	"log"
	"os"

	"github.com/piopon/domesticity/services/text-event-service/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/handlers"
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
