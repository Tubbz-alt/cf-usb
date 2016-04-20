package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateWorkspaceHandlerFunc turns a function with the right signature into a create workspace handler
type CreateWorkspaceHandlerFunc func(CreateWorkspaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateWorkspaceHandlerFunc) Handle(params CreateWorkspaceParams) middleware.Responder {
	return fn(params)
}

// CreateWorkspaceHandler interface for that can handle valid create workspace params
type CreateWorkspaceHandler interface {
	Handle(CreateWorkspaceParams) middleware.Responder
}

// NewCreateWorkspace creates a new http.Handler for the create workspace operation
func NewCreateWorkspace(ctx *middleware.Context, handler CreateWorkspaceHandler) *CreateWorkspace {
	return &CreateWorkspace{Context: ctx, Handler: handler}
}

/*CreateWorkspace swagger:route POST /workspaces workspace createWorkspace

Create new workspace

*/
type CreateWorkspace struct {
	Context *middleware.Context
	Handler CreateWorkspaceHandler
}

func (o *CreateWorkspace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateWorkspaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
