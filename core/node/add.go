package node

import (
	"fmt"

	"warp/db"
)

const (
	nodeCollection = "nodes"
)

func init() {

}

type Node struct {

	Id string `gorethink:"id,omitempty"`
	Name string
	Hostname string
	IPAddress string
	SshPort string
	DestinationPort string
}

// New creates a new Node instance.
func New(name, hostname, ip, sshPort, destPort string) *Node {

	return &Node {
		Name: name,
		Hostname: hostname,
		IPAddress: ip,
		SshPort: sshPort,
		DestinationPort: destPort,
	}
}
