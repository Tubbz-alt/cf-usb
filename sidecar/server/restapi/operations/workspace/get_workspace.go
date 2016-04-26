package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetWorkspaceHandlerFunc turns a function with the right signature into a get workspace handler
type GetWorkspaceHandlerFunc func(GetWorkspaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWorkspaceHandlerFunc) Handle(params GetWorkspaceParams) middleware.Responder {
	return fn(params)
}

// GetWorkspaceHandler interface for that can handle valid get workspace params
type GetWorkspaceHandler interface {
	Handle(GetWorkspaceParams) middleware.Responder
}

// NewGetWorkspace creates a new http.Handler for the get workspace operation
func NewGetWorkspace(ctx *middleware.Context, handler GetWorkspaceHandler) *GetWorkspace {
	return &GetWorkspace{Context: ctx, Handler: handler}
}

/*GetWorkspace swagger:route GET /workspaces/{workspace_id} workspace getWorkspace

Get the details for the specified

*/
type GetWorkspace struct {
	Context *middleware.Context
	Handler GetWorkspaceHandler
}

func (o *GetWorkspace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetWorkspaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}