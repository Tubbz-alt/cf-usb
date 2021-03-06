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

// DeleteConnectionReader is a Reader for the DeleteConnection structure.
type DeleteConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteConnectionOK creates a DeleteConnectionOK with default headers values
func NewDeleteConnectionOK() *DeleteConnectionOK {
	return &DeleteConnectionOK{}
}

/*DeleteConnectionOK handles this case with default header values.

delete connection
*/
type DeleteConnectionOK struct {
}

func (o *DeleteConnectionOK) Error() string {
	return fmt.Sprintf("[DELETE /workspaces/{workspace_id}/connections/{connection_id}][%d] deleteConnectionOK ", 200)
}

func (o *DeleteConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConnectionDefault creates a DeleteConnectionDefault with default headers values
func NewDeleteConnectionDefault(code int) *DeleteConnectionDefault {
	return &DeleteConnectionDefault{
		_statusCode: code,
	}
}

/*DeleteConnectionDefault handles this case with default header values.

generic error response
*/
type DeleteConnectionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete connection default response
func (o *DeleteConnectionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteConnectionDefault) Error() string {
	return fmt.Sprintf("[DELETE /workspaces/{workspace_id}/connections/{connection_id}][%d] deleteConnection default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
