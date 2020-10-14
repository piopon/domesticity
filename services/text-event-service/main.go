package main

import (
	"log"
	"net/http"
	"os"

	"github.com/piopon/domesticity/services/text-event-service/handlers"
)

func main() {
	handlersLogger := log.New(os.Stdout, "text-event-service > ", log.LstdFlags|log.Lmsgprefix)

	homeHandler := handlers.NewHome(handlersLogger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", homeHandler)

	server := http.Server{
		Addr:    ":10000",
		Handler: serveMux,
	}
	server.ListenAndServe()
}
