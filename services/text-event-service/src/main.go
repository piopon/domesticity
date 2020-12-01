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

func main() {
	config := utils.NewConfig("")
	logger := log.New(os.Stdout, config.Name+" > ", log.LstdFlags|log.Lmsgprefix)
	dataservice, dbError := dataservice.NewDatabase(config)
	if dbError != nil {
		logger.Println(dbError.Error())
		return
	}
	homeHandler := handlers.NewHome("resources/index.html", logger, config)
	docsHandler := handlers.NewDocs("resources/swagger.yaml")
	eventsHandler := handlers.NewEvents(logger, utils.NewValidator(), dataservice)

	routerMain := mux.NewRouter()

	routerGET := routerMain.Methods(http.MethodGet).Subrouter()
	routerGET.Path("/").HandlerFunc(homeHandler.GetIndex)
	routerGET.Path("/docs").HandlerFunc(docsHandler.GetDocumentation)
	routerGET.Path("/resources/swagger.yaml").HandlerFunc(docsHandler.GetSwagger)
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
		Addr:         config.Server.IP + ":" + config.Server.Port,
		Handler:      routerMain,
		ErrorLog:     logger,
		IdleTimeout:  time.Duration(config.Server.Timeout.Idle) * time.Second,
		ReadTimeout:  time.Duration(config.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(config.Server.Timeout.Write) * time.Second,
	}

	go func() {
		logger.Println("Starting server on port", config.Server.Port)
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
