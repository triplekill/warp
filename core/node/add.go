package node

import (

	"warp/db"
	"warp/utils"
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

	Id string			`gorethink:"id,omitempty"`
	Name string			`gorethink:"name"`
	Hostname string			`gorethink:"hostname"`
	IPAddress string		`gorethink:"ip_address"`
	SshPort string			`gorethink:"ssh_port"`
	DestinationPort string		`gorethink:"destination_port"`
}

func init() {

	db.Readify()

	createTables()
	createIndices()
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

func Save(node *Node) error {

	_, err := db.InsertRow(nodeTable, *node)
	if err != nil {
		return err
	}

	return nil
}

func createTables() {

	for _, table := range allNodeTables {

		exist, err := db.DoesTableExist(table);
		if err != nil {
			utils.Panicf(
				"Could not verify table %q. Error: %s",
				table,
				err,
			)
		}

		if !exist {
			_, err := db.CreateTable(table)
			if err != nil {
				utils.Panicf(
					"Unable to create %q table. Error: %s",
					table,
					err,
				)
			}
		}
	}
}

func createIndices() {

	for _, index := range allNodeIndices {

		indexExist, err := db.DoesIndexExist(index, nodeTable)
		if err != nil {
			utils.Panicf(
				"Could not verify index %q for table %q. Error: %s",
				index,
				nodeTable,
				err,
			)
		}

		if !indexExist {
			_, err := db.CreateIndex(index, nodeTable)
			if err != nil {
				utils.Panicf(
					"Unable to create index %q for table %q. Error: %s",
					index,
					nodeTable,
					err,
				)
			}
		}


	}
}

