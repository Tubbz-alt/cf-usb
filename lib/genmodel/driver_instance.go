package genmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DriverInstance driver instance

swagger:model driver_instance
*/
type DriverInstance struct {

	/* configuration
	 */
	Configuration interface{} `json:"configuration,omitempty"`

	/* dials
	 */
	Dials []string `json:"dials,omitempty"`

	/* driver id

	Required: true
	Max Length: 36
	Min Length: 36
	*/
	DriverID *string `json:"driver_id"`

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* name

	Required: true
	*/
	Name *string `json:"name"`

	/* service
	 */
<<<<<<< HEAD
	Service string `json:"service,omitempty"`

	/* target URL
	 */
	TargetURL string `json:"targetURL,omitempty"`
=======
	Service *string `json:"service,omitempty"`

	TargetURL string `json:"target,omitempty"`
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers
}

// Validate validates this driver instance
func (m *DriverInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDials(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDriverID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DriverInstance) validateDials(formats strfmt.Registry) error {

	if swag.IsZero(m.Dials) { // not required
		return nil
	}

	return nil
}

func (m *DriverInstance) validateDriverID(formats strfmt.Registry) error {

	if err := validate.Required("driver_id", "body", m.DriverID); err != nil {
		return err
	}

	if err := validate.MinLength("driver_id", "body", string(*m.DriverID), 36); err != nil {
		return err
	}

	if err := validate.MaxLength("driver_id", "body", string(*m.DriverID), 36); err != nil {
		return err
	}

	return nil
}

func (m *DriverInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
