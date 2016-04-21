package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDriverInstanceParams creates a new DeleteDriverInstanceParams object
// with the default values initialized.
func NewDeleteDriverInstanceParams() DeleteDriverInstanceParams {
	var ()
	return DeleteDriverInstanceParams{}
}

// DeleteDriverInstanceParams contains all the bound params for the delete driver instance operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteDriverInstance
type DeleteDriverInstanceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Driver Instance ID
	  Required: true
	  In: path
	*/
	DriverInstanceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteDriverInstanceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rDriverInstanceID, rhkDriverInstanceID, _ := route.Params.GetOK("driver_instance_id")
	if err := o.bindDriverInstanceID(rDriverInstanceID, rhkDriverInstanceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteDriverInstanceParams) bindDriverInstanceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DriverInstanceID = raw

	return nil
}
