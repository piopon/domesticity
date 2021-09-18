package model

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Events is a type definition for slice of Event pointers
type Events []*Event

// availableFilters is a registered list of filters (used only in memory dataservice)
var availableFilters = map[string]interface{}{
	"limit":    filterLimit,
	"offset":   filterOffset,
	"title":    filterTitle,
	"owner":    filterOwner,
	"dayStart": filterDayStart,
	"dayStop":  filterDayStop,
	"category": filterCategory,
	"content":  filterContent,
}

// Filter filters current event list according to provided params (used only in memory dataservice)
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

// filterLimit limits the result list to specified value (used only in memory dataservice)
func filterLimit(input *Events, limit string) error {
	limitParsed, error := strconv.Atoi(limit)
	if error != nil {
		return fmt.Errorf("filter limit - cannot parse limit value %s", limit)
	}
	if limitParsed > len(*input) {
		limitParsed = len(*input)
	}
	*input = (*input)[:limitParsed]
	return nil
}

// filterOffset shifts the start of result list by specified value (used only in memory dataservice)
func filterOffset(input *Events, offset string) error {
	offsetParsed, error := strconv.Atoi(offset)
	if error != nil {
		return fmt.Errorf("filter limit - cannot parse offset value %s", offset)
	}
	if offsetParsed > len(*input) {
		offsetParsed = 0
	}
	*input = (*input)[offsetParsed:]
	return nil
}

// filterOwner filters results by owner (used only in memory dataservice)
func filterOwner(input *Events, owner string) error {
	return filterEventField(input, func(event *Event) bool {
		return event.Owner == owner
	})
}

// filterCategory filters results by category (used only in memory dataservice)
func filterCategory(input *Events, category string) error {
	return filterEventField(input, func(event *Event) bool {
		return event.Category == category
	})
}

// filterTitle filters results by title (used only in memory dataservice)
func filterTitle(input *Events, title string) error {
	return filterEventField(input, func(event *Event) bool {
		return strings.Contains(event.Title, title)
	})
}

// filterContent filters results by content (used only in memory dataservice)
func filterContent(input *Events, content string) error {
	return filterEventField(input, func(event *Event) bool {
		return strings.Contains(event.Content, content)
	})
}

// filterDayStart filters results by event start date (used only in memory dataservice)
func filterDayStart(input *Events, dateStart string) error {
	return filterEventField(input, func(event *Event) bool {
		return event.Occurence.Start.Format("2006-01-02") == dateStart
	})
}

// filterDayStop filters results by event stop date (used only in memory dataservice)
func filterDayStop(input *Events, dateStop string) error {
	return filterEventField(input, func(event *Event) bool {
		return event.Occurence.Stop.Format("2006-01-02") == dateStop
	})
}

// filterEventField s a generic method with filter predicate (used only in memory dataservice)
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
