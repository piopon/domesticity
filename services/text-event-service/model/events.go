package model

import (
	"fmt"
	"net/url"
	"strconv"
)

// Events is a type definition for slice of Event pointers
type Events []*Event

var availableFilters = map[string]interface{}{
	"limit": filterLimit,
}

// Filter filters current event list according to provided params
func (events *Events) Filter(params url.Values) (*Events, error) {
	filteredEvents := *events
	for key, value := range params {
		if filter, ok := availableFilters[key]; ok {
			error := filter.(func(*Events, string) error)(&filteredEvents, value[0])
			if error != nil {
				return events, error
			}
		}
	}
	return &filteredEvents, nil
}

func filterLimit(input *Events, limit string) error {
	limitParsed, error := strconv.Atoi(limit)
	if error != nil {
		fmt.Println("Filter limit: cannot parse limit value", limit)
		return error
	}
	if limitParsed > len(*input) {
		limitParsed = len(*input)
	}
	*input = (*input)[:limitParsed]
	return nil
}
