package main

import (
	"fmt"

	"github.com/TwiN/go-color"
)

func main() {
	s := "\n\t======== %s %s ========\n\t======  Made by Jota ======\n\n"
	fmt.Printf(s,
		color.Ize(color.Green, "NET"), color.Ize(color.Cyan, "TOOLS"))
}
