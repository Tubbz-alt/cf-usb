package lib

import (
	"encoding/json"
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"path/filepath"
	"runtime"

	"github.com/hpcloud/cf-usb/lib/config"
	"github.com/hpcloud/cf-usb/lib/model"
	"github.com/natefinch/pie"
)

type DriverProvider struct {
	driverType string
	client     *rpc.Client

	DriverProperties config.DriverProperties
}

func NewDriverProvider(driverType string, driverProperties config.DriverProperties) (*DriverProvider, error) {
	provider := DriverProvider{}

	driverPath := os.Getenv("USB_DRIVER_PATH")

	if driverPath == "" {
		driverPath = "drivers"
	}
	driverPath = filepath.Join(driverPath, driverType)
	if runtime.GOOS == "windows" {
		driverPath = driverPath + ".exe"
	}

	client, err := pie.StartProviderCodec(jsonrpc.NewClientCodec, os.Stdout, driverPath)
	if err != nil {
		return &provider, err
	}

	provider.client = client
	provider.DriverProperties = driverProperties
	provider.driverType = driverType

	_, err = provider.Init(driverProperties)

	return &provider, err
}

func (p *DriverProvider) Init(driverProperties config.DriverProperties) (string, error) {
	var result string
	err := p.client.Call(fmt.Sprintf("%s.Init", p.driverType), driverProperties, &result)
	return result, err
}

func (p *DriverProvider) Provision(provisonRequest model.DriverProvisionRequest) (string, error) {
	var result string
	err := p.client.Call(fmt.Sprintf("%s.Provision", p.driverType), provisonRequest, &result)
	return result, err
}

func (p *DriverProvider) Deprovision(deprovisionRequest model.DriverDeprovisionRequest) (string, error) {
	var result string
	err := p.client.Call(fmt.Sprintf("%s.Deprovision", p.driverType), deprovisionRequest, &result)
	return result, err
}

func (p *DriverProvider) Bind(bindRequest model.DriverBindRequest) (json.RawMessage, error) {
	var result json.RawMessage
	err := p.client.Call(fmt.Sprintf("%s.Bind", p.driverType), bindRequest, &result)
	return result, err
}

func (p *DriverProvider) Unbind(unbindRequest model.DriverUnbindRequest) (string, error) {
	var result string
	err := p.client.Call(fmt.Sprintf("%s.Unbind", p.driverType), unbindRequest, &result)
	return result, err
}
