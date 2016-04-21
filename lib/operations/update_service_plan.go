package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateServicePlanHandlerFunc turns a function with the right signature into a update service plan handler
type UpdateServicePlanHandlerFunc func(UpdateServicePlanParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateServicePlanHandlerFunc) Handle(params UpdateServicePlanParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateServicePlanHandler interface for that can handle valid update service plan params
type UpdateServicePlanHandler interface {
	Handle(UpdateServicePlanParams, interface{}) middleware.Responder
}

// NewUpdateServicePlan creates a new http.Handler for the update service plan operation
func NewUpdateServicePlan(ctx *middleware.Context, handler UpdateServicePlanHandler) *UpdateServicePlan {
	return &UpdateServicePlan{Context: ctx, Handler: handler}
}

/*UpdateServicePlan swagger:route PUT /plans/{plan_id} updateServicePlan

Updates the plan with the id **planID** for the service id **serviceID**

*/
type UpdateServicePlan struct {
	Context *middleware.Context
	Handler UpdateServicePlanHandler
}

func (o *UpdateServicePlan) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewUpdateServicePlanParams()

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
