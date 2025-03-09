package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 || args[1] == "--help" {
		PrintHelp()
		return
	} else if len(args) == 2 {
		PrintCommandInfo(args[1])
	}
}
