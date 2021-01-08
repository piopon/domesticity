package dataservice_test

import (
	"net/url"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewInMemoryReturnsObjectAndNoError(t *testing.T) {
	memory, error := dataservice.NewInMemory(nil)
	if error != nil {
		t.Errorf("InMemory factory method returned with error: %s", error.Error())
	}
	if memory == nil {
		t.Error("InMemory factory method returned nil object")
	}
}

func TestShutdownDoesNotInvokePanic(t *testing.T) {
	memory, _ := dataservice.NewInMemory(nil)
	memory.Shutdown(nil)
}

func TestGetEventsReturnsItemsWithEmptyQuery(t *testing.T) {
	memory, _ := initializeTest()
	queryParams := url.Values{}
	events, error := memory.GetEvents(queryParams)
	if error != nil {
		t.Error("Get events with empty query should not return an error")
	}
	if len(*events) != 1 {
		t.Error("Returned incorrect number of events")
	}
	closeTest(memory)
}

func TestGetEventsReturnsItemsWithCorrectQuery(t *testing.T) {
	memory, _ := initializeTest()
	queryParams := url.Values{"owner": []string{"PNK"}}
	events, error := memory.GetEvents(queryParams)
	if error != nil {
		t.Error("Get events with correct query should not return an error")
	}
	if len(*events) != 0 {
		t.Error("Returned incorrect number of events")
	}
	closeTest(memory)
}

func TestGetEventsReturnsErrorWhenIncorrectQuery(t *testing.T) {
	customFilters := dataservice.Filters{}
	memory, _ := dataservice.NewInMemory(&customFilters)
	queryParams := url.Values{"owner": []string{"PNK"}}
	_, error := memory.GetEvents(queryParams)
	if error == nil {
		t.Error("Get events with invalid query should return an error")
	}
	closeTest(memory)
}

func TestGetEventShouldReturnItemWithSpecifiedID(t *testing.T) {
	memory, id := initializeTest()
	event, error := memory.GetEvent(id)
	if error != nil {
		t.Error("Get event should return item and not return an error")
	}
	if event.Owner != "Admin" || event.Category != "Notes" || event.Title != "This is my first event" {
		t.Errorf("Returned item is incorrect: %v", event)
	}
	closeTest(memory)
}

func TestGetEventReturnsErrorWhenItemNotExists(t *testing.T) {
	memory, _ := initializeTest()
	event, error := memory.GetEvent(primitive.NilObjectID)
	if error == nil {
		t.Error("Get event should throw an error when ID does not exists")
	}
	if event != nil {
		t.Error("Returned event should be nil if ID does not exists")
	}
}

func TestAddEventShouldInsertNewItem(t *testing.T) {
	memory, _ := initializeTest()
	error := memory.AddEvent(&model.Event{})
	if error != nil {
		t.Error("Add event should insert a new item and not return an error")
	}
	closeTest(memory)
}

func TestAddEventShouldFailIfAddingEventWithNilID(t *testing.T) {
	memory, _ := initializeTest()
	existingID := getInitialObjectId(memory)
	error := memory.AddEvent(&model.Event{ID: existingID})
	if error == nil {
		t.Error("Add event should return an error if ID = nil")
	}
	closeTest(memory)
}

func TestUpdateExistingEventShouldUpdateItem(t *testing.T) {
	memory, id := initializeTest()
	error := memory.UpdateEvent(id, &model.Event{})
	if error != nil {
		t.Error("Update event should not return an error when ID exists")
	}
	closeTest(memory)
}

func TestUpdateNotExistingEventShouldReturnError(t *testing.T) {
	memory, _ := initializeTest()
	error := memory.UpdateEvent(primitive.NilObjectID, &model.Event{})
	if error == nil {
		t.Error("Update event should throw an error when ID does not exists")
	}
	closeTest(memory)
}

func TestDeleteExistingEventShouldRemoveItem(t *testing.T) {
	memory, id := initializeTest()
	error := memory.DeleteEvent(id)
	if error != nil {
		t.Error("Delete event should not return an error when ID exists")
	}
	closeTest(memory)
}

func TestDeleteNotExistingEventShouldReturnError(t *testing.T) {
	memory, _ := initializeTest()
	error := memory.DeleteEvent(primitive.NilObjectID)
	if error == nil {
		t.Error("Delete event should theow an error when ID does not exists")
	}
	closeTest(memory)
}

func initializeTest() (*dataservice.InMemory, primitive.ObjectID) {
	config := utils.ConfigServer{TypeDB: "memory"}
	filters := dataservice.NewFilters(&config)
	memory, _ := dataservice.NewInMemory(filters)
	id := getInitialObjectId(memory)
	return memory, id
}

func closeTest(memory *dataservice.InMemory) {
	memory.Shutdown(nil)
}

func getInitialObjectId(db *dataservice.InMemory) primitive.ObjectID {
	events, errors := db.GetEvents(nil)
	if errors != nil {
		return primitive.NilObjectID
	}
	return (*events)[0].ID
}
