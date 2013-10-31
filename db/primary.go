package db

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

var primary_session *r.Session
var primary_db string = "warpdb"

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

func invalidSession() (bool, error) {

	return false, fmt.Errorf("Not session is set. Db must be initialize before use.")
}

// Initialize creates a session/connection to the database and any necessary
// work before it can be used.
func Initialize(host, port string) (*r.Session, error) {

	var err error
	primary_session, err = r.Connect(map[string]interface{} {
		"address": fmt.Sprintf("%s:%s", host, port),
		"database": primary_db,
	})
	if err != nil {
		return nil, err
	}

	return primary_session, nil
}
