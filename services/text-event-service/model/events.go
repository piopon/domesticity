package model

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Events is a type definition for slice of Event pointers
type Events []*Event

var availableFilters = map[string]interface{}{
	"limit":    filterLimit,
	"title":    filterTitle,
	"owner":    filterOwner,
	"category": filterCategory,
	"content":  filterContent,
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

func filterOwner(input *Events, owner string) error {
	return filterEventField(input, func(event *Event) bool {
		return event.Owner == owner
	})
}

func filterCategory(input *Events, category string) error {
	return filterEventField(input, func(event *Event) bool {
		return event.Category == category
	})
}

func filterTitle(input *Events, title string) error {
	return filterEventField(input, func(event *Event) bool {
		return strings.Contains(event.Title, title)
	})
}

func filterContent(input *Events, content string) error {
	return filterEventField(input, func(event *Event) bool {
		return strings.Contains(event.Content, content)
	})
}

func filterEventField(input *Events, predicate func(*Event) bool) error {
	filteredEvents := Events{}
	for i := range *input {
		if predicate((*input)[i]) {
			filteredEvents = append(filteredEvents, (*input)[i])
		}
	}
	*input = filteredEvents
	return nil
}
