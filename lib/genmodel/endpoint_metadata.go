package genmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*EndpointMetadata endpoint metadata

swagger:model endpointMetadata
*/
type EndpointMetadata struct {

	/* The name of the service to be displayed in graphical clients

	 */
	DisplayName string `json:"display_name,omitempty"`

	/* Link to a documentation page for the service

	 */
	DocumentationURL string `json:"documentation_url,omitempty"`

	/* The URL to an image

	 */
	ImageURL string `json:"image_url,omitempty"`

	/* Long description

	 */
	LongDescription string `json:"long_description,omitempty"`

	/* The name of the upstream entity providing the actual service

	 */
	ProviderDisplayName string `json:"provider_display_name,omitempty"`

	/* Link to support information for the service

	 */
	SupportURL string `json:"support_url,omitempty"`
}

// Validate validates this endpoint metadata
func (m *EndpointMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}