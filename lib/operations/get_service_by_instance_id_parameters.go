package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServiceByInstanceIDParams creates a new GetServiceByInstanceIDParams object
// with the default values initialized.
func NewGetServiceByInstanceIDParams() GetServiceByInstanceIDParams {
	var ()
	return GetServiceByInstanceIDParams{}
}

// GetServiceByInstanceIDParams contains all the bound params for the get service by instance Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getServiceByInstanceId
type GetServiceByInstanceIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Driver instance ID
	  Required: true
	  In: query
	*/
	DriverInstanceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetServiceByInstanceIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDriverInstanceID, qhkDriverInstanceID, _ := qs.GetOK("driver_instance_id")
	if err := o.bindDriverInstanceID(qDriverInstanceID, qhkDriverInstanceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceByInstanceIDParams) bindDriverInstanceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("driver_instance_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("driver_instance_id", "query", raw); err != nil {
		return err
	}

	o.DriverInstanceID = raw

	return nil
}
