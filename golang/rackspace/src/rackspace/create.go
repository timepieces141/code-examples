package rackspace

import (
	"errors"
	"log"
	"net"
	"os/exec"
	"epeters/cloudservice"
	"regexp"
)

// fatal/error regex
var fatalPattern = regexp.MustCompile(`(?is)FATAL.*`)

// field regexs
var instancePattern = regexp.MustCompile(`(?m)^Instance\sID:\s(.+)$`)
var hostIDPattern = regexp.MustCompile(`(?m)^Host\sID:\s(.+)$`)
var nodeNamePattern = regexp.MustCompile(`(?m)^Name:\s(.+)$`)
var flavorPattern = regexp.MustCompile(`(?m)^Flavor:\s(.+)$`)
var imagePattern = regexp.MustCompile(`(?m)^Image:\s(.+)$`)
var publicDNSPattern = regexp.MustCompile(`(?m)^Public\sDNS\sName:\s(.+)$`)
var publicIPPattern = regexp.MustCompile(`(?m)^Public\sIP\sAddress:\s(.+)$`)
var privateIPPattern = regexp.MustCompile(`(?m)^Private\sIP\sAddress:\s(.+)$`)
var passwordPattern = regexp.MustCompile(`(?m)^Password:\s(.+)$`)

func CreateServer(rackspace Rackspace, server RackspaceServer) (cloudservice.Server, error) {
	// assemble the arguments
	args := make([]string, 0)
	args = append(args, rackspace.name, "server", "create", "--image", server.image.imageID, "--flavor", server.flavor.flavorID, "--node-name", server.name, "--server-name", server.hostname)

	// execute the command
	cmd := exec.Command("knife", args...)
	output, err := cmd.CombinedOutput()

	// error check the command run
	if err != nil {
		log.Printf("[ERROR] Knife encountered an error, the output is:\n%s\n", string(output))
		return nil, err
	}

	// make the resulting RackspaceServer
	newServer, err2 := parseResults(string(output), server)
	if err2 != nil {
		return nil, err2
	}

	// populate a few lost fields
	newServer.hostname = server.hostname
	newServer.state = "active"

	return newServer, nil
}

func parseResults(results string, originalServer RackspaceServer) (RackspaceServer, error) {
	server := RackspaceServer{}

	// FATAL
	if fatalPattern.MatchString(results) {
		err := fatalPattern.FindString(results)
		return RackspaceServer{}, errors.New(err)
	}

	// Instance ID
	server.instanceID = instancePattern.FindStringSubmatch(results)[1]

	// Host ID
	server.hostID = hostIDPattern.FindStringSubmatch(results)[1]

	// Node Name
	server.name = nodeNamePattern.FindStringSubmatch(results)[1]

	// Flavor
	originalServer.flavor.name = flavorPattern.FindStringSubmatch(results)[1]
	server.flavor = originalServer.flavor

	// Image
	originalServer.image.name = imagePattern.FindStringSubmatch(results)[1]
	server.image = originalServer.image

	// PublicDNSName
	server.publicDNSName = publicDNSPattern.FindStringSubmatch(results)[1]

	// PublicIP
	pubAddr, _ := net.ResolveIPAddr("ip", publicIPPattern.FindStringSubmatch(results)[1])
	server.publicIP = *pubAddr

	// PrivateIP
	privAddr, _ := net.ResolveIPAddr("ip", privateIPPattern.FindStringSubmatch(results)[1])
	server.privateIP = *privAddr

	// Password
	server.password = passwordPattern.FindStringSubmatch(results)[1]

	return server, nil
}