package brokermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*UnbindParameters Parameters needed to unbind a service instance

swagger:model UnbindParameters
*/
type UnbindParameters struct {

	/* ID of the plan from the catalog. While not strictly necessary, some brokers might make use of this ID.
	 */
	PlanID string `json:"plan_id,omitempty"`

	/* ID of the service from the catalog. While not strictly necessary, some brokers might make use of this ID.
	 */
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this unbind parameters
func (m *UnbindParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}