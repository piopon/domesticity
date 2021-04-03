// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddEvent(params *AddEventParams) (*AddEventOK, error)

	DeleteEvent(params *DeleteEventParams) (*DeleteEventNoContent, error)

	GetEvent(params *GetEventParams) (*GetEventOK, error)

	GetEvents(params *GetEventsParams) (*GetEventsOK, error)

	UpdateEvent(params *UpdateEventParams) (*UpdateEventOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddEvent Creates a new event and adds it to DB
*/
func (a *Client) AddEvent(params *AddEventParams) (*AddEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addEvent",
		Method:             "POST",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteEvent Deletes an event from DB by specified ID parameter
*/
func (a *Client) DeleteEvent(params *DeleteEventParams) (*DeleteEventNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEvent",
		Method:             "DELETE",
		PathPattern:        "/events/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEventNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEvent Returns a event with provided id
*/
func (a *Client) GetEvent(params *GetEventParams) (*GetEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEvent",
		Method:             "GET",
		PathPattern:        "/event/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEvents Returns a list of currently stored events (all if no query params is used or filtered otherwise)
*/
func (a *Client) GetEvents(params *GetEventsParams) (*GetEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEvents",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEventsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateEvent Updates an event in DB by specified ID parameter
*/
func (a *Client) UpdateEvent(params *UpdateEventParams) (*UpdateEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEvent",
		Method:             "PUT",
		PathPattern:        "/events/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}