package remote

import (
	"fmt"
	"strings"
	"strconv"
	"os/exec"

	"warp/utils"
)

const (
	sshPath = "/usr/bin/ssh"
	netstatPath = "/bin/netstat"
	netstatArgs = "-tuln"
)

var portRange = utils.NumberRange(10000, 65535)

var remoteNetstat = netstatPath + " " + netstatArgs
var remoteNodeFmt string = "%s@%s"

// remoteNetstatCmd returns output of `netstat -tln` from a remote machine
// over SSH. That is, all used ports of a remote machine.
func remoteNetstatCmd(username string, host string) (string, error) {

	remoteNode := fmt.Sprintf(remoteNodeFmt, username, host)

	cmd := exec.Command(
		sshPath,
		remoteNode,
		remoteNetstat)

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf(
			"unable to retrieve used ports from `%s` :  %s",
			remoteNode,
			err)
	}

	return string(out), nil
}

// parseNetstatPorts parses the output from `netstat` and returns a slice
// representing the ports.
func parseNetstatPorts(output string) []int {

	raw_ports := make([]string, 0)
	raw_lines := strings.Split(output, "\n")
	raw_lines = raw_lines[2:]  // Getting rid of useless output.

	for _, line := range raw_lines {

		if strings.Index(line, "tcp6") == 0 || strings.Index(line, "udp6") == 0 {
			continue
		}

		units := strings.Split(line, " ")
		for _, unit := range units {

			if strings.Index(unit, ":") > 0 {
			//if unit != "" && strings.Index(unit, ":") > 0 {

				items := strings.Split(unit, ":")
				if len(items) == 2 {
					raw_ports = append(raw_ports, items[1])
					break
				}
			}
		}

	}

	ports := make([]int, 0)
	for _, p := range raw_ports {

		converted, err := strconv.Atoi(p)
		if err != nil {
			continue
		}

		ports = append(ports, converted)
	}

	return ports
}

// UsablePort returns an available/usable port from the `host` machine.
func UsablePort(username, host string) (int, error) {

	output, err := remoteNetstatCmd(username, host)
	if err != nil {
		return 0, err
	}

	usedPorts := parseNetstatPorts(output)

	available := false
	for _, p := range portRange {

		for _, u := range usedPorts {

			if p == u {
				available = false
				break
			}

			available = true
		}

		if available {
			return p, nil
		}
	}

	return 0, fmt.Errorf("No port available to for use?!")
}

// IsPortAvailable checks if a specific port is available on the remote
// at the time it is called.
func IsPortAvailable(port int, username, host string) (bool, error) {

	output, err := remoteNetstatCmd(username, host)
	if err != nil {
		return false, err
	}

	usedPorts := parseNetstatPorts(output)

	for _, u := range usedPorts {

		if u == port {

			return false, nil
		}
	}

	return true, nil
}
