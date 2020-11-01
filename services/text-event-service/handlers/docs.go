package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	// empty docs package import so documentation is correctly generated
	_ "github.com/piopon/domesticity/services/text-event-service/docs"
)

// Docs is a handler struct used to handle documentation calls
type Docs struct {
	swaggerFile string
}

// NewDocs is a factory method to create Docs service handler with defined Swagger YAML file path
func NewDocs(swaggerFile string) *Docs {
	return &Docs{swaggerFile}
}

// GetDocumentation is used to serve doc handle
//
// swagger:route GET /docs utilities getDocumentation
// Returns a HTML and JS page with this documentation<br>
// NOTE: This page was auto-generated by ReDoc (https://github.com/Redocly/redoc)
// responses:
//  204: responseNoContent
func (docs *Docs) GetDocumentation(response http.ResponseWriter, request *http.Request) {
	options := middleware.RedocOpts{
		SpecURL: docs.swaggerFile,
		Title:   "Text Event Service API docs",
	}
	middleware.Redoc(options, nil).ServeHTTP(response, request)
}

// GetSwagger is used to retrieve Swagger YAML file
//
// swagger:route GET /scripts/swagger.yaml utilities getSwagger
// Returns swagger.yaml configuration file needed for generating this documentation<br>
// NOTE: This file should be located in scripts directory
// responses:
//  204: responseNoContent
func (docs *Docs) GetSwagger(response http.ResponseWriter, request *http.Request) {
	http.FileServer(http.Dir("")).ServeHTTP(response, request)
}
