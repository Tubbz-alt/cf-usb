package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/hpcloud/cf-usb/lib/genmodel"
)

/*GetDialSchemaOK OK

swagger:response getDialSchemaOK
*/
type GetDialSchemaOK struct {

	// In: body
	Payload genmodel.DialSchema `json:"body,omitempty"`
}

// NewGetDialSchemaOK creates GetDialSchemaOK with default headers values
func NewGetDialSchemaOK() *GetDialSchemaOK {
	return &GetDialSchemaOK{}
}

// WithPayload adds the payload to the get dial schema o k response
func (o *GetDialSchemaOK) WithPayload(payload genmodel.DialSchema) *GetDialSchemaOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *GetDialSchemaOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetDialSchemaNotFound Not Found

swagger:response getDialSchemaNotFound
*/
type GetDialSchemaNotFound struct {
}

// NewGetDialSchemaNotFound creates GetDialSchemaNotFound with default headers values
func NewGetDialSchemaNotFound() *GetDialSchemaNotFound {
	return &GetDialSchemaNotFound{}
}

// WriteResponse to the client
func (o *GetDialSchemaNotFound) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(404)
}

/*GetDialSchemaInternalServerError Unexpected error

swagger:response getDialSchemaInternalServerError
*/
type GetDialSchemaInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewGetDialSchemaInternalServerError creates GetDialSchemaInternalServerError with default headers values
func NewGetDialSchemaInternalServerError() *GetDialSchemaInternalServerError {
	return &GetDialSchemaInternalServerError{}
}

// WithPayload adds the payload to the get dial schema internal server error response
func (o *GetDialSchemaInternalServerError) WithPayload(payload string) *GetDialSchemaInternalServerError {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *GetDialSchemaInternalServerError) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
