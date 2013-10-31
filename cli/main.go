package main

import (
	"fmt"

	"warp/warplib/remote"
	"warp/db"
)

//var session *r.Session

type Slave struct {
	Id string `gorethink:"id,omitempty"`
	Name string
	Age int
}

func main() {

	//out, err := remoteNetstatCmd("toppatch", "test.toppatch.com")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//_, err := r.Connect(map[string]interface{} {

	//	"address": "localhost:28015",
	//	"database": "Slark",
	//})

	//if err != nil {

	//	fmt.Println("Unable to connect.")
	//}

	//results, err := r.Db(dbName).TableCreate(tableName).RunWrite(session)
	//if err != nil {
	//	return false, err
	//}

	_, err := db.Initialize("localhost", "28015")
	if err != nil {
		fmt.Println("Could not init db")
	}

	slave := &Slave{
		Name: "miguel",
		Age: 12,
	}
	fmt.Println(db.CreateTable("slaves", "warpdb"))
	fmt.Println(db.InsertRow("slaves", slave))

	fmt.Println(db.CreateIndex("name", "slaves"))

	//fmt.Println(db.DoesTableExist("testers22s111", "test", nil))


	fmt.Println(remote.UsablePort("toppatch", "test.toppatch.com"))
	fmt.Println(remote.IsPortAvailable(15000, "toppatch", "test.toppatch.com"))

	//ports := parseNetstatPorts(output)
	//for x := 0; x < len(ports); x++ {
	//	fmt.Println(ports[x])
	//	fmt.Println("type:", reflect.TypeOf(ports[x]))
	//}
}
