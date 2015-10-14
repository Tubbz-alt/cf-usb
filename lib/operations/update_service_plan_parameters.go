package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// UpdateServicePlanParams contains all the bound params for the update service plan operation
// typically these are obtained from a http.Request
type UpdateServicePlanParams struct {
	// Update plan
	Plan genmodel.Plan
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateServicePlanParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := route.Consumer.Consume(r.Body, &o.Plan); err != nil {
		res = append(res, errors.NewParseError("plan", "body", "", err))
	} else {
		if err := o.Plan.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}