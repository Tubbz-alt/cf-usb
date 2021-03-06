package brokermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Binding Information to bind the service to an application.

swagger:model Binding
*/
type Binding struct {

	/* GUID of the application that you want to bind your service to. Will be included when users bind applications to service instances.
	 */
	AppGUID string `json:"app_guid,omitempty"`

	/* bind resource
	 */
	BindResource *BindResource `json:"bind_resource,omitempty"`

	/* parameters
	 */
	Parameters *Parameter `json:"parameters,omitempty"`

	/* ID of the plan from the catalog. While not strictly necessary, some brokers might make use of this ID.
	 */
	PlanID string `json:"plan_id,omitempty"`

	/* ID of the service from the catalog. While not strictly necessary, some brokers might make use of this ID.
	 */
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this binding
func (m *Binding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindResource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Binding) validateBindResource(formats strfmt.Registry) error {

	if swag.IsZero(m.BindResource) { // not required
		return nil
	}

	if m.BindResource != nil {

		if err := m.BindResource.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Binding) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {

		if err := m.Parameters.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
