package db

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

// CreateTable creates a new table.
func CreateTable(tableName string) (bool, error) {

	if Session() == nil {
		return false, invalidSession()
	}

	results, err := r.Db(Db()).TableCreate(tableName).RunWrite(Session())
	if err != nil {
		return false, err
	}

	if results.Created == createdSuccessful {

		return true, nil
	}

	// *NOTE* Another way of getting results. For reference.
	//results, err := r.Db(dbName).TableCreate(tableName).Run(session)
	//if err != nil {
	//	return false, err
	//}

	//if results.Next() {
	//	var row TableResult
	//	results.Scan(&row)
	//	if row.Created == createdSuccssful {

	//		return true, nil
	//	}
	//}
	// -----------------------------------------------------------------

	return false, fmt.Errorf("Unable to create table `%s`.", tableName)
}

// DoesTableExist checks if a table already exist.
func DoesTableExist(tableName string) (bool, error) {

	if Session() == nil {
		return false, invalidSession()
	}

	result, err := r.Db(Db()).TableList().Run(Session())
	if err !=  nil {
		return false, err
	}

	if result.Next() {
		var tables []string
		result.Scan(&tables)

		for _, t := range tables {
			if t == tableName {
				return true, nil
			}
		}
	}

	return false, nil
}

func CreateIndex(indexName, tableName string) (bool, error) {

	results, err := r.Table(tableName).IndexCreate(indexName).RunWrite(Session())
	if err != nil {
		return false, fmt.Errorf("Unable to create index `%s` for table `%s`. %s",
					indexName, tableName, err,
				)
	}

	if results.Created == createdSuccessful {
		return true, nil
	}

	return false, fmt.Errorf("Unable to create index `%s` for table `%s`.",
				indexName, tableName,
			)
}

func DoesIndexExist(indexName, tableName string) (bool, error) {

	if Session() == nil {
		return false, invalidSession()
	}

	result, err := r.Db(Db()).Table(tableName).IndexList().Run(Session())
	if err !=  nil {
		return false, err
	}

	if result.Next() {
		var indices []string
		result.Scan(&indices)

		for _, i := range indices {
			if i == indexName {
				return true, nil
			}
		}
	}

	return false, nil
}
