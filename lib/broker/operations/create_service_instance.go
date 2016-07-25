package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateServiceInstanceHandlerFunc turns a function with the right signature into a create service instance handler
type CreateServiceInstanceHandlerFunc func(CreateServiceInstanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateServiceInstanceHandlerFunc) Handle(params CreateServiceInstanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateServiceInstanceHandler interface for that can handle valid create service instance params
type CreateServiceInstanceHandler interface {
	Handle(CreateServiceInstanceParams, interface{}) middleware.Responder
}

// NewCreateServiceInstance creates a new http.Handler for the create service instance operation
func NewCreateServiceInstance(ctx *middleware.Context, handler CreateServiceInstanceHandler) *CreateServiceInstance {
	return &CreateServiceInstance{Context: ctx, Handler: handler}
}

/*CreateServiceInstance swagger:route PUT /service_instances/{instance_id} createServiceInstance

Provisions a service instance

When the broker receives a provision request from Cloud Controller, it should synchronously take whatever action is necessary to create a new service resource for the developer. The result of provisioning varies by service type, although there are a few common actions that work for many services.

*/
type CreateServiceInstance struct {
	Context *middleware.Context
	Handler CreateServiceInstanceHandler
}

func (o *CreateServiceInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateServiceInstanceParams()

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