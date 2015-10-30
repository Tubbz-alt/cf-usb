package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// GetDriverInstanceHandlerFunc turns a function with the right signature into a get driver instance handler
type GetDriverInstanceHandlerFunc func(GetDriverInstanceParams) (*genmodel.DriverInstance, error)

func (fn GetDriverInstanceHandlerFunc) Handle(params GetDriverInstanceParams) (*genmodel.DriverInstance, error) {
	return fn(params)
}

// GetDriverInstanceHandler interface for that can handle valid get driver instance params
type GetDriverInstanceHandler interface {
	Handle(GetDriverInstanceParams) (*genmodel.DriverInstance, error)
}

// NewGetDriverInstance creates a new http.Handler for the get driver instance operation
func NewGetDriverInstance(ctx *middleware.Context, handler GetDriverInstanceHandler) *GetDriverInstance {
	return &GetDriverInstance{Context: ctx, Handler: handler}
}

/*
Gets driver configurations

*/
type GetDriverInstance struct {
	Context *middleware.Context
	Params  GetDriverInstanceParams
	Handler GetDriverInstanceHandler
}

func (o *GetDriverInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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