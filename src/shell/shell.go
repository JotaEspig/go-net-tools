package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/TwiN/go-color"
)

type Shell struct {
	isWindows bool
	filepath  string
}

func execCommand(in string, args ...string) error {
	cmd := exec.Command(in, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Init() *Shell {
	var s *Shell
	ex, err := os.Executable()
	if err != nil {
		return nil
	}
	path := filepath.Dir(ex)
	isW := runtime.GOOS == "windows"
	s = &Shell{
		isWindows: isW,
		filepath:  path,
	}
	return s
}

// Executes a command in shell or executes a tool if it exists
func (s *Shell) execute(cmd string) error {
	args := strings.Split(cmd, " ")
	if args[0] == "cd" {
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		return nil
	}

	if value, ok := validTools[args[0]]; ok {
		toolpath := s.filepath + value + args[0]
		if s.isWindows {
			toolpath = strings.Replace(toolpath, "/", "\\", -1)
			toolpath += ".exe"
		}
		if len(args) >= 2 {
			return execCommand(toolpath, args[1:]...)
		} else {
			return execCommand(toolpath, []string{}...)
		}
	}

	if len(args) >= 2 {
		return execCommand(args[0], args[1:]...)
	}
	return execCommand(args[0], []string{}...)
}

func (s *Shell) Run() {
	var r, path string
	var in *bufio.Scanner
	for {
		path, _ = os.Getwd()
		fmt.Printf("%s\n%s@%s> ",
			color.Ize(color.Cyan, path),
			color.Ize(color.Red, config["name"]),
			color.Ize(color.Yellow, config["version"]))

		in = bufio.NewScanner(os.Stdin)
		in.Scan()
		r = in.Text()

		if r == "exit" {
			break
		}

		err := s.execute(r)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
