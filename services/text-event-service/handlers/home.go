package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Home is a service handler used after visiting main service URL
type Home struct {
	logger *log.Logger
}

// NewHome is a factory method to create Home service handler
func NewHome(logger *log.Logger) *Home {
	return &Home{logger}
}

func (home *Home) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	body, error := ioutil.ReadAll(request.Body)
	if error != nil {
		http.Error(response, "Bad request...", http.StatusBadRequest)
		return
	}
	home.logger.Printf("Request: %s\n", body)
	fmt.Fprintf(response, "Response: %s", body)
}
