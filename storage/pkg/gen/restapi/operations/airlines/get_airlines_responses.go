// Code generated by go-swagger; DO NOT EDIT.

package airlines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/amadeusitgroup/miniapp/storage/pkg/gen/models"
)

// GetAirlinesOKCode is the HTTP code returned for type GetAirlinesOK
const GetAirlinesOKCode int = 200

/*GetAirlinesOK list of airlines

swagger:response getAirlinesOK
*/
type GetAirlinesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Airline `json:"body,omitempty"`
}

// NewGetAirlinesOK creates GetAirlinesOK with default headers values
func NewGetAirlinesOK() *GetAirlinesOK {

	return &GetAirlinesOK{}
}

// WithPayload adds the payload to the get airlines o k response
func (o *GetAirlinesOK) WithPayload(payload []*models.Airline) *GetAirlinesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get airlines o k response
func (o *GetAirlinesOK) SetPayload(payload []*models.Airline) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAirlinesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Airline, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetAirlinesBadRequestCode is the HTTP code returned for type GetAirlinesBadRequest
const GetAirlinesBadRequestCode int = 400

/*GetAirlinesBadRequest generic error response

swagger:response getAirlinesBadRequest
*/
type GetAirlinesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAirlinesBadRequest creates GetAirlinesBadRequest with default headers values
func NewGetAirlinesBadRequest() *GetAirlinesBadRequest {

	return &GetAirlinesBadRequest{}
}

// WithPayload adds the payload to the get airlines bad request response
func (o *GetAirlinesBadRequest) WithPayload(payload *models.Error) *GetAirlinesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get airlines bad request response
func (o *GetAirlinesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAirlinesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAirlinesDefault generic error response

swagger:response getAirlinesDefault
*/
type GetAirlinesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAirlinesDefault creates GetAirlinesDefault with default headers values
func NewGetAirlinesDefault(code int) *GetAirlinesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAirlinesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get airlines default response
func (o *GetAirlinesDefault) WithStatusCode(code int) *GetAirlinesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get airlines default response
func (o *GetAirlinesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get airlines default response
func (o *GetAirlinesDefault) WithPayload(payload *models.Error) *GetAirlinesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get airlines default response
func (o *GetAirlinesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAirlinesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
