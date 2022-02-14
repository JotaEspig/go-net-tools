package dnsresolver

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

// Runs the dns resolver
func Run(host string, file string) bool {
	if file == "" {
		addr, err := net.LookupIP(host)
		if err != nil {
			fmt.Println("Unknown host")
			return false
		}
		fmt.Printf("%s --> %s\n",
			host, color.Ize(color.Red, addr[0].String()))
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return false
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
	return true
}
