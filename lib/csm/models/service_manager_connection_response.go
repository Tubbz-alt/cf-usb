package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ServiceManagerConnectionResponse service manager connection response

swagger:model ServiceManagerConnectionResponse
*/
type ServiceManagerConnectionResponse struct {

	/* key value map with connection details (Service manager won't interpret the details of the map it will send this back to the requester as is)
	 */
	Details map[string]interface{} `json:"details,omitempty"`

	/* Processing type

	Required: true
	*/
	ProcessingType *string `json:"processing_type"`

	/* status

	Required: true
	*/
	Status *string `json:"status"`
}

// Validate validates this service manager connection response
func (m *ServiceManagerConnectionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcessingType(formats); err != nil {
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

func (m *ServiceManagerConnectionResponse) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	return nil
}

var serviceManagerConnectionResponseTypeProcessingTypePropEnum []interface{}

// prop value enum
func (m *ServiceManagerConnectionResponse) validateProcessingTypeEnum(path, location string, value string) error {
	if serviceManagerConnectionResponseTypeProcessingTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["none","default","extension"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serviceManagerConnectionResponseTypeProcessingTypePropEnum = append(serviceManagerConnectionResponseTypeProcessingTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serviceManagerConnectionResponseTypeProcessingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceManagerConnectionResponse) validateProcessingType(formats strfmt.Registry) error {

	if err := validate.Required("processing_type", "body", m.ProcessingType); err != nil {
		return err
	}

	// value enum
	if err := m.validateProcessingTypeEnum("processing_type", "body", *m.ProcessingType); err != nil {
		return err
	}

	return nil
}

var serviceManagerConnectionResponseTypeStatusPropEnum []interface{}

// prop value enum
func (m *ServiceManagerConnectionResponse) validateStatusEnum(path, location string, value string) error {
	if serviceManagerConnectionResponseTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["none","unknown","successful","failed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serviceManagerConnectionResponseTypeStatusPropEnum = append(serviceManagerConnectionResponseTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serviceManagerConnectionResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceManagerConnectionResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}
