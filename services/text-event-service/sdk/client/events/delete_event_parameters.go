// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteEventParams creates a new DeleteEventParams object
// with the default values initialized.
func NewDeleteEventParams() *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEventParamsWithTimeout creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEventParamsWithTimeout(timeout time.Duration) *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		timeout: timeout,
	}
}

// NewDeleteEventParamsWithContext creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEventParamsWithContext(ctx context.Context) *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		Context: ctx,
	}
}

// NewDeleteEventParamsWithHTTPClient creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEventParamsWithHTTPClient(client *http.Client) *DeleteEventParams {
	var ()
	return &DeleteEventParams{
		HTTPClient: client,
	}
}

/*DeleteEventParams contains all the parameters to send to the API endpoint
for the delete event operation typically these are written to a http.Request
*/
type DeleteEventParams struct {

	/*ID
	  The id of the event for which the operation relates<br>
	NOTE: The type is primitive.ObjectID which is the BSON ObjectID type

	*/
	ID []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete event params
func (o *DeleteEventParams) WithTimeout(timeout time.Duration) *DeleteEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete event params
func (o *DeleteEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete event params
func (o *DeleteEventParams) WithContext(ctx context.Context) *DeleteEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete event params
func (o *DeleteEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete event params
func (o *DeleteEventParams) WithHTTPClient(client *http.Client) *DeleteEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete event params
func (o *DeleteEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete event params
func (o *DeleteEventParams) WithID(id []int64) *DeleteEventParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete event params
func (o *DeleteEventParams) SetID(id []int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesID []string
	for _, v := range o.ID {
		valuesID = append(valuesID, swag.FormatInt64(v))
	}

	joinedID := swag.JoinByFormat(valuesID, "")
	// path array param id
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedID) > 0 {
		if err := r.SetPathParam("id", joinedID[0]); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
