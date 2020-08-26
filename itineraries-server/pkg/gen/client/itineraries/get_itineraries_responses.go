// Code generated by go-swagger; DO NOT EDIT.

package itineraries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/amadeusitgroup/miniplanes/itineraries-server/pkg/gen/models"
)

// GetItinerariesReader is a Reader for the GetItineraries structure.
type GetItinerariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItinerariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetItinerariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetItinerariesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetItinerariesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetItinerariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetItinerariesOK creates a GetItinerariesOK with default headers values
func NewGetItinerariesOK() *GetItinerariesOK {
	return &GetItinerariesOK{}
}

/*GetItinerariesOK handles this case with default header values.

list of itineraries
*/
type GetItinerariesOK struct {
	Payload []*models.Itinerary
}

func (o *GetItinerariesOK) Error() string {
	return fmt.Sprintf("[GET /itineraries][%d] getItinerariesOK  %+v", 200, o.Payload)
}

func (o *GetItinerariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItinerariesBadRequest creates a GetItinerariesBadRequest with default headers values
func NewGetItinerariesBadRequest() *GetItinerariesBadRequest {
	return &GetItinerariesBadRequest{}
}

/*GetItinerariesBadRequest handles this case with default header values.

generic error response
*/
type GetItinerariesBadRequest struct {
	Payload *models.Error
}

func (o *GetItinerariesBadRequest) Error() string {
	return fmt.Sprintf("[GET /itineraries][%d] getItinerariesBadRequest  %+v", 400, o.Payload)
}

func (o *GetItinerariesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItinerariesNotFound creates a GetItinerariesNotFound with default headers values
func NewGetItinerariesNotFound() *GetItinerariesNotFound {
	return &GetItinerariesNotFound{}
}

/*GetItinerariesNotFound handles this case with default header values.

not found
*/
type GetItinerariesNotFound struct {
	Payload *models.Error
}

func (o *GetItinerariesNotFound) Error() string {
	return fmt.Sprintf("[GET /itineraries][%d] getItinerariesNotFound  %+v", 404, o.Payload)
}

func (o *GetItinerariesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItinerariesDefault creates a GetItinerariesDefault with default headers values
func NewGetItinerariesDefault(code int) *GetItinerariesDefault {
	return &GetItinerariesDefault{
		_statusCode: code,
	}
}

/*GetItinerariesDefault handles this case with default header values.

generic error response
*/
type GetItinerariesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get itineraries default response
func (o *GetItinerariesDefault) Code() int {
	return o._statusCode
}

func (o *GetItinerariesDefault) Error() string {
	return fmt.Sprintf("[GET /itineraries][%d] GetItineraries default  %+v", o._statusCode, o.Payload)
}

func (o *GetItinerariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
