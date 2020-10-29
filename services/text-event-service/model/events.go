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
	"offset":   filterOffset,
	"title":    filterTitle,
	"owner":    filterOwner,
	"dayStart": filterDayStart,
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
		return fmt.Errorf("Filter limit: cannot parse limit value %s", limit)
	}
	if limitParsed > len(*input) {
		limitParsed = len(*input)
	}
	*input = (*input)[:limitParsed]
	return nil
}

func filterOffset(input *Events, offset string) error {
	offsetParsed, error := strconv.Atoi(offset)
	if error != nil {
		return fmt.Errorf("Filter limit: cannot parse offset value %s", offset)
	}
	if offsetParsed > len(*input) {
		offsetParsed = 0
	}
	*input = (*input)[offsetParsed:]
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

func filterDayStart(input *Events, dateStart string) error {
	return filterEventField(input, func(event *Event) bool {
		return event.Occurence.Start.Format("2006-01-02") == dateStart
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
