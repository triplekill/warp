package main

import (
	"fmt"
	"reflect"
)

func main() {

	out, err := remoteNetstatCmd("toppatch", "test.toppatch.com")
	if err != nil {
		fmt.Println(err)
	}

	output := string(out)

	//fmt.Println(output)

	ports := parseNetstatPorts(output)
	for x := 0; x < len(ports); x++ {
		fmt.Println(ports[x])
		fmt.Println("type:", reflect.TypeOf(ports[x]))
	}

}
