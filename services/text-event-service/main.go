package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		fmt.Println("got request!")
	})
	http.ListenAndServe(":10000", nil)
}
