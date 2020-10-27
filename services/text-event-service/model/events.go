package model

import (
	"fmt"
	"net/url"
)

// Events is a type definition for slice of Event pointers
type Events []*Event

// Filter filters current event list according to provided params
func (events *Events) Filter(params url.Values) (*Events, error) {
	filteredEvents := *events
	for key, value := range params {
		fmt.Println(key, " => ", value)
	}
	return &filteredEvents, nil
}

