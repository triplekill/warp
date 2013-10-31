package db

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

func InsertRow(tableName string, item interface{})  (bool, error) {

	if Session() == nil {
		return invalidSession()
	}

	results, err := r.Table(tableName).Insert(item).RunWrite(Session())
	if err != nil {
		return false, fmt.Errorf("Unable to insert row: %s", err)
	}

	if results.Inserted == insertedSuccessful {
		return true, nil
	}
	return false, fmt.Errorf("Unable to insert row `%s`.", item)
}
