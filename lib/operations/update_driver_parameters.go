package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// UpdateDriverParams contains all the bound params for the update driver operation
// typically these are obtained from a http.Request
type UpdateDriverParams struct {
	// Authorization token
	Authorization string
	// Driver ID
	DriverID string
	// Driver to be updated
	Driver genmodel.Driver
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateDriverParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := o.bindAuthorization(r.Header.Get("authorization"), route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindDriverID(route.Params.Get("driver_id"), route.Formats); err != nil {
		res = append(res, err)
	}

	if err := route.Consumer.Consume(r.Body, &o.Driver); err != nil {
		res = append(res, errors.NewParseError("driver", "body", "", err))
	} else {
		if err := o.Driver.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDriverParams) bindAuthorization(raw string, formats strfmt.Registry) error {
	if err := validate.RequiredString("authorization", "header", raw); err != nil {
		return err
	}

	o.Authorization = raw

	return nil
}

func (o *UpdateDriverParams) bindDriverID(raw string, formats strfmt.Registry) error {

	o.DriverID = raw

	return nil
}
