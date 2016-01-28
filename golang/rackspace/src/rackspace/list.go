package rackspace

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"epeters/cloudservice"
	"regexp"
	"strconv"
	"strings"
)

func list(command string) (string, error) {
	cmd := exec.Command("knife", "rackspace", command, "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(output))
	}

	return string(output), nil
}

func search(command string, strRegexp string) (string, error) {
	output, err := list(command)
	if err != nil {
		return "", err
	}

	// search the output: start of line or at least 2 spaces, the search
	// term, and then either at least 2 more spaces or the end of the line
	// (but capture the whole line)
	re := fmt.Sprintf("(?m)(^|^.*[ ]{2,}?)(%s)($|[ ]{2,}?.*$)", strRegexp)
	pattern := regexp.MustCompile(re)
	if pattern.MatchString(output) {
		found := pattern.FindString(output)
		return found, nil
	}

	return "", nil
}

func ServerList() ([]cloudservice.Server, error) {
	output, err := list("server")
	if err != nil {
		return make([]cloudservice.Server, 0), err
	}
	return parseServerList(output, true), nil
}

func parseServerList(output string, skipHeader bool) []cloudservice.Server {
	servers := make([]cloudservice.Server, 0)
	lines := strings.Split(output, "\n")

	// loop through the lines parsing out a RackspaceServer, ommiting the header
	// if required
	for index, line := range lines {
		if skipHeader && index == 0 {
			continue
		}

		var server = RackspaceServer{}
		fields := strings.Fields(line)
		if len(fields) == 7 {
			server.instanceID = fields[0]
			server.name = fields[1]
			pubAddr, _ := net.ResolveIPAddr("ip", fields[2])
			server.publicIP = *pubAddr
			privAddr, _ := net.ResolveIPAddr("ip", fields[3])
			server.privateIP = *privAddr
			server.flavor = NewRackspaceFlavor(fields[4], "", 0, 0, 0)
			server.image = NewRackspaceImage(fields[5], "")
			server.state = fields[6]
			servers = append(servers, server)
			continue
		}
		// hack for missing public IP
		if len(fields) == 6 {
			server.instanceID = fields[0]
			server.name = fields[1]
			privAddr, _ := net.ResolveIPAddr("ip", fields[2])
			server.privateIP = *privAddr
			server.flavor = NewRackspaceFlavor(fields[3], "", 0, 0, 0)
			server.image = NewRackspaceImage(fields[4], "")
			server.state = fields[5]
			servers = append(servers, server)
			continue
		}
	}

	return servers
}

func SearchForServer(strRegexp string) (cloudservice.Server, error) {
	lines, err := search("server", strRegexp)
	if err != nil {
		return new(RackspaceServer), err
	}

	if strings.EqualFold(lines, "") {
		return new(RackspaceServer), nil
	}

	return parseServerList(lines, false)[0], nil
}

func CheckForServer(strRegexp string) (bool, error) {
	lines, err := search("server", strRegexp)
	if err != nil {
		return false, err
	}

	return !strings.EqualFold(lines, ""), nil
}

func ImageList() ([]cloudservice.Image, error) {
	output, err := list("image")
	if err != nil {
		return make([]cloudservice.Image, 0), err
	}
	return parseImageList(output, true), nil
}

func parseImageList(output string, skipHeader bool) []cloudservice.Image {
	images := make([]cloudservice.Image, 0)
	lines := strings.Split(output, "\n")

	// loop through the lines parsing out a RackspaceImage, ommiting the header
	// if required
	for index, line := range lines {
		if skipHeader && index == 0 {
			continue
		}

		var image = RackspaceImage{}
		fields := strings.Fields(line)
		if len(fields) > 1 {
			image.imageID = fields[0]
			image.name = strings.Join(fields[1:], " ")
			images = append(images, image)
		}
	}

	return images
}

func SearchForImage(strRegexp string) (cloudservice.Image, error) {
	lines, err := search("image", strRegexp)
	if err != nil {
		return new(RackspaceImage), err
	}

	if strings.EqualFold(lines, "") {
		return new(RackspaceImage), nil
	}

	return parseImageList(lines, false)[0], nil
}

func CheckForImage(strRegexp string) (bool, error) {
	lines, err := search("image", strRegexp)
	if err != nil {
		return false, err
	}

	return !strings.EqualFold(lines, ""), nil
}

func FlavorList() ([]cloudservice.Flavor, error) {
	output, err := list("flavor")
	if err != nil {
		return make([]cloudservice.Flavor, 0), err
	}
	return parseFlavorList(output, true), nil
}

func parseFlavorList(output string, skipHeader bool) []cloudservice.Flavor {
	flavors := make([]cloudservice.Flavor, 0)
	lines := strings.Split(output, "\n")

	// loop through the lines parsing out a RackspaceFlavor, ommiting the header
	// if required
	for index, line := range lines {
		if skipHeader && index == 0 {
			continue
		}

		var flavor = RackspaceFlavor{}
		pattern := regexp.MustCompile("\\s{2,}")
		fields := pattern.Split(line, -1)
		diskPattern := regexp.MustCompile(" GB[ ]*")
		if len(fields) > 1 {
			flavor.flavorID = fields[0]
			flavor.name = fields[1]
			flavor.cpus, _ = strconv.Atoi(fields[2])
			flavor.ram, _ = strconv.Atoi(fields[3])
			flavor.disk, _ = strconv.Atoi(diskPattern.ReplaceAllString(fields[4], ""))
			flavors = append(flavors, flavor)
		}
	}

	return flavors
}

func SearchForFlavor(strRegexp string) (cloudservice.Flavor, error) {
	lines, err := search("flavor", strRegexp)
	if err != nil {
		return new(RackspaceFlavor), err
	}

	if strings.EqualFold(lines, "") {
		return new(RackspaceFlavor), nil
	}

	return parseFlavorList(lines, false)[0], nil
}

func CheckForFlavor(strRegexp string) (bool, error) {
	lines, err := search("flavor", strRegexp)
	if err != nil {
		return false, err
	}

	return !strings.EqualFold(lines, ""), nil
}