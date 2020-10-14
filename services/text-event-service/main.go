package main

import (
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/handlers"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", handlers.NewHome())
	http.ListenAndServe(":10000", serveMux)
}
