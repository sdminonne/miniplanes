// Code generated by go-swagger; DO NOT EDIT.

package airports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAirportsParams creates a new GetAirportsParams object
// with the default values initialized.
func NewGetAirportsParams() *GetAirportsParams {

	return &GetAirportsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAirportsParamsWithTimeout creates a new GetAirportsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAirportsParamsWithTimeout(timeout time.Duration) *GetAirportsParams {

	return &GetAirportsParams{

		timeout: timeout,
	}
}

// NewGetAirportsParamsWithContext creates a new GetAirportsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAirportsParamsWithContext(ctx context.Context) *GetAirportsParams {

	return &GetAirportsParams{

		Context: ctx,
	}
}

// NewGetAirportsParamsWithHTTPClient creates a new GetAirportsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAirportsParamsWithHTTPClient(client *http.Client) *GetAirportsParams {

	return &GetAirportsParams{
		HTTPClient: client,
	}
}

/*GetAirportsParams contains all the parameters to send to the API endpoint
for the get airports operation typically these are written to a http.Request
*/
type GetAirportsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get airports params
func (o *GetAirportsParams) WithTimeout(timeout time.Duration) *GetAirportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get airports params
func (o *GetAirportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get airports params
func (o *GetAirportsParams) WithContext(ctx context.Context) *GetAirportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get airports params
func (o *GetAirportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get airports params
func (o *GetAirportsParams) WithHTTPClient(client *http.Client) *GetAirportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get airports params
func (o *GetAirportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAirportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
