package data

import (
	"encoding/json"
	"io"
)

// FromJSON is a method called on Event struct with specified IO Reader
func (event *Event) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(event)
}

// ToJSON is a method called on Event slice with specified IO Writer
func (events *Events) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(events)
}
