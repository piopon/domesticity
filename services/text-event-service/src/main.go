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
	app, error := initialize()
	if error != nil {
		app.logger.Println("Cannot initialize application:")
		app.logger.Println(error.Error())
		os.Exit(1)
	}
	homeHandler := handlers.NewHome("resources/index.html", app.logger, app.config)
	docsHandler := handlers.NewDocs("resources/swagger.yaml")
	eventsHandler := handlers.NewEvents(app.logger, utils.NewValidator(), app.database)

	server := &http.Server{
		Addr:         app.config.Server.IP + ":" + app.config.Server.Port,
		Handler:      createRouter(homeHandler, docsHandler, eventsHandler),
		ErrorLog:     app.logger,
		IdleTimeout:  time.Duration(app.config.Server.Timeout.Idle) * time.Second,
		ReadTimeout:  time.Duration(app.config.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(app.config.Server.Timeout.Write) * time.Second,
	}

	go func() {
		app.logger.Println("Starting server on port", app.config.Server.Port)
		workError := server.ListenAndServe()
		if workError != nil {
			app.logger.Fatal("Error starting server:", workError)
			os.Exit(1)
		}
	}()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	app.logger.Println("Shutting down by", <-quitChannel)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	app.database.Shutdown(shutdownCtx)
	server.Shutdown(shutdownCtx)
	app.logger.Println("Correctly shutdown database and server.")
}

// Application is a struct containing all top-level settings
type Application struct {
	logger   *log.Logger
	config   *utils.Config
	database dataservice.Database
}

// initialize is used create top-level Application struct
func initialize() (*Application, error) {
	config := utils.NewConfig("")
	logger := log.New(os.Stdout, config.Name+" > ", log.LstdFlags|log.Lmsgprefix)
	dataservice, dbError := dataservice.NewDatabase(config)
	return &Application{logger, config, dataservice}, dbError
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
