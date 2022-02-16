package main

import (
	"fmt"
	"my-cli/shell"

	"github.com/TwiN/go-color"
)

func main() {
	s := "\n\n================ %s %s ================\n"
	s += "==============  Made by Jota ==============\n\n\n"
	fmt.Printf(s,
		color.Ize(color.Cyan, "NET"), color.Ize(color.Yellow, "TOOLS"))
	sh := shell.Init()
	sh.Run()

}
