package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/lib/genmodel"
)

/*GetDriverEndpointsOK OK

swagger:response getDriverEndpointsOK
*/
type GetDriverEndpointsOK struct {

	// In: body
	Payload []*genmodel.DriverEndpoint `json:"body,omitempty"`
}

// NewGetDriverEndpointsOK creates GetDriverEndpointsOK with default headers values
func NewGetDriverEndpointsOK() *GetDriverEndpointsOK {
	return &GetDriverEndpointsOK{}
}

// WithPayload adds the payload to the get driver endpoints o k response
func (o *GetDriverEndpointsOK) WithPayload(payload []*genmodel.DriverEndpoint) *GetDriverEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver endpoints o k response
func (o *GetDriverEndpointsOK) SetPayload(payload []*genmodel.DriverEndpoint) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetDriverEndpointsInternalServerError Unexpected error

swagger:response getDriverEndpointsInternalServerError
*/
type GetDriverEndpointsInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewGetDriverEndpointsInternalServerError creates GetDriverEndpointsInternalServerError with default headers values
func NewGetDriverEndpointsInternalServerError() *GetDriverEndpointsInternalServerError {
	return &GetDriverEndpointsInternalServerError{}
}

// WithPayload adds the payload to the get driver endpoints internal server error response
func (o *GetDriverEndpointsInternalServerError) WithPayload(payload string) *GetDriverEndpointsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver endpoints internal server error response
func (o *GetDriverEndpointsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverEndpointsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
