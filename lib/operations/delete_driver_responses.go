package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

/*DeleteDriverNoContent Driver deleted

swagger:response deleteDriverNoContent
*/
type DeleteDriverNoContent struct {
}

// NewDeleteDriverNoContent creates DeleteDriverNoContent with default headers values
func NewDeleteDriverNoContent() *DeleteDriverNoContent {
	return &DeleteDriverNoContent{}
}

// WriteResponse to the client
func (o *DeleteDriverNoContent) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(204)
}

/*DeleteDriverNotFound Not Found

swagger:response deleteDriverNotFound
*/
type DeleteDriverNotFound struct {
}

// NewDeleteDriverNotFound creates DeleteDriverNotFound with default headers values
func NewDeleteDriverNotFound() *DeleteDriverNotFound {
	return &DeleteDriverNotFound{}
}

// WriteResponse to the client
func (o *DeleteDriverNotFound) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(404)
}

/*DeleteDriverInternalServerError Unexpected error

swagger:response deleteDriverInternalServerError
*/
type DeleteDriverInternalServerError struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewDeleteDriverInternalServerError creates DeleteDriverInternalServerError with default headers values
func NewDeleteDriverInternalServerError() *DeleteDriverInternalServerError {
	return &DeleteDriverInternalServerError{}
}

// WithPayload adds the payload to the delete driver internal server error response
func (o *DeleteDriverInternalServerError) WithPayload(payload string) *DeleteDriverInternalServerError {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *DeleteDriverInternalServerError) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
