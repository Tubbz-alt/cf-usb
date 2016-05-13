package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUsbMgmtAPI creates a new UsbMgmt instance
func NewUsbMgmtAPI(spec *loads.Document) *UsbMgmtAPI {
	o := &UsbMgmtAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*UsbMgmtAPI Universal Service Broker Management API */
type UsbMgmtAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// AuthorizationAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	AuthorizationAuth func(string) (interface{}, error)

	// GetDriverEndpointHandler sets the operation handler for the get driver endpoint operation
	GetDriverEndpointHandler GetDriverEndpointHandler
	// GetDriverEndpointsHandler sets the operation handler for the get driver endpoints operation
	GetDriverEndpointsHandler GetDriverEndpointsHandler
	// GetInfoHandler sets the operation handler for the get info operation
	GetInfoHandler GetInfoHandler
	// PingDriverEndpointHandler sets the operation handler for the ping driver endpoint operation
	PingDriverEndpointHandler PingDriverEndpointHandler
	// RegisterDriverEndpointHandler sets the operation handler for the register driver endpoint operation
	RegisterDriverEndpointHandler RegisterDriverEndpointHandler
	// UnregisterDriverInstanceHandler sets the operation handler for the unregister driver instance operation
	UnregisterDriverInstanceHandler UnregisterDriverInstanceHandler
	// UpdateCatalogHandler sets the operation handler for the update catalog operation
	UpdateCatalogHandler UpdateCatalogHandler
	// UpdateDriverEndpointHandler sets the operation handler for the update driver endpoint operation
	UpdateDriverEndpointHandler UpdateDriverEndpointHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup
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

	if o.AuthorizationAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.GetDriverEndpointHandler == nil {
		unregistered = append(unregistered, "GetDriverEndpointHandler")
	}

	if o.GetDriverEndpointsHandler == nil {
		unregistered = append(unregistered, "GetDriverEndpointsHandler")
	}

	if o.GetInfoHandler == nil {
		unregistered = append(unregistered, "GetInfoHandler")
	}

	if o.PingDriverEndpointHandler == nil {
		unregistered = append(unregistered, "PingDriverEndpointHandler")
	}

	if o.RegisterDriverEndpointHandler == nil {
		unregistered = append(unregistered, "RegisterDriverEndpointHandler")
	}

	if o.UnregisterDriverInstanceHandler == nil {
		unregistered = append(unregistered, "UnregisterDriverInstanceHandler")
	}

	if o.UpdateCatalogHandler == nil {
		unregistered = append(unregistered, "UpdateCatalogHandler")
	}

	if o.UpdateDriverEndpointHandler == nil {
		unregistered = append(unregistered, "UpdateDriverEndpointHandler")
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
func (o *UsbMgmtAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "Authorization":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, func(tok string) (interface{}, error) { return o.AuthorizationAuth(tok) })

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *UsbMgmtAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *UsbMgmtAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *UsbMgmtAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *UsbMgmtAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/driver_endpoints/{driver_endpoint_id}"] = NewGetDriverEndpoint(o.context, o.GetDriverEndpointHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/driver_endpoints"] = NewGetDriverEndpoints(o.context, o.GetDriverEndpointsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/info"] = NewGetInfo(o.context, o.GetInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/driver_endpoint/{driver_endpoint_id}/ping"] = NewPingDriverEndpoint(o.context, o.PingDriverEndpointHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/driver_endpoints"] = NewRegisterDriverEndpoint(o.context, o.RegisterDriverEndpointHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/driver_endpoints/{driver_endpoint_id}"] = NewUnregisterDriverInstance(o.context, o.UnregisterDriverInstanceHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/update_catalog"] = NewUpdateCatalog(o.context, o.UpdateCatalogHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/driver_endpoints/{driver_endpoint_id}"] = NewUpdateDriverEndpoint(o.context, o.UpdateDriverEndpointHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *UsbMgmtAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}