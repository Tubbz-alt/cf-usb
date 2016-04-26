package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServicePlansParams creates a new GetServicePlansParams object
// with the default values initialized.
func NewGetServicePlansParams() GetServicePlansParams {
	var ()
	return GetServicePlansParams{}
}

// GetServicePlansParams contains all the bound params for the get service plans operation
// typically these are obtained from a http.Request
//
// swagger:parameters getServicePlans
type GetServicePlansParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Instance ID
	  In: query
	*/
	InstanceID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetServicePlansParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qInstanceID, qhkInstanceID, _ := qs.GetOK("instance_id")
	if err := o.bindInstanceID(qInstanceID, qhkInstanceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServicePlansParams) bindInstanceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.InstanceID = &raw

	return nil
}
