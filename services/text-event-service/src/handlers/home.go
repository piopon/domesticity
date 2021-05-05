package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
)

var (
	//BuildTime holds binary build time information
	BuildTime string
	//VersionAPI holds binary version information
	VersionAPI string
	//CommitSHA holds last commit hash information
	CommitSHA string
	//GoVersion holds Golang version information
	GoVersion string
)

// Home is a service handler used after visiting main service URL
type Home struct {
	html   string
	logger *log.Logger
	data   *HomeData
}

// HomeData is a struct for storing data displayed in index.html
type HomeData struct {
	Build  *BuildData
	Config *utils.Config
}

// BuildData is a struct holding all build information
type BuildData struct {
	API    string
	Time   string
	SHA    string
	Golang string
}

// NewHome is a factory method to create Home service handler
func NewHome(html string, logger *log.Logger, config *utils.Config) *Home {
	buildData := &BuildData{
		API:    VersionAPI,
		Time:   BuildTime,
		SHA:    CommitSHA,
		Golang: GoVersion,
	}
	return &Home{html, logger, &HomeData{buildData, config}}
}

// GetIndex is used to serve main page of service
func (home *Home) GetIndex(response http.ResponseWriter, request *http.Request) {
	home.logger.Println("Handling GET INDEX")
	response.Header().Add("Content-Type", "text/html; charset=utf-8")
	template, parseError := template.ParseFiles(home.html)
	if parseError != nil {
		home.logger.Println("Got error while parsing template: " + parseError.Error())
		response.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericError{Message: "Cannot parse home template: " + parseError.Error()}, response)
		return
	}
	executeError := template.Execute(response, home.data)
	if executeError != nil {
		home.logger.Println("Got error while executing template: " + executeError.Error())
		response.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&model.GenericError{Message: "Cannot run home template: " + executeError.Error()}, response)
		return
	}
}

// GetHealth is used to serve service health check
func (home *Home) GetHealth(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	utils.ToJSON(&model.Health{Status: "OK"}, response)
}
