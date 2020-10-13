package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("got request!")
		body, error := ioutil.ReadAll(request.Body)
		if error == nil {
			fmt.Printf("Data: %s\n", body)
		}
	})
	http.ListenAndServe(":10000", nil)
}
