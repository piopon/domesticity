package model

import (
	"time"
)

// Date is an struct representing event time
type Date struct {
	Start time.Time `json:"start" validate:"required,datetime"`
	Stop  time.Time `json:"stop" validate:"required,datetime"`
}
