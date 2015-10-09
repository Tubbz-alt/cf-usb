package genmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Service service

swagger:model service
*/
type Service struct {

	/* Bindable bindable
	 */
	Bindable bool `json:"bindable,omitempty"`

	/* Description description
	 */
	Description string `json:"description,omitempty"`

	/* DriverType driver type
	 */
	DriverType string `json:"driver_type,omitempty"`

	/* ID id

	Required: true
	*/
	ID string `json:"id"`

	/* Metadata metadata
	 */
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	/* Name name

	Required: true
	*/
	Name string `json:"name"`

	/* Plans plans
	 */
	Plans []Plan `json:"plans,omitempty"`

	/* Tags tags
	 */
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

func (m *Service) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}
