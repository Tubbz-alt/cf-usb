package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// CreateServiceHandlerFunc turns a function with the right signature into a create service handler
type CreateServiceHandlerFunc func(CreateServiceParams) error

func (fn CreateServiceHandlerFunc) Handle(params CreateServiceParams) error {
	return fn(params)
}

// CreateServiceHandler interface for that can handle valid create service params
type CreateServiceHandler interface {
	Handle(CreateServiceParams) error
}

// NewCreateService creates a new http.Handler for the create service operation
func NewCreateService(ctx *middleware.Context, handler CreateServiceHandler) *CreateService {
	return &CreateService{Context: ctx, Handler: handler}
}

/*
Create a new `service` definition
*/
type CreateService struct {
	Context *middleware.Context
	Params  CreateServiceParams
	Handler CreateServiceHandler
}

func (o *CreateService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	err := o.Handler.Handle(o.Params) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	o.Context.Respond(rw, r, route.Produces, route, nil)

}