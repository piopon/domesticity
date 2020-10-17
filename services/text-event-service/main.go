package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/handlers"
)

func main() {
	logger := log.New(os.Stdout, "text-event-service > ", log.LstdFlags|log.Lmsgprefix)

	homeHandler := handlers.NewHome(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", homeHandler)

	server := &http.Server{
		Addr:         ":10000",
		Handler:      serveMux,
		IdleTimeout:  300 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		workError := server.ListenAndServe()
		if workError != nil {
			logger.Fatal(workError)
		}
	}()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if cancel != nil {
		logger.Fatal(cancel)
	}
	server.Shutdown(shutdownCtx)
}
