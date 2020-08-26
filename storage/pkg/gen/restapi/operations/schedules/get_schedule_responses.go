// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/amadeusitgroup/miniplanes/storage/pkg/gen/models"
)

// GetScheduleOKCode is the HTTP code returned for type GetScheduleOK
const GetScheduleOKCode int = 200

/*GetScheduleOK list of schedules

swagger:response getScheduleOK
*/
type GetScheduleOK struct {

	/*
	  In: Body
	*/
	Payload *models.Schedule `json:"body,omitempty"`
}

// NewGetScheduleOK creates GetScheduleOK with default headers values
func NewGetScheduleOK() *GetScheduleOK {

	return &GetScheduleOK{}
}

// WithPayload adds the payload to the get schedule o k response
func (o *GetScheduleOK) WithPayload(payload *models.Schedule) *GetScheduleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get schedule o k response
func (o *GetScheduleOK) SetPayload(payload *models.Schedule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScheduleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
