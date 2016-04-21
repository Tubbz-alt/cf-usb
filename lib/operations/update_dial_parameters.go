package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hpcloud/cf-usb/lib/genmodel"
)

// NewUpdateDialParams creates a new UpdateDialParams object
// with the default values initialized.
func NewUpdateDialParams() UpdateDialParams {
	var ()
	return UpdateDialParams{}
}

// UpdateDialParams contains all the bound params for the update dial operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateDial
type UpdateDialParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Updated dial
	  Required: true
	  In: body
	*/
	Dial *genmodel.Dial
	/*ID of the dial
	  Required: true
	  In: path
	*/
	DialID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateDialParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body genmodel.Dial
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("dial", "body"))
			} else {
				res = append(res, errors.NewParseError("dial", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Dial = &body
			}
		}

	} else {
		res = append(res, errors.Required("dial", "body"))
	}

	rDialID, rhkDialID, _ := route.Params.GetOK("dial_id")
	if err := o.bindDialID(rDialID, rhkDialID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDialParams) bindDialID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DialID = raw

	return nil
}
