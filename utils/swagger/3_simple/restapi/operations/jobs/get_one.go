// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetOneHandlerFunc turns a function with the right signature into a get one handler
type GetOneHandlerFunc func(GetOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOneHandlerFunc) Handle(params GetOneParams) middleware.Responder {
	return fn(params)
}

// GetOneHandler interface for that can handle valid get one params
type GetOneHandler interface {
	Handle(GetOneParams) middleware.Responder
}

// NewGetOne creates a new http.Handler for the get one operation
func NewGetOne(ctx *middleware.Context, handler GetOneHandler) *GetOne {
	return &GetOne{Context: ctx, Handler: handler}
}

/*GetOne swagger:route GET /{id} jobs getOne

GetOne get one API

*/
type GetOne struct {
	Context *middleware.Context
	Handler GetOneHandler
}

func (o *GetOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
