// Code generated by go-swagger; DO NOT EDIT.

package utilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new utilities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for utilities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetDocumentation(params *GetDocumentationParams) (*GetDocumentationOK, error)

	GetSwagger(params *GetSwaggerParams) (*GetSwaggerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDocumentation Returns a HTML and JS page with this documentation<br>
NOTE: This page was auto-generated by ReDoc (https://github.com/Redocly/redoc)
*/
func (a *Client) GetDocumentation(params *GetDocumentationParams) (*GetDocumentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDocumentationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDocumentation",
		Method:             "GET",
		PathPattern:        "/docs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDocumentationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDocumentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDocumentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSwagger Returns swagger.yaml configuration file needed for generating this documentation<br>
NOTE: This file should be located in resources directory
*/
func (a *Client) GetSwagger(params *GetSwaggerParams) (*GetSwaggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSwaggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSwagger",
		Method:             "GET",
		PathPattern:        "/resources/swagger.yaml",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSwaggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSwaggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSwagger: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
