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

// NewDeprovisionServiceInstanceParams creates a new DeprovisionServiceInstanceParams object
// with the default values initialized.
func NewDeprovisionServiceInstanceParams() DeprovisionServiceInstanceParams {
	var ()
	return DeprovisionServiceInstanceParams{}
}

// DeprovisionServiceInstanceParams contains all the bound params for the deprovision service instance operation
// typically these are obtained from a http.Request
//
// swagger:parameters deprovisionServiceInstance
type DeprovisionServiceInstanceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The instance_id of a service instance is provided by the Cloud Controller. This ID will be used for future requests (bind and deprovision), so the broker must use it to correlate the resource it creates.
	  Required: true
	  In: path
	*/
	InstanceID string
	/*the plan id as stored in the catalog
	  Required: true
	  In: query
	*/
	PlanID string
	/*the service id as stored in the catalog
	  Required: true
	  In: query
	*/
	ServiceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeprovisionServiceInstanceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rInstanceID, rhkInstanceID, _ := route.Params.GetOK("instance_id")
	if err := o.bindInstanceID(rInstanceID, rhkInstanceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPlanID, qhkPlanID, _ := qs.GetOK("plan_id")
	if err := o.bindPlanID(qPlanID, qhkPlanID, route.Formats); err != nil {
		res = append(res, err)
	}

	qServiceID, qhkServiceID, _ := qs.GetOK("service_id")
	if err := o.bindServiceID(qServiceID, qhkServiceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeprovisionServiceInstanceParams) bindInstanceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.InstanceID = raw

	return nil
}

func (o *DeprovisionServiceInstanceParams) bindPlanID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("plan_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("plan_id", "query", raw); err != nil {
		return err
	}

	o.PlanID = raw

	return nil
}

func (o *DeprovisionServiceInstanceParams) bindServiceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("service_id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("service_id", "query", raw); err != nil {
		return err
	}

	o.ServiceID = raw

	return nil
}
