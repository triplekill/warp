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

func GetByIndex(tableName string, index interface{}, args ...interface{}) ([]interface{}, error) {

	//kv := make([]interface{}, 0)

	//for k, v := range args {

	//	kv = append(kv, k, v)
	//}

	results, err := r.Table(tableName).GetAllByIndex(index, args...).Run(Session())
	if err != nil {
		return nil, fmt.Errorf("Unable to verify index %q. Error: %s", index, err)
	}

	items := make([]interface{}, 0)

	for results.Next() {

		var i interface{}

		err := results.Scan(&i)
		if err != nil {
			continue
		}

		items = append(items, i)
	}

	return items, nil
}

func IndexedValueExist(table string, args map[string]interface{}) {

}
