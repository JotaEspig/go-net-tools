package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/TwiN/go-color"
)

const (
	PORT_MIN_D = 1    // Default value of the first port to check
	PORT_MAX_D = 1024 // Default value of the last port to check
)

func run(host string, iPort uint64, fPort uint64) {
	var i, j int
	var openPorts []int
	wg := &sync.WaitGroup{}

	if iPort == 0 {
		iPort = PORT_MIN_D
	}
	if fPort == 0 {
		fPort = PORT_MAX_D
	}

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
