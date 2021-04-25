package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/handlers"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
)

func main() {
	// create main service application objects
	app, error := initialize(utils.NewConfig(""))
	if error != nil {
		app.logger.Println("Cannot initialize application:")
		app.logger.Println(error.Error())
		os.Exit(1)
	}
	server := createServer(app)
	// run server in background and serve incoming requests
	go func() {
		app.logger.Println("Starting server on port", app.config.Server.Port)
		workError := server.ListenAndServe()
		if workError != http.ErrServerClosed {
			app.logger.Fatal("Error starting server:", workError)
			os.Exit(1)
		}
	}()
	// watch interrupt or kill signal (to stop program execution)
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	app.logger.Println("Shutting down by", <-quitChannel)
	// gracefully shutdown all connections
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
func initialize(config *utils.Config) (*Application, error) {
	logger := log.New(os.Stdout, config.Name+" > ", log.LstdFlags|log.Lmsgprefix)
	dataservice, dbError := dataservice.NewDatabase(config)
	return &Application{logger, config, dataservice}, dbError
}

// createServer is used to create a server object instance
func createServer(app *Application) *http.Server {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8100"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	return &http.Server{
		Addr:         app.config.Server.IP + ":" + app.config.Server.Port,
		Handler:      corsHandler.Handler(createRouter(createHandlers(app))),
		ErrorLog:     app.logger,
		IdleTimeout:  time.Duration(app.config.Server.Timeout.Idle) * time.Second,
		ReadTimeout:  time.Duration(app.config.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(app.config.Server.Timeout.Write) * time.Second,
	}
}

// createHandlers is used to create all neccessary handlers
func createHandlers(app *Application) (*handlers.Home, *handlers.Docs, *handlers.Events) {
	homeHandler := handlers.NewHome("resources/index.html", app.logger, app.config)
	docsHandler := handlers.NewDocs("resources/swagger.yaml")
	eventsHandler := handlers.NewEvents(app.logger, utils.NewValidator(), app.database)
	return homeHandler, docsHandler, eventsHandler
}

// createRouter is used to create a new endpoints routes and connect them with handlers
func createRouter(home *handlers.Home, docs *handlers.Docs, events *handlers.Events) *mux.Router {
	router := mux.NewRouter()
	// bind GET method paths to concrete methods handlers
	routerGET := router.Methods(http.MethodGet).Subrouter()
	routerGET.Path("/").HandlerFunc(home.GetIndex)
	routerGET.Path("/docs").HandlerFunc(docs.GetDocumentation)
	routerGET.Path("/resources/swagger.yaml").HandlerFunc(docs.GetSwagger)
	routerGET.Path("/events").HandlerFunc(events.GetEvents)
	routerGET.Path("/events/{id}").HandlerFunc(events.GetEvent)
	// bind POST method paths to concrete methods handlers
	routerPOST := router.Methods(http.MethodPost).Subrouter()
	routerPOST.Use(events.ValidationMiddleware)
	routerPOST.Path("/events").HandlerFunc(events.AddEvent)
	// bind PUT method paths to concrete methods handlers
	routerPUT := router.Methods(http.MethodPut).Subrouter()
	routerPUT.Use(events.ValidationMiddleware)
	routerPUT.Path("/events/{id}").HandlerFunc(events.UpdateEvent)
	// bind DELETE method paths to concrete methods handlers
	routerDELETE := router.Methods(http.MethodDelete).Subrouter()
	routerDELETE.Path("/events/{id}").HandlerFunc(events.DeleteEvent)
	return router
}
