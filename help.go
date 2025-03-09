package main

import "fmt"

func PrintHelp() {
	fmt.Println(`[USAGE]   | All TinyGit commands
                --config   | sets or outputs the name of a commit author
                --help     | prints the help page
                --add      | adds a file to the list of tracked files or outputs the list
                --log      | shows all commits
                --commit   | saves the file changes and the author name
                --checkout | allows you to switch between commits and restore previous file state`)
}

func PrintCommandInfo(command string) {
	switch command {
	case CONFIG:
		{
			fmt.Println(`--config   | sets or outputs the name of a commit author
                        [USAGE] ./tinygit --config {NAME}`)
		}
	case ADD:
		{
			fmt.Println(`--add   |  adds a file to the list of tracked files or outputs the list
                        [USAGE] ./tinygit --add {FILE_NAME}`)
		}
	case LOG:
		{

			fmt.Println(`--log      | shows all commits
                        [USAGE] ./tinygit --log`)
		}
	case COMMIT:
		{

			fmt.Println(`--commit   | saves the file changes and the author name
                [USAGE] ./tinygit --commit {FILE_NAME}`)
		}

	case CHECKOUT:
		{

			fmt.Println(`--checkout | allows you to switch between commits and restore previous file state
                [USAGE] ./tinygit --checkout {COMMIT_HASH}`)
		}
	case HELP:
		PrintHelp()
	default:
		fmt.Println("Unknown command: ", command, "| try --help for more info")
	}
}
