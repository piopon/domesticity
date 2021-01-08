package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/handlers"
	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInitializeShouldCorrectlyCreateTopLevelStruct(t *testing.T) {
	app, err := initialize(createTestConfig())
	if app == nil {
		t.Fatalf("Initialize should create an top-level struct but it failed: %v", err)
	}
}

func TestCreateHandlersShouldCorrectlyCreateHandlers(t *testing.T) {
	app, _ := initialize(createTestConfig())
	home, docs, events := createHandlers(app)
	if home == nil {
		t.Fatalf("Create handler should create home handler but it failed")
	}
	if docs == nil {
		t.Fatalf("Create handler should create docs handler but it failed")
	}
	if events == nil {
		t.Fatalf("Create handler should create events handler but it failed")
	}
}

func TestCreateServerShouldCorrectlyCreateServerObject(t *testing.T) {
	app, _ := initialize(createTestConfig())
	server := createServer(app)
	if server == nil {
		t.Fatal("Create server should create server object but it failed")
	}
}

func TestCreateRouterCreatesCorrectPathRouter(t *testing.T) {
	helper := newHelper()
	initID := helper.getDatabaseIds()[0]
	router := createRouter(helper.createTestHandlers())
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	testTable := []struct {
		name     string
		server   *httptest.Server
		method   string
		url      string
		event    *model.Event
		exitCode int
	}{
		{"path: get home", testServer, "GET", "/", nil, 200},
		{"path: get docs", testServer, "GET", "/docs", nil, 200},
		{"path: get all events ", testServer, "GET", "/events", nil, 200},
		{"path: get single event", testServer, "GET", "/events/" + initID.Hex(), nil, 200},
		{"path: post new event", testServer, "POST", "/events", helper.createEvent(), 200},
		{"path: update event", testServer, "PUT", "/events/" + initID.Hex(), helper.createEvent(), 200},
		{"path: delete event", testServer, "DELETE", "/events/" + initID.Hex(), nil, 204},
		{"path: not existing", testServer, "GET", "/not-existing", nil, 404},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if testCase.event != nil {
				error := utils.ToJSON(testCase.event, &buffer)
				if error != nil {
					t.Errorf("Could not convert event to JSON")
				}
			}
			requestBody := strings.NewReader(buffer.String())
			request, err := http.NewRequest(testCase.method, testCase.server.URL+testCase.url, requestBody)
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

func (h *helper) getDatabaseIds() []primitive.ObjectID {
	events, errors := h.db.GetEvents(nil)
	if errors != nil {
		return nil
	}
	result := []primitive.ObjectID{}
	for _, event := range *events {
		result = append(result, event.ID)
	}
	return result
}

func (h *helper) createEvent() *model.Event {
	return &model.Event{
		Title: "This is my first event",
		Owner: "Admin",
		Occurence: model.TimeSpan{
			Start: time.Date(2020, 05, 26, 14, 15, 00, 00, time.Local),
			Stop:  time.Date(2020, 05, 27, 10, 30, 00, 00, time.Local)},
		Category: "Notes",
		Content:  "Test event number 1",
	}
}
