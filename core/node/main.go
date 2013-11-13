package node

import (

	"fmt"

	"warp/db"
)

const (
	nodeTable = "nodes"
	nodeNameIndex = "name"
)

var allNodeTables = []string {
	nodeTable,
}

var allNodeIndices = []string {
	nodeNameIndex,
}

type Node struct {

	Id string                       `gorethink:"id,omitempty"`
	Name string                     `gorethink:"name"`
	Hostname string                 `gorethink:"hostname"`
	IPAddress string                `gorethink:"ip_address"`
	SshPort string                  `gorethink:"ssh_port"`
}

func init() {

	db.Readify()

	db.CreateTables(allNodeTables)
	db.CreateIndices(nodeTable, allNodeIndices)
}


// New creates a new Node instance.
func New(name, hostname, ip, sshPort string) *Node {

	return &Node {
		Name: name,
		Hostname: hostname,
		IPAddress: ip,
		SshPort: sshPort,
	}
}

func Save(node *Node) error {

	exist, err := DoesNodeExist(node.Name)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("Node with name %q already exist.", node.Name)
	}

	keys, err := db.InsertRow(nodeTable, *node)
	if err != nil {
		return err
	}

	if len(keys) >= 1 {
		node.Id = keys[0]
		return nil
	}

	return fmt.Errorf("No ID was generated.")
}

func Delete(name string) error {

	err := db.DeleteByIndex(nodeTable, nodeNameIndex, name)
	if err != nil {
		return fmt.Errorf("Unable to delete node %q. %s", name, err)
	}

	return nil
}

func DoesNodeExist(name string) (bool, error) {

	items, err := db.GetByIndex(
		nodeTable,
		nodeNameIndex,
		name,
	)
	if err != nil {
		return false, err
	}

	if len(items) > 0 {
		return true, nil
	}

	return false, nil
}
