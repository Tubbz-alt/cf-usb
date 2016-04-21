package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/lib/genmodel"
)

/*UpdateDriverInstanceOK Sucessfull response

swagger:response updateDriverInstanceOK
*/
type UpdateDriverInstanceOK struct {

	// In: body
	Payload *genmodel.DriverInstance `json:"body,omitempty"`
}

// NewUpdateDriverInstanceOK creates UpdateDriverInstanceOK with default headers values
func NewUpdateDriverInstanceOK() *UpdateDriverInstanceOK {
	return &UpdateDriverInstanceOK{}
}

// WithPayload adds the payload to the update driver instance o k response
func (o *UpdateDriverInstanceOK) WithPayload(payload *genmodel.DriverInstance) *UpdateDriverInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update driver instance o k response
func (o *UpdateDriverInstanceOK) SetPayload(payload *genmodel.DriverInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateDriverInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateDriverInstanceNotFound Not Found

swagger:response updateDriverInstanceNotFound
*/
type UpdateDriverInstanceNotFound struct {
}

// NewUpdateDriverInstanceNotFound creates UpdateDriverInstanceNotFound with default headers values
func NewUpdateDriverInstanceNotFound() *UpdateDriverInstanceNotFound {
	return &UpdateDriverInstanceNotFound{}
}

// WriteResponse to the client
func (o *UpdateDriverInstanceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*UpdateDriverInstanceConflict Conflict

swagger:response updateDriverInstanceConflict
*/
type UpdateDriverInstanceConflict struct {
}

// NewUpdateDriverInstanceConflict creates UpdateDriverInstanceConflict with default headers values
func NewUpdateDriverInstanceConflict() *UpdateDriverInstanceConflict {
	return &UpdateDriverInstanceConflict{}
}

// WriteResponse to the client
func (o *UpdateDriverInstanceConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
}

/*UpdateDriverInstanceInternalServerError Unexpected error

swagger:response updateDriverInstanceInternalServerError
*/
type UpdateDriverInstanceInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewUpdateDriverInstanceInternalServerError creates UpdateDriverInstanceInternalServerError with default headers values
func NewUpdateDriverInstanceInternalServerError() *UpdateDriverInstanceInternalServerError {
	return &UpdateDriverInstanceInternalServerError{}
}

// WithPayload adds the payload to the update driver instance internal server error response
func (o *UpdateDriverInstanceInternalServerError) WithPayload(payload string) *UpdateDriverInstanceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update driver instance internal server error response
func (o *UpdateDriverInstanceInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateDriverInstanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
