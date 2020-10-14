package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Home is a service handler used after visiting main service URL
type Home struct{}

// NewHome is a factory method to create Home service handler
func NewHome() *Home {
	return &Home{}
}

func (home *Home) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	body, error := ioutil.ReadAll(request.Body)
	if error != nil {
		http.Error(response, "Bad request...", http.StatusBadRequest)
		return
	}
	fmt.Printf("Request: %s\n", body)
	fmt.Fprintf(response, "Response: %s", body)
}
