package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/sidecar/server/models"
)

/*DeleteConnectionOK delete connection

swagger:response deleteConnectionOK
*/
type DeleteConnectionOK struct {
}

// NewDeleteConnectionOK creates DeleteConnectionOK with default headers values
func NewDeleteConnectionOK() *DeleteConnectionOK {
	return &DeleteConnectionOK{}
}

// WriteResponse to the client
func (o *DeleteConnectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*DeleteConnectionDefault generic error response

swagger:response deleteConnectionDefault
*/
type DeleteConnectionDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConnectionDefault creates DeleteConnectionDefault with default headers values
func NewDeleteConnectionDefault(code int) *DeleteConnectionDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteConnectionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete connection default response
func (o *DeleteConnectionDefault) WithStatusCode(code int) *DeleteConnectionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete connection default response
func (o *DeleteConnectionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete connection default response
func (o *DeleteConnectionDefault) WithPayload(payload *models.Error) *DeleteConnectionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete connection default response
func (o *DeleteConnectionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConnectionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
