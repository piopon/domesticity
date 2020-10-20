package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/piopon/domesticity/services/text-event-service/handlers"
)

var addressIP = ""
var addressPort = "9999"

func main() {
	logger := log.New(os.Stdout, "text-event-service > ", log.LstdFlags|log.Lmsgprefix)

	homeHandler := handlers.NewHome(logger)
	eventsHandler := handlers.NewEvents(logger)

	routerMain := mux.NewRouter()

	routerGET := routerMain.Methods(http.MethodGet).Subrouter()

	routerPOST := routerMain.Methods(http.MethodPost).Subrouter()

	routerPUT := routerMain.Methods(http.MethodPut).Subrouter()

	server := &http.Server{
		Addr:         addressIP + ":" + addressPort,
		Handler:      routerMain,
		ErrorLog:     logger,
		IdleTimeout:  300 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port", addressPort)
		workError := server.ListenAndServe()
		if workError != nil {
			logger.Fatal("Error starting server:", workError)
			os.Exit(1)
		}
	}()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt)
	signal.Notify(quitChannel, os.Kill)
	logger.Println("Shutting down by", <-quitChannel)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(shutdownCtx)
}
