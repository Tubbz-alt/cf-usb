package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/lib/genmodel"
)

/*UpdateDriverOK Driver updated

swagger:response updateDriverOK
*/
type UpdateDriverOK struct {

	// In: body
	Payload *genmodel.Driver `json:"body,omitempty"`
}

// NewUpdateDriverOK creates UpdateDriverOK with default headers values
func NewUpdateDriverOK() *UpdateDriverOK {
	return &UpdateDriverOK{}
}

// WithPayload adds the payload to the update driver o k response
func (o *UpdateDriverOK) WithPayload(payload *genmodel.Driver) *UpdateDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update driver o k response
func (o *UpdateDriverOK) SetPayload(payload *genmodel.Driver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateDriverNotFound Not Found

swagger:response updateDriverNotFound
*/
type UpdateDriverNotFound struct {
}

// NewUpdateDriverNotFound creates UpdateDriverNotFound with default headers values
func NewUpdateDriverNotFound() *UpdateDriverNotFound {
	return &UpdateDriverNotFound{}
}

// WriteResponse to the client
func (o *UpdateDriverNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*UpdateDriverConflict A driver with the same type already exists

swagger:response updateDriverConflict
*/
type UpdateDriverConflict struct {
}

// NewUpdateDriverConflict creates UpdateDriverConflict with default headers values
func NewUpdateDriverConflict() *UpdateDriverConflict {
	return &UpdateDriverConflict{}
}

// WriteResponse to the client
func (o *UpdateDriverConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
}

/*UpdateDriverInternalServerError Unexpected error

swagger:response updateDriverInternalServerError
*/
type UpdateDriverInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewUpdateDriverInternalServerError creates UpdateDriverInternalServerError with default headers values
func NewUpdateDriverInternalServerError() *UpdateDriverInternalServerError {
	return &UpdateDriverInternalServerError{}
}

// WithPayload adds the payload to the update driver internal server error response
func (o *UpdateDriverInternalServerError) WithPayload(payload string) *UpdateDriverInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update driver internal server error response
func (o *UpdateDriverInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateDriverInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
