package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/TwiN/go-color"
)

const (
	PORT_MIN = 1
	PORT_MAX = 1024
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
		iPort = PORT_MIN
		fPort, err = strconv.ParseUint(os.Args[3], 10, 16)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case 2:
		host = os.Args[1]
		iPort = PORT_MIN
		fPort = PORT_MAX

	default:
		fmt.Println("How to use: port_scanner <host> [port_min] [port_max]")
		//return
	}

	run(host, iPort, fPort)
}

func run(host string, iPort uint64, fPort uint64) {
	var i, j int
	var openPorts []int
	wg := &sync.WaitGroup{}

	fmt.Println(host)
	for i = int(iPort); i <= int(fPort); i += 100 {
		iP := i       // initial port
		fP := i + 100 // final port
		wg.Add(100)
		for j = iP; j < fP; j++ {
			port := j
			go func() {
				if scanPort(host, port) {
					openPorts = append(openPorts, port)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
	sort.Ints(openPorts)
	printPorts(openPorts)
}

func scanPort(host string, port int) bool {
	addr := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func printPorts(ports []int) {
	for _, port := range ports {
		fmt.Println(color.Ize(color.Cyan, strconv.Itoa(port)))
	}
}
