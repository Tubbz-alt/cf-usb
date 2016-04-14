package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
<<<<<<< HEAD
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
=======
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers
)

// NewDeleteConnectionParams creates a new DeleteConnectionParams object
// with the default values initialized.
func NewDeleteConnectionParams() *DeleteConnectionParams {
	var ()
	return &DeleteConnectionParams{}
}

/*DeleteConnectionParams contains all the parameters to send to the API endpoint
for the delete connection operation typically these are written to a http.Request
*/
type DeleteConnectionParams struct {

	/*ConnectionID
	  connection ID

	*/
	ConnectionID string
	/*WorkspaceID
	  Workspace ID

	*/
	WorkspaceID string
}

// WithConnectionID adds the connectionId to the delete connection params
func (o *DeleteConnectionParams) WithConnectionID(connectionId string) *DeleteConnectionParams {
	o.ConnectionID = connectionId
	return o
}

// WithWorkspaceID adds the workspaceId to the delete connection params
func (o *DeleteConnectionParams) WithWorkspaceID(workspaceId string) *DeleteConnectionParams {
	o.WorkspaceID = workspaceId
	return o
}

// WriteToRequest writes these params to a swagger request
<<<<<<< HEAD
func (o *DeleteConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
=======
func (o *DeleteConnectionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers

	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID); err != nil {
		return err
	}

	// path param workspace_id
	if err := r.SetPathParam("workspace_id", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
