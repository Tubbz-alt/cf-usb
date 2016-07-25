package brokermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Listing Listing.

swagger:model Listing
*/
type Listing struct {

	/* blurb
	 */
	Blurb string `json:"blurb,omitempty"`

	/* image URL.
	 */
	ImageURL string `json:"imageUrl,omitempty"`

	/* Long Description
	 */
	LongDescription string `json:"longDescription,omitempty"`
}

// Validate validates this listing
func (m *Listing) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}