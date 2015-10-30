package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// GetServicePlanHandlerFunc turns a function with the right signature into a get service plan handler
type GetServicePlanHandlerFunc func(GetServicePlanParams) (*genmodel.Plan, error)

func (fn GetServicePlanHandlerFunc) Handle(params GetServicePlanParams) (*genmodel.Plan, error) {
	return fn(params)
}

// GetServicePlanHandler interface for that can handle valid get service plan params
type GetServicePlanHandler interface {
	Handle(GetServicePlanParams) (*genmodel.Plan, error)
}

// NewGetServicePlan creates a new http.Handler for the get service plan operation
func NewGetServicePlan(ctx *middleware.Context, handler GetServicePlanHandler) *GetServicePlan {
	return &GetServicePlan{Context: ctx, Handler: handler}
}

/*
Gets the `plan` with the **planID** for the *serviceID*
*/
type GetServicePlan struct {
	Context *middleware.Context
	Params  GetServicePlanParams
	Handler GetServicePlanHandler
}

func (o *GetServicePlan) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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