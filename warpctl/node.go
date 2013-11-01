package main

import (
	"fmt"
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


func runNode(cmd *Command, args []string) {

	if addNode {
		fmt.Printf("Add new node %s.\n", nodeName)
	}

}

func setNodeFlags(cmd *Command) {

	cmd.Flag.BoolVar(&addNode, "add", false, "")
	cmd.Flag.StringVar(&nodeName, "name", "", "")

}
