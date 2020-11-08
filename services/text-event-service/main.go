package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/piopon/domesticity/services/text-event-service/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/handlers"
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

var addressIP = ""
var addressPort = "9999"

func main() {
	logger := log.New(os.Stdout, "text-event-service > ", log.LstdFlags|log.Lmsgprefix)
	dataservice := dataservice.NewInMemory()

	homeHandler := handlers.NewHome(logger)
	docsHandler := handlers.NewDocs("scripts/swagger.yaml")
	eventsHandler := handlers.NewEvents(logger, utils.NewValidator(), dataservice)

	routerMain := mux.NewRouter()

	routerGET := routerMain.Methods(http.MethodGet).Subrouter()
	routerGET.Path("/").HandlerFunc(homeHandler.ServeHTTP)
	routerGET.Path("/docs").HandlerFunc(docsHandler.GetDocumentation)
	routerGET.Path("/scripts/swagger.yaml").HandlerFunc(docsHandler.GetSwagger)
	routerGET.Path("/events").HandlerFunc(eventsHandler.GetEvents)
	routerGET.Path("/events/{id}").HandlerFunc(eventsHandler.GetEvent)

	routerPOST := routerMain.Methods(http.MethodPost).Subrouter()
	routerPOST.Use(eventsHandler.ValidationMiddleware)
	routerPOST.Path("/events").HandlerFunc(eventsHandler.AddEvent)

	routerPUT := routerMain.Methods(http.MethodPut).Subrouter()
	routerPUT.Use(eventsHandler.ValidationMiddleware)
	routerPUT.Path("/events/{id}").HandlerFunc(eventsHandler.UpdateEvent)

	routerDELETE := routerMain.Methods(http.MethodDelete).Subrouter()
	routerDELETE.Path("/events/{id}").HandlerFunc(eventsHandler.DeleteEvent)

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
	signal.Notify(quitChannel, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	logger.Println("Shutting down by", <-quitChannel)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	dataservice.Shutdown(shutdownCtx)
	server.Shutdown(shutdownCtx)
}
