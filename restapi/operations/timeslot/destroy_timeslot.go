// Code generated by go-swagger; DO NOT EDIT.

package timeslot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DestroyTimeslotHandlerFunc turns a function with the right signature into a destroy timeslot handler
type DestroyTimeslotHandlerFunc func(DestroyTimeslotParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DestroyTimeslotHandlerFunc) Handle(params DestroyTimeslotParams) middleware.Responder {
	return fn(params)
}

// DestroyTimeslotHandler interface for that can handle valid destroy timeslot params
type DestroyTimeslotHandler interface {
	Handle(DestroyTimeslotParams) middleware.Responder
}

// NewDestroyTimeslot creates a new http.Handler for the destroy timeslot operation
func NewDestroyTimeslot(ctx *middleware.Context, handler DestroyTimeslotHandler) *DestroyTimeslot {
	return &DestroyTimeslot{Context: ctx, Handler: handler}
}

/* DestroyTimeslot swagger:route DELETE /users/{user_id}/time-slots/{time_slot_id} timeslot destroyTimeslot

Delete time slot according to given id

*/
type DestroyTimeslot struct {
	Context *middleware.Context
	Handler DestroyTimeslotHandler
}

func (o *DestroyTimeslot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDestroyTimeslotParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
