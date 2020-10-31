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
func (docs *Docs) GetDocumentation(response http.ResponseWriter, request *http.Request) {
	middleware.Redoc(middleware.RedocOpts{SpecURL: docs.swaggerFile}, nil).ServeHTTP(response, request)
}

// GetSwagger is used to retrieve Swagger YAML file
func (docs *Docs) GetSwagger(response http.ResponseWriter, request *http.Request) {
	http.FileServer(http.Dir("")).ServeHTTP(response, request)
}
