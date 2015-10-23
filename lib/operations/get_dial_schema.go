package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// GetDialSchemaHandlerFunc turns a function with the right signature into a get dial schema handler
type GetDialSchemaHandlerFunc func(GetDialSchemaParams) (*genmodel.DialSchema, error)

func (fn GetDialSchemaHandlerFunc) Handle(params GetDialSchemaParams) (*genmodel.DialSchema, error) {
	return fn(params)
}

// GetDialSchemaHandler interface for that can handle valid get dial schema params
type GetDialSchemaHandler interface {
	Handle(GetDialSchemaParams) (*genmodel.DialSchema, error)
}

// NewGetDialSchema creates a new http.Handler for the get dial schema operation
func NewGetDialSchema(ctx *middleware.Context, handler GetDialSchemaHandler) *GetDialSchema {
	return &GetDialSchema{Context: ctx, Handler: handler}
}

/*
Get dial schema
*/
type GetDialSchema struct {
	Context *middleware.Context
	Params  GetDialSchemaParams
	Handler GetDialSchemaHandler
}

func (o *GetDialSchema) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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
