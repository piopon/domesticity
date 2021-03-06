package model

import (
	"math"
	"time"
)

// TimeSpan is an struct representing event start and end time
// swagger:model
type TimeSpan struct {
	// The event start time in RFC3339 standard ("2006-01-02T15:04:05Z07:00")
	//
	// required: true
	Start time.Time `json:"start" bson:"start" validate:"required,date-time"`
	// The event stop time in RFC3339 standard ("2006-01-02T15:04:05Z07:00")
	//
	// required: true
	Stop time.Time `json:"stop" bson:"stop" validate:"required,date-time,gtfield=Start"`
}

// Duration calulates event time duration
func (span *TimeSpan) Duration() time.Duration {
	return span.Stop.Sub(span.Start)
}

// DaysAffected checks if Date lasts over one day
func (span *TimeSpan) DaysAffected() int {
	return int(math.Round(span.Duration().Hours() / float64(24)))
}
