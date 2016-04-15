package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"mysql-driver/models"
)

// NewCreateConnectionParams creates a new CreateConnectionParams object
// with the default values initialized.
func NewCreateConnectionParams() CreateConnectionParams {
	var ()
	return CreateConnectionParams{}
}

// CreateConnectionParams contains all the bound params for the create connection operation
// typically these are obtained from a http.Request
//
// swagger:parameters createConnection
type CreateConnectionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The service JSON you want to post
	  Required: true
	  In: body
	*/
	ConnectionCreateRequest *models.ServiceManagerConnectionCreateRequest
	/*Workspace ID
	  Required: true
	  In: path
	*/
	WorkspaceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateConnectionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if httpkit.HasBody(r) {
		defer r.Body.Close()
		var body models.ServiceManagerConnectionCreateRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("connectionCreateRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("connectionCreateRequest", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ConnectionCreateRequest = &body
			}
		}

	} else {
		res = append(res, errors.Required("connectionCreateRequest", "body"))
	}

	rWorkspaceID, rhkWorkspaceID, _ := route.Params.GetOK("workspace_id")
	if err := o.bindWorkspaceID(rWorkspaceID, rhkWorkspaceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateConnectionParams) bindWorkspaceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.WorkspaceID = raw

	return nil
}
