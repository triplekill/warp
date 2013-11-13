package db

import (
	"fmt"

	r "github.com/dancannon/gorethink"

	"warp/utils"
)

var primary_session *r.Session
var primary_db = "warpdb"

var dbReady = false

func IsDbReady() bool {
	return dbReady
}

// Db returns the current database being used.
func Db() string {
	return primary_db
}

// SetDb sets the current session.
func SetDb(name string) error {

	if name == "" {
		return fmt.Errorf("Must provide a valid database name.")
	}

	primary_db = name
	return nil
}

// Session returns the current session being used.
func Session() *r.Session {
	return primary_session
}

func SetSession(session *r.Session) error {

	if session == nil {
		return fmt.Errorf("Must set a valid (not `nil`)  session.")
	}

	primary_session = session
	return nil
}

func invalidSession() error {

	return fmt.Errorf("No session is set. Db must be initialize before use.")
}

// Initialize creates a session/connection to the database and any necessary
// work before it can be used.
func Initialize(host, port string) (*r.Session, error) {

	var err error
	primary_session, err = r.Connect(map[string]interface{} {
		"address": fmt.Sprintf("%s:%s", host, port),
		"database": Db(),
	})
	if err != nil {
		return nil, err
	}

	dbReady = true
	return primary_session, nil
}

func Readify() {
	if !IsDbReady() {
		_, err := Initialize("localhost", "28015")
		if err != nil {
			panic(err)
		}
	}
}

// CreateTables is a helper to create multiple tables in one shot.
// Will return a panic if unable to create a table.
func CreateTables(tables []string) {

	for _, table := range tables {

		exist, err := DoesTableExist(table);
		if err != nil {
			utils.Panicf(
				"Could not verify table %q. Error: %s",
				table,
				err,
			)
		}

		if !exist {
			_, err := CreateTable(table)
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

// CreateIndices is a helper to create multiple indices in one shot.
// Will return a panic if unable to create an index.
func CreateIndices(tableName string, indices []string) {

	for _, index := range indices {

		indexExist, err := DoesIndexExist(index, tableName)
		if err != nil {
			utils.Panicf(
				"Could not verify index %q for table %q. Error: %s",
				index,
				tableName,
				err,
			)
		}

		if !indexExist {
			_, err := CreateIndex(index, tableName)
			if err != nil {
				utils.Panicf(
					"Unable to create index %q for table %q. Error: %s",
					index,
					tableName,
					err,
				)
			}
		}


	}
}
