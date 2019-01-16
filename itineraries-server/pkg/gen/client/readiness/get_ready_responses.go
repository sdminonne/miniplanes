// Code generated by go-swagger; DO NOT EDIT.

package readiness

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/amadeusitgroup/miniapp/itineraries-server/pkg/gen/models"
)

// GetReadyReader is a Reader for the GetReady structure.
type GetReadyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReadyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReadyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 503:
		result := NewGetReadyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetReadyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReadyOK creates a GetReadyOK with default headers values
func NewGetReadyOK() *GetReadyOK {
	return &GetReadyOK{}
}

/*GetReadyOK handles this case with default header values.

readiness probe
*/
type GetReadyOK struct {
}

func (o *GetReadyOK) Error() string {
	return fmt.Sprintf("[GET /ready][%d] getReadyOK ", 200)
}

func (o *GetReadyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReadyServiceUnavailable creates a GetReadyServiceUnavailable with default headers values
func NewGetReadyServiceUnavailable() *GetReadyServiceUnavailable {
	return &GetReadyServiceUnavailable{}
}

/*GetReadyServiceUnavailable handles this case with default header values.

if not ready
*/
type GetReadyServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetReadyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ready][%d] getReadyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReadyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReadyDefault creates a GetReadyDefault with default headers values
func NewGetReadyDefault(code int) *GetReadyDefault {
	return &GetReadyDefault{
		_statusCode: code,
	}
}

/*GetReadyDefault handles this case with default header values.

generic error response
*/
type GetReadyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get ready default response
func (o *GetReadyDefault) Code() int {
	return o._statusCode
}

func (o *GetReadyDefault) Error() string {
	return fmt.Sprintf("[GET /ready][%d] GetReady default  %+v", o._statusCode, o.Payload)
}

func (o *GetReadyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
