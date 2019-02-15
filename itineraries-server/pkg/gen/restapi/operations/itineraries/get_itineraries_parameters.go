// Code generated by go-swagger; DO NOT EDIT.

package itineraries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetItinerariesParams creates a new GetItinerariesParams object
// with the default values initialized.
func NewGetItinerariesParams() GetItinerariesParams {

	var (
		// initialize parameters with default values

		departureTimeDefault = string("1200")

		returnTimeDefault = string("1200")
	)

	return GetItinerariesParams{
		DepartureTime: &departureTimeDefault,

		ReturnTime: &returnTimeDefault,
	}
}

// GetItinerariesParams contains all the bound params for the get itineraries operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetItineraries
type GetItinerariesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	DepartureDate *string
	/*
	  In: query
	  Default: "1200"
	*/
	DepartureTime *string
	/*
	  In: query
	*/
	From *string
	/*
	  In: query
	*/
	ReturnDate *string
	/*
	  In: query
	  Default: "1200"
	*/
	ReturnTime *string
	/*
	  In: query
	*/
	To *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetItinerariesParams() beforehand.
func (o *GetItinerariesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDepartureDate, qhkDepartureDate, _ := qs.GetOK("departureDate")
	if err := o.bindDepartureDate(qDepartureDate, qhkDepartureDate, route.Formats); err != nil {
		res = append(res, err)
	}

	qDepartureTime, qhkDepartureTime, _ := qs.GetOK("departureTime")
	if err := o.bindDepartureTime(qDepartureTime, qhkDepartureTime, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qReturnDate, qhkReturnDate, _ := qs.GetOK("returnDate")
	if err := o.bindReturnDate(qReturnDate, qhkReturnDate, route.Formats); err != nil {
		res = append(res, err)
	}

	qReturnTime, qhkReturnTime, _ := qs.GetOK("returnTime")
	if err := o.bindReturnTime(qReturnTime, qhkReturnTime, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDepartureDate binds and validates parameter DepartureDate from query.
func (o *GetItinerariesParams) bindDepartureDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DepartureDate = &raw

	return nil
}

// bindDepartureTime binds and validates parameter DepartureTime from query.
func (o *GetItinerariesParams) bindDepartureTime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetItinerariesParams()
		return nil
	}

	o.DepartureTime = &raw

	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetItinerariesParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.From = &raw

	return nil
}

// bindReturnDate binds and validates parameter ReturnDate from query.
func (o *GetItinerariesParams) bindReturnDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ReturnDate = &raw

	return nil
}

// bindReturnTime binds and validates parameter ReturnTime from query.
func (o *GetItinerariesParams) bindReturnTime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetItinerariesParams()
		return nil
	}

	o.ReturnTime = &raw

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetItinerariesParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.To = &raw

	return nil
}
