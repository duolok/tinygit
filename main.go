package main

import (
	"fmt"
	"os"
)

type Command struct {
	commandName string
	values      []string
}

func main() {
	args := os.Args
	processArgs(args)
}

func processArgs(args []string) {
	switch {
	case len(args) < 2 || args[1] == HELP:
		PrintHelp()
	case len(args) == 2:
		PrintCommandInfo(args[1])
	case len(args) > 2:
		gitArgs := Command{
			commandName: args[1],
			values:  args[2:],
		}
		CreateConfigDirectory()
		handleCommand(gitArgs)
	default:
		CreateConfigDirectory()
	}
}

func handleCommand(command Command) {
	switch command.commandName {
	case CONFIG:
		handleConfig(command.values)
	case ADD:
		handleAdd()
	default:
		return
	}
}

func handleConfig(values []string) {
	if len(values) != 1 {
		fmt.Println("ERROR: only one value should be used as author name")
		return
	}
	WriteToFile(configPath, values[0])
}

func handleAdd() {}
