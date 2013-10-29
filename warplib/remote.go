package main

import (
	"fmt"
	"strings"
	"os/exec"
)

const (
	sshPath = "/usr/bin/ssh"
	netstatPath = "/bin/netstat"
	netstatArgs = "-tuln"
)

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

// parseNetstatPorts parses the output from `netstat` and returns a slice of
// strings representing the ports.
func parseNetstatPorts(output string) []string {

	ports := make([]string, 0)
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
					ports = append(ports, items[1])
					break
				}
			}
		}

	}

	return ports
}
