package rackspace

import (
	"net"
	"epeters/cloudservice"
)

type RackspaceServer struct {
	instanceID string
	hostID string
	name string
	hostname string
	publicDNSName string
	publicIP net.IPAddr
	privateIP net.IPAddr
	flavor RackspaceFlavor
	image RackspaceImage
	state string
	password string
}

func NewRackspaceServer(name string, hostname string, flavor RackspaceFlavor, image RackspaceImage) RackspaceServer {
	server := RackspaceServer{}
	server.name = name
	server.hostname = hostname
	server.flavor = flavor
	server.image = image
	return server
}

func (rss RackspaceServer) InstanceID() string {
	return rss.instanceID
}

func (rss RackspaceServer) Name() string {
	return rss.name
}

func (rss RackspaceServer) PublicIP() net.IPAddr {
	return rss.publicIP
}

func (rss RackspaceServer) PrivateIP() net.IPAddr {
	return rss.privateIP
}

func (rss RackspaceServer) Flavor() cloudservice.Flavor {
	return &rss.flavor
}

func (rss RackspaceServer) Image() cloudservice.Image {
	return &rss.image
}

func (rss RackspaceServer) State() string {
	return rss.state
}