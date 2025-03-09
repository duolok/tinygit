package main

import (
	"os"
)

func main() {
	args := os.Args
	processArgs(args)
}

func processArgs(args []string) {
	switch {
	case len(args) < 2 || args[1] == "--help":
		PrintHelp()
	case len(args) == 2:
		PrintCommandInfo(args[1])
		CreateConfigDirectory()
	default:
		CreateConfigDirectory()
	}

}
