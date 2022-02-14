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
	var host, filename string
	if len(os.Args) == 3 {
		host = os.Args[1]
		filename = os.Args[2]
	} else if len(os.Args) == 2 {
		host = os.Args[1]
	} else {
		fmt.Println("How to use:\ndns_resolver <host> [wordlist of subdomains]")
		return
	}

	run(host, filename)
}
