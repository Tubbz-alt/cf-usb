package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

/*DeleteDriverInstanceNoContent Driver deleted

swagger:response deleteDriverInstanceNoContent
*/
type DeleteDriverInstanceNoContent struct {
}

// NewDeleteDriverInstanceNoContent creates DeleteDriverInstanceNoContent with default headers values
func NewDeleteDriverInstanceNoContent() *DeleteDriverInstanceNoContent {
	return &DeleteDriverInstanceNoContent{}
}

// WriteResponse to the client
func (o *DeleteDriverInstanceNoContent) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(204)
}

/*DeleteDriverInstanceNotFound Not Found

swagger:response deleteDriverInstanceNotFound
*/
type DeleteDriverInstanceNotFound struct {
}

// NewDeleteDriverInstanceNotFound creates DeleteDriverInstanceNotFound with default headers values
func NewDeleteDriverInstanceNotFound() *DeleteDriverInstanceNotFound {
	return &DeleteDriverInstanceNotFound{}
}

// WriteResponse to the client
func (o *DeleteDriverInstanceNotFound) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(404)
}

/*DeleteDriverInstanceInternalServerError Unexpected error

swagger:response deleteDriverInstanceInternalServerError
*/
type DeleteDriverInstanceInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewDeleteDriverInstanceInternalServerError creates DeleteDriverInstanceInternalServerError with default headers values
func NewDeleteDriverInstanceInternalServerError() *DeleteDriverInstanceInternalServerError {
	return &DeleteDriverInstanceInternalServerError{}
}

// WithPayload adds the payload to the delete driver instance internal server error response
func (o *DeleteDriverInstanceInternalServerError) WithPayload(payload string) *DeleteDriverInstanceInternalServerError {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *DeleteDriverInstanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
