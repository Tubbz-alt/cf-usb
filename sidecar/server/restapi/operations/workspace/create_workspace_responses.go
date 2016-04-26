package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/sidecar/server/models"
)

/*CreateWorkspaceCreated create workspace

swagger:response createWorkspaceCreated
*/
type CreateWorkspaceCreated struct {

	// In: body
	Payload *models.ServiceManagerWorkspaceResponse `json:"body,omitempty"`
}

// NewCreateWorkspaceCreated creates CreateWorkspaceCreated with default headers values
func NewCreateWorkspaceCreated() *CreateWorkspaceCreated {
	return &CreateWorkspaceCreated{}
}

// WithPayload adds the payload to the create workspace created response
func (o *CreateWorkspaceCreated) WithPayload(payload *models.ServiceManagerWorkspaceResponse) *CreateWorkspaceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create workspace created response
func (o *CreateWorkspaceCreated) SetPayload(payload *models.ServiceManagerWorkspaceResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkspaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateWorkspaceDefault generic error response

swagger:response createWorkspaceDefault
*/
type CreateWorkspaceDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateWorkspaceDefault creates CreateWorkspaceDefault with default headers values
func NewCreateWorkspaceDefault(code int) *CreateWorkspaceDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateWorkspaceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create workspace default response
func (o *CreateWorkspaceDefault) WithStatusCode(code int) *CreateWorkspaceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create workspace default response
func (o *CreateWorkspaceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create workspace default response
func (o *CreateWorkspaceDefault) WithPayload(payload *models.Error) *CreateWorkspaceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create workspace default response
func (o *CreateWorkspaceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkspaceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}