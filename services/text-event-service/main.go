package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("format")
	http.ListenAndServe(":10000", nil)
}
