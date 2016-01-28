package openstack

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

var blankIP, _ = net.ResolveIPAddr("ip", "0.0.0.0")

type OpenStack struct {
	name string
	FloatingIPBlock []net.IPAddr
	SSHKey string
	availableIps []net.IPAddr
	usedIPs []net.IPAddr
	mutex sync.Mutex
}

func NewOpenStack(floatingIPBlock []net.IPAddr, sshKey string) *OpenStack {
	block := append(floatingIPBlock)
	available := append(floatingIPBlock)
	return &OpenStack{"openstack", block, sshKey, available, make([]net.IPAddr, 0), sync.Mutex{}}
}

func (cloud *OpenStack) ServiceName() string {
	return cloud.name
}

func (cloud *OpenStack) AllocateFloatingIP() (net.IPAddr, error) {
	if cloud.availableIps != nil {
		// mutex
		cloud.mutex.Lock()

		// grab the first available address
		addr := cloud.availableIps[0]

		// slice it out of the available
		if len(cloud.availableIps) > 1 {
			cloud.availableIps = cloud.availableIps[1:]
		} else {
			cloud.availableIps = nil
		}

		// put the address in used
		cloud.usedIPs = append(cloud.usedIPs, addr)

		// mutex
		cloud.mutex.Unlock()

		return addr, nil
	}

	// no addresses left
	return *blankIP, errors.New("The floating IP Address block has been exhausted")
}

func (cloud *OpenStack) ReturnFloatingIP(addr net.IPAddr) error {
	// grab the address out of the used
	index := indexOfAddress(cloud.usedIPs, addr)
	if index == -1 {
		return errors.New(fmt.Sprintf("'%s' address was not previously allocated", addr.String()))
	}

	// mutex
	cloud.mutex.Lock()

	// remove from usedIPs
	cloud.usedIPs = cloud.usedIPs[:index+copy(cloud.usedIPs[index:], cloud.usedIPs[index+1:])]

	// recreate the availableIPs slice if it was set to nil
	if cloud.availableIps == nil {
		cloud.availableIps = make([]net.IPAddr, 0)
	}

	// add the address to available
	cloud.availableIps = append(cloud.availableIps, addr)

	// mutex
	cloud.mutex.Unlock()

	return nil
}

func indexOfAddress(slice []net.IPAddr, search net.IPAddr) int {
	for index, addr := range slice {
		if compareIPAddr(addr, search) {
			return index
		}
	}
	return -1
}

func compareIPAddr(first net.IPAddr, second net.IPAddr) bool {
	firstBytes := first.IP
	secondBytes := second.IP
	for index, b := range firstBytes {
		if secondBytes[index] != b {
			return false
		}
	}
	return true
}