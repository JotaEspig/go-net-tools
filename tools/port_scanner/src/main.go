package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var err error
	var host, portsRange string
	var ports []string
	var iPort, fPort int

	flag.StringVar(&host, "t", "", "Host that will be scanned")
	flag.StringVar(&portsRange, "p", "0-0",
		"The range of ports that will be scanned. Ex.: '1-1000' (ports 1 to 1000 will be scanned)")
	flag.Parse()

	ports = strings.Split(portsRange, "-")
	if len(ports) >= 1 {
		iPort, err = strconv.Atoi(ports[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if len(ports) == 2 {
		fPort, err = strconv.Atoi(ports[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	run(host, iPort, fPort)
}
