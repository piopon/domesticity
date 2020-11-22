package model_test

import (
	"testing"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/model"
)

func TestDurationReturnsCorrectTimesDifference(t *testing.T) {
	baseTime := time.Now()
	expectedDuration := time.Duration(90)
	span := model.TimeSpan{
		Start: baseTime,
		Stop:  baseTime.Add(expectedDuration),
	}
	diff := span.Duration()
	if diff != expectedDuration {
		t.Errorf("TimeSpan duration returned incorrect result: %v", diff)
	}
}

func TestDaysAffectedReturnsCorrectDaysDifference(t *testing.T) {
	startTime, startError := time.Parse("2006-01-02", "2020-12-01")
	if startError != nil {
		t.Error("Cannot parse start time" + startError.Error())
	}
	stopTime, stopError := time.Parse("2006-01-02", "2020-12-08")
	if stopError != nil {
		t.Error("Cannot parse stop time" + stopError.Error())
	}
	span := model.TimeSpan{Start: startTime, Stop: stopTime}
	days := span.DaysAffected()
	if days != 7 {
		t.Errorf("TimeSpan days affected returned incorrect result: %v", days)
	}
}
