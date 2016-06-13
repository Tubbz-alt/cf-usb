package genmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DriverEndpoint driver endpoint

swagger:model driverEndpoint
*/
type DriverEndpoint struct {

	/* An authentication key used by the USB when communicating with the
	driver endpoint.

	*/
	AuthenticationKey string `json:"authenticationKey,omitempty"`

	/* URL for the driver endpoint. Used by the USB to create service
	instances, generate credentials, discover plans and schemas.

	*/
	EndpointURL string `json:"endpointURL,omitempty"`

	/* USB generated ID for the driver endpoint.

	 */
	ID string `json:"id,omitempty"`

	/* Optional metadata configuration used by graphical clients to display
	information about a service.

	*/
	Metadata *EndpointMetadata `json:"metadata,omitempty"`

	/* The name of the driver endpoint. It's displayed by the Cloud Foundry
	CLI when the user lists available service offerings.


	Required: true
	*/
	Name *string `json:"name"`
}

// Validate validates this driver endpoint
func (m *DriverEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DriverEndpoint) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
