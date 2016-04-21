package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/lib/genmodel"
)

/*GetDialOK Sucessfull response

swagger:response getDialOK
*/
type GetDialOK struct {

	// In: body
	Payload *genmodel.Dial `json:"body,omitempty"`
}

// NewGetDialOK creates GetDialOK with default headers values
func NewGetDialOK() *GetDialOK {
	return &GetDialOK{}
}

// WithPayload adds the payload to the get dial o k response
func (o *GetDialOK) WithPayload(payload *genmodel.Dial) *GetDialOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dial o k response
func (o *GetDialOK) SetPayload(payload *genmodel.Dial) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDialNotFound Not Found

swagger:response getDialNotFound
*/
type GetDialNotFound struct {
}

// NewGetDialNotFound creates GetDialNotFound with default headers values
func NewGetDialNotFound() *GetDialNotFound {
	return &GetDialNotFound{}
}

// WriteResponse to the client
func (o *GetDialNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*GetDialInternalServerError Unexpected error

swagger:response getDialInternalServerError
*/
type GetDialInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewGetDialInternalServerError creates GetDialInternalServerError with default headers values
func NewGetDialInternalServerError() *GetDialInternalServerError {
	return &GetDialInternalServerError{}
}

// WithPayload adds the payload to the get dial internal server error response
func (o *GetDialInternalServerError) WithPayload(payload string) *GetDialInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dial internal server error response
func (o *GetDialInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
