package model

import (
	"time"
)

// TimeSpan is an struct representing event start and end time
type TimeSpan struct {
	Start time.Time `json:"start" validate:"required,datetime"`
	Stop  time.Time `json:"stop" validate:"required,datetime"`
}
