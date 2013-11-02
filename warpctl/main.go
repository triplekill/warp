package main

import (
	"fmt"
	"flag"
	"os"
	"strings"


//	"warp/warplib/remote"
//	"warp/db"
)


var commands = []*Command {

	nodeCommand,
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	//if args[0] == "help" {
	//	help(args[1:])
	//	return
	//}


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

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {

				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}

			cmd.Run(cmd, args)
			return
		}
	}

	//_, err := db.Initialize("localhost", "28015")
	//if err != nil {
	//	fmt.Println("Could not init db")
	//}

	//slave := &Slave{
	//	Name: "miguel",
	//	Age: 12,
	//}

	//fmt.Println(db.CreateTable("slaves", "warpdb"))
	//fmt.Println(db.InsertRow("slaves", slave))

	//fmt.Println(db.CreateIndex("name", "slaves"))

	//fmt.Println(db.DoesTableExist("testers22s111", "test", nil))


	//fmt.Println(remote.UsablePort("toppatch", "test.toppatch.com"))
	//fmt.Println(remote.IsPortAvailable(15000, "toppatch", "test.toppatch.com"))

	//ports := parseNetstatPorts(output)
	//for x := 0; x < len(ports); x++ {
	//	fmt.Println(ports[x])
	//	fmt.Println("type:", reflect.TypeOf(ports[x]))
	//}

	fmt.Fprintf(os.Stderr, "warp: unknown command %q\n", args[0])
}

func usage() {

	fmt.Println("Not using warpctl correctly.")
	os.Exit(1)
}


// A Command is an implementation of a go command
// like go build or go fix.
type Command struct {

	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	Short string

	// Long is the long message shown in the 'go help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own flag parsing.
	CustomFlags bool
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}

	return name
}

func (c* Command) Usage() {

	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}
