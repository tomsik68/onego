package onego

import (
	"github.com/onego-project/onego/resources"
	"github.com/onego-project/onego/services"
	"github.com/onego-project/xmlrpc"
)

// Client structure contains XML-RPC client and virtual machine
type Client struct {
	VirtualMachineService services.VirtualMachineService
}

// CreateClient method
func CreateClient(endpoint, key string) *Client {
	return &Client{VirtualMachineService: services.VirtualMachineService{RPC: &resources.RPC{Client: xmlrpc.NewClient(endpoint), Key: key}}}
}
