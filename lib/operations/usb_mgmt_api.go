package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewUsbMgmtAPI creates a new UsbMgmt instance
func NewUsbMgmtAPI(spec *spec.Document) *UsbMgmtAPI {
	o := &UsbMgmtAPI{
		spec:            spec,
		handlers:        make(map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
	}

	return o
}

/* UsbMgmtAPI  */
type UsbMgmtAPI struct {
	spec            *spec.Document
	context         *middleware.Context
	handlers        map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer httpkit.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer httpkit.Producer

	// GetServicePlanByIDHandler sets the operation handler for the get service plan by ID operation
	GetServicePlanByIDHandler GetServicePlanByIDHandler

	// GetInfoHandler sets the operation handler for the get info operation
	GetInfoHandler GetInfoHandler

	// GetDriverConfigsHandler sets the operation handler for the get driver configs operation
	GetDriverConfigsHandler GetDriverConfigsHandler

	// GetServiceByIDHandler sets the operation handler for the get service by ID operation
	GetServiceByIDHandler GetServiceByIDHandler

	// CreateDriverConfigHandler sets the operation handler for the create driver config operation
	CreateDriverConfigHandler CreateDriverConfigHandler

	// DeleteDriverConfigHandler sets the operation handler for the delete driver config operation
	DeleteDriverConfigHandler DeleteDriverConfigHandler

	// DeleteServicePlanHandler sets the operation handler for the delete service plan operation
	DeleteServicePlanHandler DeleteServicePlanHandler

	// GetServicePlansHandler sets the operation handler for the get service plans operation
	GetServicePlansHandler GetServicePlansHandler

	// GetDriverConfigByIDHandler sets the operation handler for the get driver config by ID operation
	GetDriverConfigByIDHandler GetDriverConfigByIDHandler

	// UpdateDriverConfigHandler sets the operation handler for the update driver config operation
	UpdateDriverConfigHandler UpdateDriverConfigHandler

	// DeleteServiceHandler sets the operation handler for the delete service operation
	DeleteServiceHandler DeleteServiceHandler

	// GetDriversHandler sets the operation handler for the get drivers operation
	GetDriversHandler GetDriversHandler

	// GetServicesHandler sets the operation handler for the get services operation
	GetServicesHandler GetServicesHandler

	// CreateServiceHandler sets the operation handler for the create service operation
	CreateServiceHandler CreateServiceHandler

	// CreateServicePlanHandler sets the operation handler for the create service plan operation
	CreateServicePlanHandler CreateServicePlanHandler

	// UpdateServiceHandler sets the operation handler for the update service operation
	UpdateServiceHandler UpdateServiceHandler

	// UpdateServicePlanHandler sets the operation handler for the update service plan operation
	UpdateServicePlanHandler UpdateServicePlanHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)
}

// SetDefaultProduces sets the default produces media type
func (o *UsbMgmtAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *UsbMgmtAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *UsbMgmtAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *UsbMgmtAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *UsbMgmtAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *UsbMgmtAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the UsbMgmtAPI
func (o *UsbMgmtAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetServicePlanByIDHandler == nil {
		unregistered = append(unregistered, "GetServicePlanByIDHandler")
	}

	if o.GetInfoHandler == nil {
		unregistered = append(unregistered, "GetInfoHandler")
	}

	if o.GetDriverConfigsHandler == nil {
		unregistered = append(unregistered, "GetDriverConfigsHandler")
	}

	if o.GetServiceByIDHandler == nil {
		unregistered = append(unregistered, "GetServiceByIDHandler")
	}

	if o.CreateDriverConfigHandler == nil {
		unregistered = append(unregistered, "CreateDriverConfigHandler")
	}

	if o.DeleteDriverConfigHandler == nil {
		unregistered = append(unregistered, "DeleteDriverConfigHandler")
	}

	if o.DeleteServicePlanHandler == nil {
		unregistered = append(unregistered, "DeleteServicePlanHandler")
	}

	if o.GetServicePlansHandler == nil {
		unregistered = append(unregistered, "GetServicePlansHandler")
	}

	if o.GetDriverConfigByIDHandler == nil {
		unregistered = append(unregistered, "GetDriverConfigByIDHandler")
	}

	if o.UpdateDriverConfigHandler == nil {
		unregistered = append(unregistered, "UpdateDriverConfigHandler")
	}

	if o.DeleteServiceHandler == nil {
		unregistered = append(unregistered, "DeleteServiceHandler")
	}

	if o.GetDriversHandler == nil {
		unregistered = append(unregistered, "GetDriversHandler")
	}

	if o.GetServicesHandler == nil {
		unregistered = append(unregistered, "GetServicesHandler")
	}

	if o.CreateServiceHandler == nil {
		unregistered = append(unregistered, "CreateServiceHandler")
	}

	if o.CreateServicePlanHandler == nil {
		unregistered = append(unregistered, "CreateServicePlanHandler")
	}

	if o.UpdateServiceHandler == nil {
		unregistered = append(unregistered, "UpdateServiceHandler")
	}

	if o.UpdateServicePlanHandler == nil {
		unregistered = append(unregistered, "UpdateServicePlanHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *UsbMgmtAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *UsbMgmtAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *UsbMgmtAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *UsbMgmtAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation id
func (o *UsbMgmtAPI) HandlerFor(operationID string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	h, ok := o.handlers[operationID]
	return h, ok
}

func (o *UsbMgmtAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	o.handlers = make(map[string]http.Handler)

	o.handlers["getServicePlanByID"] = NewGetServicePlanByID(o.context, o.GetServicePlanByIDHandler)

	o.handlers["getInfo"] = NewGetInfo(o.context, o.GetInfoHandler)

	o.handlers["getDriverConfigs"] = NewGetDriverConfigs(o.context, o.GetDriverConfigsHandler)

	o.handlers["getServiceByID"] = NewGetServiceByID(o.context, o.GetServiceByIDHandler)

	o.handlers["createDriverConfig"] = NewCreateDriverConfig(o.context, o.CreateDriverConfigHandler)

	o.handlers["deleteDriverConfig"] = NewDeleteDriverConfig(o.context, o.DeleteDriverConfigHandler)

	o.handlers["deleteServicePlan"] = NewDeleteServicePlan(o.context, o.DeleteServicePlanHandler)

	o.handlers["getServicePlans"] = NewGetServicePlans(o.context, o.GetServicePlansHandler)

	o.handlers["getDriverConfigByID"] = NewGetDriverConfigByID(o.context, o.GetDriverConfigByIDHandler)

	o.handlers["updateDriverConfig"] = NewUpdateDriverConfig(o.context, o.UpdateDriverConfigHandler)

	o.handlers["deleteService"] = NewDeleteService(o.context, o.DeleteServiceHandler)

	o.handlers["getDrivers"] = NewGetDrivers(o.context, o.GetDriversHandler)

	o.handlers["getServices"] = NewGetServices(o.context, o.GetServicesHandler)

	o.handlers["createService"] = NewCreateService(o.context, o.CreateServiceHandler)

	o.handlers["createServicePlan"] = NewCreateServicePlan(o.context, o.CreateServicePlanHandler)

	o.handlers["updateService"] = NewUpdateService(o.context, o.UpdateServiceHandler)

	o.handlers["updateServicePlan"] = NewUpdateServicePlan(o.context, o.UpdateServicePlanHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve())
func (o *UsbMgmtAPI) Serve() http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler()
}