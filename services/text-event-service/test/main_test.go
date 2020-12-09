package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/handlers"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
)

func TestCreateRouterCreatesCorrectPathRouter(t *testing.T) {
	helper := newHelper()
	router := createRouter(helper.createTestHandlers())
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	testTable := []struct {
		name     string
		server   *httptest.Server
		method   string
		url      string
		exitCode int
	}{
		{"home index path", testServer, "GET", "/", 200},
		{"documentation handler", testServer, "GET", "/docs", 200},
		{"get all events handler", testServer, "GET", "/events", 200},
		{"not existing path", testServer, "GET", "/not-existing", 404},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			request, err := http.NewRequest(testCase.method, testCase.server.URL+testCase.url, nil)
			if err != nil {
				t.Fatalf("Cannot create request for client: %v", err)
			}
			response, err := testCase.server.Client().Do(request)
			if err != nil {
				t.Fatalf("Cannot process created request: %v", err)
			}
			if response.StatusCode != testCase.exitCode {
				t.Fatalf("Expected a status code of %v, got %v", testCase.exitCode, response.StatusCode)
			}
		})
	}
}

type helper struct {
	db     dataservice.Database
	config *utils.Config
	logger *log.Logger
}

func newHelper() *helper {
	config := createTestConfig()
	dataservice, _ := dataservice.NewDatabase(config)
	return &helper{
		db:     dataservice,
		config: config,
		logger: log.New(os.Stdout, config.Name+" > ", log.LstdFlags|log.Lmsgprefix),
	}
}

func createTestConfig() *utils.Config {
	return &utils.Config{
		Name:    "test-name",
		Verbose: false,
		Server: utils.ConfigServer{
			TypeDB: "memory",
		},
		MongoDB: utils.ConfigMongo{},
	}
}

func (h *helper) createTestHandlers() (*handlers.Home, *handlers.Docs, *handlers.Events) {
	home := handlers.NewHome("../resources/index.html", h.logger, h.config)
	docs := handlers.NewDocs("../resources/swagger.yaml")
	events := handlers.NewEvents(h.logger, utils.NewValidator(), h.db)
	return home, docs, events
}
