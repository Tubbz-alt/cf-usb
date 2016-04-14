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

// GetConnectionReader is a Reader for the GetConnection structure.
type GetConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
<<<<<<< HEAD
func (o *GetConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
=======
func (o *GetConnectionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers
	switch response.Code() {

	case 200:
		result := NewGetConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetConnectionOK creates a GetConnectionOK with default headers values
func NewGetConnectionOK() *GetConnectionOK {
	return &GetConnectionOK{}
}

/*GetConnectionOK handles this case with default header values.

details of specified connection
*/
type GetConnectionOK struct {
	Payload *models.ServiceManagerConnectionResponse
}

func (o *GetConnectionOK) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspace_id}/connections/{connection_id}][%d] getConnectionOK  %+v", 200, o.Payload)
}

<<<<<<< HEAD
func (o *GetConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
=======
func (o *GetConnectionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers

	o.Payload = new(models.ServiceManagerConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionDefault creates a GetConnectionDefault with default headers values
func NewGetConnectionDefault(code int) *GetConnectionDefault {
	return &GetConnectionDefault{
		_statusCode: code,
	}
}

/*GetConnectionDefault handles this case with default header values.

generic error response
*/
type GetConnectionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get connection default response
func (o *GetConnectionDefault) Code() int {
	return o._statusCode
}

func (o *GetConnectionDefault) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspace_id}/connections/{connection_id}][%d] getConnection default  %+v", o._statusCode, o.Payload)
}

<<<<<<< HEAD
func (o *GetConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
=======
func (o *GetConnectionDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {
>>>>>>> f998b3c... [HCFRO-193] Use rest for calling drivers

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
