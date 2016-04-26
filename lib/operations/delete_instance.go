package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteInstanceHandlerFunc turns a function with the right signature into a delete instance handler
type DeleteInstanceHandlerFunc func(DeleteInstanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteInstanceHandlerFunc) Handle(params DeleteInstanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteInstanceHandler interface for that can handle valid delete instance params
type DeleteInstanceHandler interface {
	Handle(DeleteInstanceParams, interface{}) middleware.Responder
}

// NewDeleteInstance creates a new http.Handler for the delete instance operation
func NewDeleteInstance(ctx *middleware.Context, handler DeleteInstanceHandler) *DeleteInstance {
	return &DeleteInstance{Context: ctx, Handler: handler}
}

/*DeleteInstance swagger:route DELETE /instances/{instance_id} deleteInstance

Delete an instance

*/
type DeleteInstance struct {
	Context *middleware.Context
	Handler DeleteInstanceHandler
}

func (o *DeleteInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteInstanceParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
