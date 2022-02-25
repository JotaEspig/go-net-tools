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

func run(host string, iPort int, fPort int) {
	var i, j int
	var openPorts []int
	wg := &sync.WaitGroup{}

	if iPort == 0 {
		iPort = fPort
	}
	if fPort == 0 {
		fPort = PORT_MAX_D
		iPort = PORT_MIN_D
	}

	t_amount := 100

	addr, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("Could not resolve host")
		return
	}

	fmt.Printf("Start scanning %s (%s) on ports %s -> %s\n",
		color.Ize(color.Green, host),
		addr[0].String(),
		color.Ize(color.Cyan, fmt.Sprint(iPort)),
		color.Ize(color.Cyan, fmt.Sprint(fPort)))

	host = addr[0].String()
	for i = iPort; i <= fPort; i += t_amount {
		iP := i            // initial port
		fP := i + t_amount // final port
		for j = iP; j < fP; j++ {
			if j > fPort {
				break
			}
			wg.Add(1)
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
	if len(openPorts) > 0 {
		sort.Ints(openPorts)
		printPorts(openPorts)
	} else {
		fmt.Println("There is no open ports in this host")
	}
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
