package remote

import (

	"fmt"

	"warp/db"
	"warp/utils"
)

const (
	destTable = "destination"
	destPortIndex = "local_port"
)

var allDestTables = []string {
	destTable,
}

var allDestIndices = []string {
	destPortIndex,
}

func init() {

	db.Readify()

	db.CreateTables(allDestTables)
	db.CreateIndices(destTable, allDestIndices)
}

var destPortRange = utils.NumberRange(10000, 11000)

type Destination struct {

	Id string		`gorethink:"id, omitempty"`
	LocalPort string	`gorethink:"local_port"`
	NodeId string		`gorethink:"node_id"`
}

func NewDestination(nodeId string) (*Destination, error) {

	d := &Destination{
		LocalPort: "555",
		NodeId: nodeId,
	}

	return d, nil
}

func SaveDestination(dest *Destination) error {

	exist, err := DoesDestinationExist(dest.LocalPort)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("Destination with port %q already exist.", dest.LocalPort)
	}

	_, err = db.InsertRow(destTable, *dest)
	if err != nil {
		return err
	}

	return nil
}

func DoesDestinationExist(port string) (bool, error) {

	items, err := db.GetByIndex(
		destTable,
		destPortIndex,
		port,
	)
	if err != nil {
		return false, err
	}

	if len(items) > 0 {
		return true, nil
	}

	return false, nil
}
