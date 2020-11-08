// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/piopon/domesticity/services/text-event-service/sdk/models"
)

// AddEventReader is a Reader for the AddEvent structure.
type AddEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAddEventUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddEventInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddEventOK creates a AddEventOK with default headers values
func NewAddEventOK() *AddEventOK {
	return &AddEventOK{}
}

/*AddEventOK handles this case with default header values.

Response with currently created event which was added to DB
*/
type AddEventOK struct {
	Payload *models.Event
}

func (o *AddEventOK) Error() string {
	return fmt.Sprintf("[POST /events][%d] addEventOK  %+v", 200, o.Payload)
}

func (o *AddEventOK) GetPayload() *models.Event {
	return o.Payload
}

func (o *AddEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEventBadRequest creates a AddEventBadRequest with default headers values
func NewAddEventBadRequest() *AddEventBadRequest {
	return &AddEventBadRequest{}
}

/*AddEventBadRequest handles this case with default header values.

Error with JSON message returned in body when bad request was invoked
*/
type AddEventBadRequest struct {
	Payload *models.GenericError
}

func (o *AddEventBadRequest) Error() string {
	return fmt.Sprintf("[POST /events][%d] addEventBadRequest  %+v", 400, o.Payload)
}

func (o *AddEventBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *AddEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEventUnprocessableEntity creates a AddEventUnprocessableEntity with default headers values
func NewAddEventUnprocessableEntity() *AddEventUnprocessableEntity {
	return &AddEventUnprocessableEntity{}
}

/*AddEventUnprocessableEntity handles this case with default header values.

Error with JSON messages returned in body when validation error occurs
*/
type AddEventUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *AddEventUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /events][%d] addEventUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AddEventUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AddEventUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEventInternalServerError creates a AddEventInternalServerError with default headers values
func NewAddEventInternalServerError() *AddEventInternalServerError {
	return &AddEventInternalServerError{}
}

/*AddEventInternalServerError handles this case with default header values.

Error with JSON message returned in body when internal server error occurs
*/
type AddEventInternalServerError struct {
	Payload *models.GenericError
}

func (o *AddEventInternalServerError) Error() string {
	return fmt.Sprintf("[POST /events][%d] addEventInternalServerError  %+v", 500, o.Payload)
}

func (o *AddEventInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *AddEventInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}