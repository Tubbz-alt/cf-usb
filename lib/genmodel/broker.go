package genmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Broker broker

swagger:model broker
*/
type Broker struct {

	/* Credentials credentials

	Required: true
	*/
	Credentials struct {

		/* Password password

		Required: true
		*/
		Password string `json:"password"`

		/* Username username

		Required: true
		*/
		Username string `json:"username"`
	} `json:"credentials"`

	/* Listen listen

	Required: true
	*/
	Listen string `json:"listen"`
}

// Validate validates this broker
func (m *Broker) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListen(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

func (m *Broker) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials"+"."+"password", "body", string(m.Credentials.Password)); err != nil {
		return err
	}

	if err := validate.Required("credentials"+"."+"username", "body", string(m.Credentials.Username)); err != nil {
		return err
	}

	return nil
}

func (m *Broker) validateListen(formats strfmt.Registry) error {

	if err := validate.Required("listen", "body", string(m.Listen)); err != nil {
		return err
	}

	return nil
}
