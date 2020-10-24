package data

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the object from JSON format string
func FromJSON(i interface{}, reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(i)
}

// ToJSON serializes the interface into a JSON format string
func ToJSON(i interface{}, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(i)
}
