package genmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*dial Dial dial

swagger:model dial
*/
type Dial struct {

	/* Configuration configuration
	 */
	Configuration interface{} `json:"configuration,omitempty"`

	/* DriverInstanceID driver instance id

	Required: true
	Max Length: 36
	Min Length: 36
	*/
	DriverInstanceID string `json:"driver_instance_id,omitempty"`

	/* ID id
	 */
	ID *string `json:"id,omitempty"`

	/* Plan plan
	 */
	Plan *string `json:"plan,omitempty"`
}

// Validate validates this dial
func (m *Dial) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriverInstanceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dial) validateDriverInstanceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("driver_instance_id", "body", string(m.DriverInstanceID)); err != nil {
		return err
	}

	if err := validate.MinLength("driver_instance_id", "body", string(m.DriverInstanceID), 36); err != nil {
		return err
	}

	if err := validate.MaxLength("driver_instance_id", "body", string(m.DriverInstanceID), 36); err != nil {
		return err
	}

	return nil
}
