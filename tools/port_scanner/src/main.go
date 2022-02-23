package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err error
	var host string
	var iPort, fPort uint64
	switch len(os.Args) {
	case 4:
		host = os.Args[1]
		iPort, err = strconv.ParseUint(os.Args[2], 10, 16)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fPort, err = strconv.ParseUint(os.Args[3], 10, 16)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case 3:
		host = os.Args[1]
		fPort, err = strconv.ParseUint(os.Args[3], 10, 16)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case 2:
		host = os.Args[1]

	default:
		fmt.Println("How to use: port_scanner <host> [port_min] [port_max]")
		return
	}

	run(host, iPort, fPort)
}
