package db

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

func InsertRow(tableName string, item interface{})  ([]string, error) {

	if Session() == nil {
		err := invalidSession()
		return nil, err
	}

	results, err := r.Table(tableName).Insert(item).RunWrite(Session())
	if err != nil {
		return nil, fmt.Errorf("Unable to insert row: %s", err)
	}

	if results.Inserted == insertedSuccessful {
		return results.GeneratedKeys, nil
	}
	return nil, fmt.Errorf("Unable to insert row `%s`.", item)
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

func DeleteByIndex(tableName string, index interface{}, args ...interface{}) error {

	response, err := r.Table(tableName).GetAllByIndex(index, args...).Delete().RunWrite(Session())
	if err != nil {
		return fmt.Errorf("Error tring to delete: %s\n", err)
	}

	if response.Deleted <= 0 {
		return fmt.Errorf("Nothing was deleted.")
	}

	return nil
}
