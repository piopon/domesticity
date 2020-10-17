package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/handlers"
)

func main() {
	logger := log.New(os.Stdout, "text-event-service > ", log.LstdFlags|log.Lmsgprefix)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", handlers.NewHome(logger))
	serveMux.Handle("/events", handlers.NewEvents(logger))

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

	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, os.Interrupt)
	signal.Notify(quitChannel, os.Kill)
	logger.Println("Shutting down by", <-quitChannel)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(shutdownCtx)
}
