package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Home struct {
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
