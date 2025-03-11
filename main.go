package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
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
	case len(args) >= 2:
		gitArgs := Command{
			commandName: args[1],
			values:      args[2:],
		}
		CreateConfigDirectory()
		handleCommand(gitArgs)
	default:
		CreateConfigDirectory()
		PrintHelp()
	}
}

func handleCommand(command Command) {
	switch command.commandName {
	case CONFIG:
		handleConfig(command.values)
	case ADD:
		handleAdd(command.values)
	case COMMIT:
		handleCommit(command.values)
	case TRACKED:
		handleTrackedFiles()
	case LOG:
		handleLog()
	case SHOW_COMMIT:
		handleShowCommit(command.values)
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

func handleAdd(values []string) {
	files, err := GetTrackedFiles()
	if err != nil {
		panic(err)
	}

	if len(values) == 0 {
		ShowTrackedFiles(files)
		return
	}

	cleanFiles := RemoveDuplicates(files)
	for _, val := range values {
		if _, err := os.Stat(val); os.IsNotExist(err) {
			fmt.Printf("WARNING: %s does not exist in the current directory\n", val)
			continue
		}

		if slices.Contains(cleanFiles, val) {
			continue
		}

		cleanFiles = append(cleanFiles, val)
	}

	content := strings.Join(cleanFiles, "\n")
	WriteToFile(filepath.Join(mainVCSPath, "index"), content)
}

func handleCommit(values []string) {
	if len(values) != 1 {
		fmt.Println("ERROR: commit message required")
		return
	}
	message := values[0]

	err := CreateCommit(message)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handleShowCommit(values []string) {
	if len(values) != 1 {
		fmt.Println("ERROR: commit ID required")
		return
	}

	commitID := values[0]
	err := DisplayCommitDetails(commitID)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handleLog() {
	LogAllCommits()
}

func handleTrackedFiles() {
	files, err := GetTrackedFiles()
	if err != nil {
		panic(err)
	}
	ShowTrackedFiles(files)
}
