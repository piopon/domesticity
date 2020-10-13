package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		body, error := ioutil.ReadAll(request.Body)
		if error != nil {
			http.Error(response, "Bad request...", http.StatusBadRequest)
			return
		}
		fmt.Printf("Request: %s\n", body)
		fmt.Fprintf(response, "Response: %s", body)
	})
	http.ListenAndServe(":10000", nil)
}
