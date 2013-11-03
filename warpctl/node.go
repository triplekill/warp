package main

import (
	"fmt"

	"warp/core/node"
)


var nodeCommand = &Command{
	UsageLine:	"node [flags]",
	Short:		"manage nodes",
	Long:`
Node allows the ability to manage nodes.
`,
}

func init() {

	nodeCommand.Run = runNode

	setNodeFlags(nodeCommand)
}

// Flags used by node.
var addNode bool		// -add
var nodeName string		// -name
var hostname string		// -host
var ipAddress string		// -ip
var sshPort string		// -sshport

func setNodeFlags(cmd *Command) {

	cmd.Flag.BoolVar(&addNode, "add", false, "")
	cmd.Flag.StringVar(&nodeName, "name", "", "")
	cmd.Flag.StringVar(&hostname, "host", "", "")
	cmd.Flag.StringVar(&ipAddress, "ip", "", "")
	cmd.Flag.StringVar(&sshPort, "sshport", "", "")

}

func runNode(cmd *Command, args []string) {

	if addNode {

		n := node.New(
			nodeName,
			"hostname",
			"ip address",
			"22",
		)
		err := node.Save(n)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Add new node %s.\n", nodeName)
		}
	}

}
