package model_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFilterCombinedReturnsCorrectReturn(t *testing.T) {
	testFilters := url.Values{
		"owner":   []string{"Admin"},
		"dayStop": []string{"2020-03-03"},
		"limit":   []string{"1"},
	}
	expectedResults := model.Events{getTestEvent3()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterLimitReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"limit": []string{"2"}}
	expectedResults := model.Events{getTestEvent1(), getTestEvent2()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterLimitReturnsInputIfLimitExceeds(t *testing.T) {
	testFilters := url.Values{"limit": []string{"100"}}
	expectedResults := model.Events{getTestEvent1(), getTestEvent2(), getTestEvent3()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterLimitFailsIfValueNotAnInt(t *testing.T) {
	testFilters := url.Values{"limit": []string{"1.1"}}
	resultEvents, resultError := getTestEvents().Filter(testFilters)
	if resultError == nil {
		t.Error("Limit value should fail when set to not an integer value")
	}
	expectedResults := model.Events{getTestEvent1(), getTestEvent2(), getTestEvent3()}
	checkError := checkResultEvents(expectedResults, *resultEvents)
	if checkError != nil {
		t.Error(checkError.Error())
	}
}

func TestFilterOffsetReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"offset": []string{"1"}}
	expectedResults := model.Events{getTestEvent2(), getTestEvent3()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterOffsetReturnsInputIfLimitExceedsResult(t *testing.T) {
	testFilters := url.Values{"offset": []string{"100"}}
	expectedResults := model.Events{getTestEvent1(), getTestEvent2(), getTestEvent3()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterOffsetFailsIfValueNotAnInt(t *testing.T) {
	testFilters := url.Values{"offset": []string{"1.1"}}
	resultEvents, resultError := getTestEvents().Filter(testFilters)
	if resultError == nil {
		t.Error("Offset value should fail when set to not an integer value")
	}
	expectedResults := model.Events{getTestEvent1(), getTestEvent2(), getTestEvent3()}
	checkError := checkResultEvents(expectedResults, *resultEvents)
	if checkError != nil {
		t.Error(checkError.Error())
	}
}

func TestFilterOwnerReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"owner": []string{"User"}}
	expectedResults := model.Events{getTestEvent2()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterCategoryReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"category": []string{"Blue"}}
	expectedResults := model.Events{getTestEvent3()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterTitleReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"title": []string{"first"}}
	expectedResults := model.Events{getTestEvent1()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterContentReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"content": []string{"number"}}
	expectedResults := model.Events{getTestEvent1(), getTestEvent2(), getTestEvent3()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterDayStartReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"dayStart": []string{"2020-02-02"}}
	expectedResults := model.Events{getTestEvent2()}
	executeEventsTest(t, testFilters, expectedResults)
}

func TestFilterDayStopReturnsCorrectResult(t *testing.T) {
	testFilters := url.Values{"dayStop": []string{"2020-03-03"}}
	expectedResults := model.Events{getTestEvent2(), getTestEvent3()}
	executeEventsTest(t, testFilters, expectedResults)
}

func executeEventsTest(t *testing.T, testFilters url.Values, expectedResults model.Events) {
	resultEvents, resultError := getTestEvents().Filter(testFilters)
	if resultError != nil {
		t.Error("Error while filtering test data: " + resultError.Error())
	}
	checkError := checkResultEvents(expectedResults, *resultEvents)
	if checkError != nil {
		t.Error(checkError.Error())
	}
}

func checkResultEvents(expectEvents model.Events, resultEvents model.Events) error {
	resultNo := len(resultEvents)
	expectNo := len(expectEvents)
	if resultNo != expectNo {
		return fmt.Errorf("Incorrect result data size: %d", resultNo)
	}
	for i, resultEvent := range resultEvents {
		expectedEvent := expectEvents[i]
		if resultEvent.ToString() != expectedEvent.ToString() {
			return fmt.Errorf("Incorrect result: %v", resultEvent)
		}
	}
	return nil
}

func getTestEvents() *model.Events {
	return &model.Events{getTestEvent1(), getTestEvent2(), getTestEvent3()}
}

func getTestEvent1() *model.Event {
	return &model.Event{
		ID:    primitive.NewObjectID(),
		Title: "This is my first event",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(2020, 01, 01, 00, 00, 00, 00, time.Local),
			Stop:  time.Date(2020, 02, 02, 00, 00, 01, 00, time.Local)},
		Category: "Red",
		Content:  "Test event number 1",
	}
}

func getTestEvent2() *model.Event {
	return &model.Event{
		ID:    primitive.NewObjectID(),
		Title: "Event no. 2",
		Owner: "User",
		Occurence: model.TimeSpan{
			Start: time.Date(2020, 02, 02, 00, 00, 00, 00, time.Local),
			Stop:  time.Date(2020, 03, 03, 00, 00, 01, 00, time.Local)},
		Category: "Red",
		Content:  "Test event number 2",
	}
}

func getTestEvent3() *model.Event {
	return &model.Event{
		ID:    primitive.NewObjectID(),
		Title: "3 EVENT",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(2020, 03, 03, 00, 00, 00, 00, time.Local),
			Stop:  time.Date(2020, 03, 03, 00, 00, 01, 00, time.Local)},
		Category: "Blue",
		Content:  "Test event number 3",
	}
}
