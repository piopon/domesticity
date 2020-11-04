package model

import (
	"math"
	"time"
)

// TimeSpan is an struct representing event start and end time
type TimeSpan struct {
	Start time.Time `json:"start" validate:"required,date-time"`
	Stop  time.Time `json:"stop" validate:"required,date-time,gtfield=Start"`
}

// Duration calulates event time duration
func (span *TimeSpan) Duration() time.Duration {
	return span.Stop.Sub(span.Start)
}

// DaysAffected checks if Date lasts over one day
func (span *TimeSpan) DaysAffected() int {
	return int(math.Round(span.Duration().Hours() / float64(24)))
}
