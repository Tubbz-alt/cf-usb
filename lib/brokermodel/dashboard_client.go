package brokermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*DashboardClient Contains the data necessary to activate the [Dashboard SSO feature](https://docs.cloudfoundry.org/services/dashboard-sso.html) for this service

swagger:model DashboardClient
*/
type DashboardClient struct {

	/* The id of the Oauth2 client that the service intends to use. The name may be taken, in which case the API will return an error to the operator
	 */
	ID string `json:"id,omitempty"`

	/* A domain for the service dashboard that will be whitelisted by the UAA to enable SSO
	 */
	RedirectURI string `json:"redirect_uri,omitempty"`

	/* A secret for the dashboard client
	 */
	Secret string `json:"secret,omitempty"`
}

// Validate validates this dashboard client
func (m *DashboardClient) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
