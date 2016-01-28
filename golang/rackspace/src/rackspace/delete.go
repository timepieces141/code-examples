package rackspace

import (
	"log"
	"os/exec"
	"strings"
)

func DeleteServer(rackspace Rackspace, server RackspaceServer) (bool, error) {
	cmd := exec.Command("knife", rackspace.name, "server", "delete", "--purge", server.instanceID)
	reader := strings.NewReader("Y")
	cmd.Stdin = reader
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error(), string(output))
		return false, err
	}

	return true, nil
}