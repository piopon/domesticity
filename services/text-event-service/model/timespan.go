package model

import (
	"math"
	"time"
)

// TimeSpan is an struct representing event start and end time
// swagger:model
type TimeSpan struct {
	// The event start time
	//
	// required: true
	Start time.Time `json:"start" validate:"required,date-time"`
	// The event stop time
	//
	// required: true
	Stop time.Time `json:"stop" validate:"required,date-time,gtfield=Start"`
}

// Duration calulates event time duration
func (span *TimeSpan) Duration() time.Duration {
	return span.Stop.Sub(span.Start)
}

// DaysAffected checks if Date lasts over one day
func (span *TimeSpan) DaysAffected() int {
	return int(math.Round(span.Duration().Hours() / float64(24)))
}
