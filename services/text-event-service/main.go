package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("got request!")
		ioutil.ReadAll(request.Body)
	})
	http.ListenAndServe(":10000", nil)
}
