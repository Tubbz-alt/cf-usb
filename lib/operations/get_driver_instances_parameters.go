package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

// GetDriverInstancesParams contains all the bound params for the get driver instances operation
// typically these are obtained from a http.Request
type GetDriverInstancesParams struct {
	// Authorization token
	Authorization string
	// Driver ID
	DriverID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetDriverInstancesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	qs := r.URL.Query()

	if err := o.bindAuthorization(r.Header.Get("authorization"), route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindDriverID(qs.Get("driver_id"), route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDriverInstancesParams) bindAuthorization(raw string, formats strfmt.Registry) error {
	if err := validate.RequiredString("authorization", "header", raw); err != nil {
		return err
	}

	o.Authorization = raw

	return nil
}

func (o *GetDriverInstancesParams) bindDriverID(raw string, formats strfmt.Registry) error {

	o.DriverID = raw

	return nil
}
