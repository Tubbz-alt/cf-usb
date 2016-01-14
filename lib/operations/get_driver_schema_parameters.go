package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetDriverSchemaParams creates a new GetDriverSchemaParams object
// with the default values initialized.
func NewGetDriverSchemaParams() GetDriverSchemaParams {
	var ()
	return GetDriverSchemaParams{}
}

// GetDriverSchemaParams contains all the bound params for the get driver schema operation
// typically these are obtained from a http.Request
//
// swagger:parameters getDriverSchema
type GetDriverSchemaParams struct {
	/*Driver ID
	  Required: true
	  In: path
	*/
	DriverID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetDriverSchemaParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	rDriverID, rhkDriverID, _ := route.Params.GetOK("driver_id")
	if err := o.bindDriverID(rDriverID, rhkDriverID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDriverSchemaParams) bindDriverID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DriverID = raw

	return nil
}
