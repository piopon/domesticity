package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/piopon/domesticity/services/text-event-service/utils"
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
func NewHome(logger *log.Logger, config *utils.Config) *Home {
	buildData := &BuildData{
		API:    VersionAPI,
		Time:   BuildTime,
		SHA:    CommitSHA,
		Golang: GoVersion,
	}
	return &Home{logger, &HomeData{buildData, config}}
}

// GetIndex is used to serve main page of service
func (home *Home) GetIndex(response http.ResponseWriter, request *http.Request) {
	template, parseError := template.ParseFiles("templates/index.html")
	if parseError != nil {
		fmt.Println("Got error while parsing template: " + parseError.Error())
		return
	}
	executeError := template.Execute(response, home.config)
	if executeError != nil {
		fmt.Println("Got error while executing template: " + executeError.Error())
	}
}
