package main

import (
	"fmt"
	"os"
	"strconv"

	"warp/core/node"
	"warp/core/remote"
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
var delNode bool		// -del
var nodeName string		// -name
var hostname string		// -host
var ipAddress string		// -ip
var sshPort string		// -port

func setNodeFlags(cmd *Command) {

	cmd.Flag.BoolVar(&addNode, "add", false, "")
	cmd.Flag.BoolVar(&delNode, "del", false, "")

	cmd.Flag.StringVar(&nodeName, "name", "", "")
	cmd.Flag.StringVar(&hostname, "host", "", "")
	cmd.Flag.StringVar(&ipAddress, "ip", "", "")
	cmd.Flag.StringVar(&sshPort, "port", "", "")

}

func runNode(cmd *Command, args []string) {

	switch {

	case addNode:
		newNode()

	case delNode:
		deleteNode()
	}


}

func deleteNode() {

	if (nodeName == "") {
		fmt.Println("Please provide a valid Node name.")
		os.Exit(1)
	}

	fmt.Printf("Delete %q for good (y/N): ", nodeName)

	confirm := "n"
	_, err := fmt.Scan(&confirm)
	if err != nil {
		fmt.Printf("Unable to read input. Error: %s\n Nothing to delete.", err)
		os.Exit(1)
	}

	if confirm == "y" || confirm == "Y" {

		err = node.Delete(nodeName)
		if err != nil {
			fmt.Println(err)
		} else {

			fmt.Printf("Node %q was deleted successfully.\n", nodeName)
		}

		os.Exit(1)
	}

	fmt.Printf("%q was NOT deleted.\n", nodeName)
}

func newNode() {

	if (nodeName == "") {
		fmt.Println("Please provide a valid Node name.")
		os.Exit(1)
	}

	if (hostname == "" && ipAddress == "") {
		fmt.Println("Please provide a valid hostname or IP address.")
		os.Exit(1)
	}

	if (sshPort == "") {
		fmt.Printf("Please provide a valid SSH port %q is listening on.\n", nodeName)
		os.Exit(1)
	}else {

		i, err := strconv.Atoi(sshPort)
		if err != nil {
			fmt.Printf("%q is not a valid port integer.\n", sshPort)
			os.Exit(1)
		}

		if (i <= 0 || i > 65535) {
			fmt.Println("Port must be > 0 and < 65535.")
			os.Exit(1)
		}
	}


	n := node.New(
		nodeName,
		hostname,
		ipAddress,
		sshPort,
	)

	err := node.Save(n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d, err := remote.NewDestination(n.Id)
	if err != nil {

		fmt.Printf("Destination port was not added for node %q.\n", nodeName)
		node.Delete(nodeName)
		os.Exit(1)
	}

	err = remote.SaveDestination(d)
	if err != nil {
		fmt.Printf("Unable to save a destination port for node %q\n", nodeName)
		node.Delete(nodeName)
		os.Exit(1)
	}

	fmt.Printf("Node %q was saved and will be using destination port %q.\n", nodeName, d.LocalPort)
	fmt.Println("** Please add a firewall rules if necessary. **")
}
