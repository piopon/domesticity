package model

import "strings"

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// ToString is a metod to return all errors as a single string
func (errors ValidationError) ToString(bullet string) string {
	return bullet + strings.Join(errors.Messages, "\n"+bullet)
}
