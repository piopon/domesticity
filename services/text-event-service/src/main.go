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
	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/handlers"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
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

	server := &http.Server{
		Addr:         config.Server.IP + ":" + config.Server.Port,
		Handler:      createRouter(homeHandler, docsHandler, eventsHandler),
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

// Application is a struct containing all top-level settings
type Application struct {
	logger   *log.Logger
	config   *utils.Config
	database dataservice.Database
}

// createRouter is used to create a new endpoints routes and connect them with handlers
func createRouter(home *handlers.Home, docs *handlers.Docs, events *handlers.Events) *mux.Router {
	router := mux.NewRouter()

	routerGET := router.Methods(http.MethodGet).Subrouter()
	routerGET.Path("/").HandlerFunc(home.GetIndex)
	routerGET.Path("/docs").HandlerFunc(docs.GetDocumentation)
	routerGET.Path("/resources/swagger.yaml").HandlerFunc(docs.GetSwagger)
	routerGET.Path("/events").HandlerFunc(events.GetEvents)
	routerGET.Path("/events/{id}").HandlerFunc(events.GetEvent)

	routerPOST := router.Methods(http.MethodPost).Subrouter()
	routerPOST.Use(events.ValidationMiddleware)
	routerPOST.Path("/events").HandlerFunc(events.AddEvent)

	routerPUT := router.Methods(http.MethodPut).Subrouter()
	routerPUT.Use(events.ValidationMiddleware)
	routerPUT.Path("/events/{id}").HandlerFunc(events.UpdateEvent)

	routerDELETE := router.Methods(http.MethodDelete).Subrouter()
	routerDELETE.Path("/events/{id}").HandlerFunc(events.DeleteEvent)

	return router
}
