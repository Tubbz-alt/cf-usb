package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateInstanceHandlerFunc turns a function with the right signature into a update instance handler
type UpdateInstanceHandlerFunc func(UpdateInstanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateInstanceHandlerFunc) Handle(params UpdateInstanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateInstanceHandler interface for that can handle valid update instance params
type UpdateInstanceHandler interface {
	Handle(UpdateInstanceParams, interface{}) middleware.Responder
}

// NewUpdateInstance creates a new http.Handler for the update instance operation
func NewUpdateInstance(ctx *middleware.Context, handler UpdateInstanceHandler) *UpdateInstance {
	return &UpdateInstance{Context: ctx, Handler: handler}
}

/*UpdateInstance swagger:route PUT /instances/{instance_id} updateInstance

Update an instance


*/
type UpdateInstance struct {
	Context *middleware.Context
	Handler UpdateInstanceHandler
}

func (o *UpdateInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewUpdateInstanceParams()

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
