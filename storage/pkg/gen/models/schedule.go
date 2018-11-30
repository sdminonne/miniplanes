// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Schedule schedule
// swagger:model schedule
type Schedule struct {

	// arrival
	Arrival string `json:"Arrival,omitempty"`

	// days operated
	DaysOperated string `json:"DaysOperated,omitempty"`

	// departure
	Departure string `json:"Departure,omitempty"`

	// destination
	Destination string `json:"Destination,omitempty"`

	// flight number
	FlightNumber string `json:"FlightNumber,omitempty"`

	// operating carrier
	OperatingCarrier string `json:"OperatingCarrier,omitempty"`

	// origin
	Origin string `json:"Origin,omitempty"`
}

// Validate validates this schedule
func (m *Schedule) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schedule) UnmarshalBinary(b []byte) error {
	var res Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
