package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

<<<<<<< HEAD
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
=======
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers

	"github.com/hpcloud/cf-usb/lib/servicemgr/models"
)

// CreateConnectionReader is a Reader for the CreateConnection structure.
type CreateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
<<<<<<< HEAD
func (o *CreateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
=======
func (o *CreateConnectionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers
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

<<<<<<< HEAD
func (o *CreateConnectionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
=======
func (o *CreateConnectionCreated) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers

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

<<<<<<< HEAD
func (o *CreateConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
=======
func (o *CreateConnectionDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
