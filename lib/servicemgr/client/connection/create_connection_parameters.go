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

	"github.com/hpcloud/cf-usb/lib/servicemgr/models"
)

// NewCreateConnectionParams creates a new CreateConnectionParams object
// with the default values initialized.
func NewCreateConnectionParams() *CreateConnectionParams {
	var ()
	return &CreateConnectionParams{}
}

/*CreateConnectionParams contains all the parameters to send to the API endpoint
for the create connection operation typically these are written to a http.Request
*/
type CreateConnectionParams struct {

	/*ConnectionCreateRequest
	  The service JSON you want to post

	*/
	ConnectionCreateRequest *models.ServiceManagerConnectionCreateRequest
	/*WorkspaceID
	  Workspace ID

	*/
	WorkspaceID string
}

// WithConnectionCreateRequest adds the connectionCreateRequest to the create connection params
func (o *CreateConnectionParams) WithConnectionCreateRequest(connectionCreateRequest *models.ServiceManagerConnectionCreateRequest) *CreateConnectionParams {
	o.ConnectionCreateRequest = connectionCreateRequest
	return o
}

// WithWorkspaceID adds the workspaceId to the create connection params
func (o *CreateConnectionParams) WithWorkspaceID(workspaceId string) *CreateConnectionParams {
	o.WorkspaceID = workspaceId
	return o
}

// WriteToRequest writes these params to a swagger request
<<<<<<< HEAD
func (o *CreateConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
=======
func (o *CreateConnectionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers

	var res []error

	if o.ConnectionCreateRequest == nil {
		o.ConnectionCreateRequest = new(models.ServiceManagerConnectionCreateRequest)
	}

	if err := r.SetBodyParam(o.ConnectionCreateRequest); err != nil {
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
