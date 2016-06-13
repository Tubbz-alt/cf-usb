package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*StatusDiagnostic status diagnostic

swagger:model StatusDiagnostic
*/
type StatusDiagnostic struct {

	/* Full description of the diagnostic (e.g. This diagnostic attempts to connect to a MySQL Server using the configured credentials)

	Required: true
	*/
	Description *string `json:"description"`

	/* Contains any details that inform the success or failure of the diagnostic that was perfomed.

	Required: true
	*/
	Message *string `json:"message"`

	/* Name of the diagnostic (e.g. Credential verification)

	Required: true
	*/
	Name *string `json:"name"`

	/* Status of the diagnostic

	Required: true
	*/
	Status *string `json:"status"`
}

// Validate validates this status diagnostic
func (m *StatusDiagnostic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusDiagnostic) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *StatusDiagnostic) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *StatusDiagnostic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StatusDiagnostic) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}
