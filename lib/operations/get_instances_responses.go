package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/lib/genmodel"
)

/*GetInstancesOK OK

swagger:response getInstancesOK
*/
type GetInstancesOK struct {

	// In: body
	Payload []*genmodel.Instance `json:"body,omitempty"`
}

// NewGetInstancesOK creates GetInstancesOK with default headers values
func NewGetInstancesOK() *GetInstancesOK {
	return &GetInstancesOK{}
}

// WithPayload adds the payload to the get instances o k response
func (o *GetInstancesOK) WithPayload(payload []*genmodel.Instance) *GetInstancesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instances o k response
func (o *GetInstancesOK) SetPayload(payload []*genmodel.Instance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstancesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetInstancesInternalServerError Unexpected error

swagger:response getInstancesInternalServerError
*/
type GetInstancesInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewGetInstancesInternalServerError creates GetInstancesInternalServerError with default headers values
func NewGetInstancesInternalServerError() *GetInstancesInternalServerError {
	return &GetInstancesInternalServerError{}
}

// WithPayload adds the payload to the get instances internal server error response
func (o *GetInstancesInternalServerError) WithPayload(payload string) *GetInstancesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instances internal server error response
func (o *GetInstancesInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstancesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
