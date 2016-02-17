package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// UpdateDialHandlerFunc turns a function with the right signature into a update dial handler
type UpdateDialHandlerFunc func(UpdateDialParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateDialHandlerFunc) Handle(params UpdateDialParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateDialHandler interface for that can handle valid update dial params
type UpdateDialHandler interface {
	Handle(UpdateDialParams, interface{}) middleware.Responder
}

// NewUpdateDial creates a new http.Handler for the update dial operation
func NewUpdateDial(ctx *middleware.Context, handler UpdateDialHandler) *UpdateDial {
	return &UpdateDial{Context: ctx, Handler: handler}
}

/*UpdateDial swagger:route PUT /dials/{dial_id} updateDial

Updates the dial with the id **dial_id**

*/
type UpdateDial struct {
	Context *middleware.Context
	Params  UpdateDialParams
	Handler UpdateDialHandler
}

func (o *UpdateDial) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewUpdateDialParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
