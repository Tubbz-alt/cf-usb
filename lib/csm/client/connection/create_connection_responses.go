package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/SUSE/cf-usb/lib/csm/models"
)

// CreateConnectionReader is a Reader for the CreateConnection structure.
type CreateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateConnectionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewCreateConnectionCreated creates a CreateConnectionCreated with default headers values
func NewCreateConnectionCreated() *CreateConnectionCreated {
	return &CreateConnectionCreated{}
}

/*CreateConnectionCreated handles this case with default header values.

create connection
*/
type CreateConnectionCreated struct {
	Payload *models.ServiceManagerConnectionResponse
}

func (o *CreateConnectionCreated) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace_id}/connections][%d] createConnectionCreated  %+v", 201, o.Payload)
}

func (o *CreateConnectionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceManagerConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectionDefault creates a CreateConnectionDefault with default headers values
func NewCreateConnectionDefault(code int) *CreateConnectionDefault {
	return &CreateConnectionDefault{
		_statusCode: code,
	}
}

/*CreateConnectionDefault handles this case with default header values.

generic error response
*/
type CreateConnectionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create connection default response
func (o *CreateConnectionDefault) Code() int {
	return o._statusCode
}

func (o *CreateConnectionDefault) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace_id}/connections][%d] createConnection default  %+v", o._statusCode, o.Payload)
}

func (o *CreateConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
