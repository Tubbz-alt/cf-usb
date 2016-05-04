package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UnregisterDriverInstanceHandlerFunc turns a function with the right signature into a unregister driver instance handler
type UnregisterDriverInstanceHandlerFunc func(UnregisterDriverInstanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UnregisterDriverInstanceHandlerFunc) Handle(params UnregisterDriverInstanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UnregisterDriverInstanceHandler interface for that can handle valid unregister driver instance params
type UnregisterDriverInstanceHandler interface {
	Handle(UnregisterDriverInstanceParams, interface{}) middleware.Responder
}

// NewUnregisterDriverInstance creates a new http.Handler for the unregister driver instance operation
func NewUnregisterDriverInstance(ctx *middleware.Context, handler UnregisterDriverInstanceHandler) *UnregisterDriverInstance {
	return &UnregisterDriverInstance{Context: ctx, Handler: handler}
}

/*UnregisterDriverInstance swagger:route DELETE /driver_endpoints/{driver_endpoint_id} unregisterDriverInstance

Unregisters a driver instance

*/
type UnregisterDriverInstance struct {
	Context *middleware.Context
	Handler UnregisterDriverInstanceHandler
}

func (o *UnregisterDriverInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewUnregisterDriverInstanceParams()

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
