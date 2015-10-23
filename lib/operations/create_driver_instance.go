package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// CreateDriverInstanceHandlerFunc turns a function with the right signature into a create driver instance handler
type CreateDriverInstanceHandlerFunc func(CreateDriverInstanceParams) (*genmodel.DriverInstance, error)

func (fn CreateDriverInstanceHandlerFunc) Handle(params CreateDriverInstanceParams) (*genmodel.DriverInstance, error) {
	return fn(params)
}

// CreateDriverInstanceHandler interface for that can handle valid create driver instance params
type CreateDriverInstanceHandler interface {
	Handle(CreateDriverInstanceParams) (*genmodel.DriverInstance, error)
}

// NewCreateDriverInstance creates a new http.Handler for the create driver instance operation
func NewCreateDriverInstance(ctx *middleware.Context, handler CreateDriverInstanceHandler) *CreateDriverInstance {
	return &CreateDriverInstance{Context: ctx, Handler: handler}
}

/*
Create a driver instance
*/
type CreateDriverInstance struct {
	Context *middleware.Context
	Params  CreateDriverInstanceParams
	Handler CreateDriverInstanceHandler
}

func (o *CreateDriverInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res, err := o.Handler.Handle(o.Params) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	o.Context.Respond(rw, r, route.Produces, route, res)

}
