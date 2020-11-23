package dataservice_test

import (
	"fmt"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

func TestNewDatabaseReturnsInMemoryAccordingToConfig(t *testing.T) {
	config := utils.Config{
		Server: utils.ConfigServer{TypeDB: "memory"},
	}
	db, dbError := dataservice.NewDatabase(&config)
	dbType := fmt.Sprintf("%T", db)
	if dbType != "*dataservice.InMemory" {
		t.Errorf("Bad DB concrete type created %s", dbType)
	}
	if dbError != nil {
		t.Errorf("InMemory database should not throw an error: %s", dbError.Error())
	}
}

func TestNewDatabaseReturnsMongoDbAccordingToConfig(t *testing.T) {
	config := utils.Config{
		Server: utils.ConfigServer{TypeDB: "mongo"},
	}
	db, _ := dataservice.NewDatabase(&config)
	dbType := fmt.Sprintf("%T", db)
	if dbType != "*dataservice.MongoDB" {
		t.Errorf("Bad DB concrete type created %s", dbType)
	}
}
