package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// UpdateDriverConfigParams contains all the bound params for the update driver config operation
// typically these are obtained from a http.Request
type UpdateDriverConfigParams struct {
	// Add driver_config
	DriverConfig genmodel.DriverConfig
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateDriverConfigParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := route.Consumer.Consume(r.Body, &o.DriverConfig); err != nil {
		res = append(res, errors.NewParseError("driverConfig", "body", "", err))
	} else {
		if err := o.DriverConfig.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}