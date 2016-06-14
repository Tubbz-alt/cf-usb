package brokermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*PreviousValues Information about the instance prior to the update.

swagger:model PreviousValues
*/
type PreviousValues struct {

	/* ID of the organization containing the instance.
	 */
	OrganizationID string `json:"organization_id,omitempty"`

	/* ID of the plan prior to the update.
	 */
	PlanID string `json:"plan_id,omitempty"`

	/* ID of the service for the instance.
	 */
	ServiceID string `json:"service_id,omitempty"`

	/* ID of the space containing the instance.
	 */
	SpaceID string `json:"space_id,omitempty"`
}

// Validate validates this previous values
func (m *PreviousValues) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
