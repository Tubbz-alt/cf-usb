package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-usb/lib/brokermodel"
)

/*UpdateServiceInstanceOK New plan is effective. The expected response body is {}.

swagger:response updateServiceInstanceOK
*/
type UpdateServiceInstanceOK struct {

	// In: body
	Payload brokermodel.Empty `json:"body,omitempty"`
}

// NewUpdateServiceInstanceOK creates UpdateServiceInstanceOK with default headers values
func NewUpdateServiceInstanceOK() *UpdateServiceInstanceOK {
	return &UpdateServiceInstanceOK{}
}

// WithPayload adds the payload to the update service instance o k response
func (o *UpdateServiceInstanceOK) WithPayload(payload brokermodel.Empty) *UpdateServiceInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service instance o k response
func (o *UpdateServiceInstanceOK) SetPayload(payload brokermodel.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*UpdateServiceInstanceAccepted Service instance update is in progress. This triggers Cloud Controller to poll the Service Instance Last Operation Endpoint for operation status.

swagger:response updateServiceInstanceAccepted
*/
type UpdateServiceInstanceAccepted struct {

	// In: body
	Payload brokermodel.Empty `json:"body,omitempty"`
}

// NewUpdateServiceInstanceAccepted creates UpdateServiceInstanceAccepted with default headers values
func NewUpdateServiceInstanceAccepted() *UpdateServiceInstanceAccepted {
	return &UpdateServiceInstanceAccepted{}
}

// WithPayload adds the payload to the update service instance accepted response
func (o *UpdateServiceInstanceAccepted) WithPayload(payload brokermodel.Empty) *UpdateServiceInstanceAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service instance accepted response
func (o *UpdateServiceInstanceAccepted) SetPayload(payload brokermodel.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceInstanceAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*UpdateServiceInstanceUnprocessableEntity May be returned if the particular plan change requested is not supported or if the request can not currently be fulfilled due to the state of the instance (eg. instance utilization is over the quota of the requested plan). Broker should include a user-facing message in the body; for details [see Broker Errors](https://docs.cloudfoundry.org/services/api.html#broker-errors).

swagger:response updateServiceInstanceUnprocessableEntity
*/
type UpdateServiceInstanceUnprocessableEntity struct {

	// In: body
	Payload *brokermodel.AsyncError `json:"body,omitempty"`
}

// NewUpdateServiceInstanceUnprocessableEntity creates UpdateServiceInstanceUnprocessableEntity with default headers values
func NewUpdateServiceInstanceUnprocessableEntity() *UpdateServiceInstanceUnprocessableEntity {
	return &UpdateServiceInstanceUnprocessableEntity{}
}

// WithPayload adds the payload to the update service instance unprocessable entity response
func (o *UpdateServiceInstanceUnprocessableEntity) WithPayload(payload *brokermodel.AsyncError) *UpdateServiceInstanceUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service instance unprocessable entity response
func (o *UpdateServiceInstanceUnprocessableEntity) SetPayload(payload *brokermodel.AsyncError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceInstanceUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateServiceInstanceDefault generic error response

swagger:response updateServiceInstanceDefault
*/
type UpdateServiceInstanceDefault struct {
	_statusCode int

	// In: body
	Payload *brokermodel.BrokerError `json:"body,omitempty"`
}

// NewUpdateServiceInstanceDefault creates UpdateServiceInstanceDefault with default headers values
func NewUpdateServiceInstanceDefault(code int) *UpdateServiceInstanceDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateServiceInstanceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update service instance default response
func (o *UpdateServiceInstanceDefault) WithStatusCode(code int) *UpdateServiceInstanceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update service instance default response
func (o *UpdateServiceInstanceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update service instance default response
func (o *UpdateServiceInstanceDefault) WithPayload(payload *brokermodel.BrokerError) *UpdateServiceInstanceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service instance default response
func (o *UpdateServiceInstanceDefault) SetPayload(payload *brokermodel.BrokerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceInstanceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}