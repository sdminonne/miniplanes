// Code generated by go-swagger; DO NOT EDIT.

package airports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddAirportHandlerFunc turns a function with the right signature into a add airport handler
type AddAirportHandlerFunc func(AddAirportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddAirportHandlerFunc) Handle(params AddAirportParams) middleware.Responder {
	return fn(params)
}

// AddAirportHandler interface for that can handle valid add airport params
type AddAirportHandler interface {
	Handle(AddAirportParams) middleware.Responder
}

// NewAddAirport creates a new http.Handler for the add airport operation
func NewAddAirport(ctx *middleware.Context, handler AddAirportHandler) *AddAirport {
	return &AddAirport{Context: ctx, Handler: handler}
}

/*AddAirport swagger:route POST /airports airports addAirport

Creates an Airport entry. Duplicates are not allowed

*/
type AddAirport struct {
	Context *middleware.Context
	Handler AddAirportHandler
}

func (o *AddAirport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddAirportParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
