package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

// Runs the dns resolver
func run(host string, file string) {
	if file == "" {
		addr, err := net.LookupIP(host)
		if err != nil {
			fmt.Println("Unknown host")
			return
		}
		fmt.Printf("%s --> %s\n",
			host, color.Ize(color.Red, addr[0].String()))
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, subd := range lines {
		hostname := fmt.Sprintf("%s.%s", subd, host)
		addr, err := net.LookupIP(hostname)
		if err != nil {
			continue
		}
		fmt.Printf("%s --> %s\n",
			hostname, color.Ize(color.Red, addr[0].String()))
	}
}

func main() {
	args := os.Args
	run(args[1], args[2])
}