package main

const (
	CONFIG      = "--config"
	COMMIT      = "--commit"
	CHECKOUT    = "--checkout"
	HELP        = "--help"
	ADD         = "--add"
	LOG         = "--log"
	TRACKED     = "--tracked-files"
	SHOW_COMMIT = "--show-commit"
)

const configPath = "./.tinygit/config"
const indexPath = "./.tinygit/index"
const commitPath = "./.tinygit/commits/"
const logPath = "./.tinygit/commit_log"
const mainVCSPath = "./.tinygit/"

const BlueColor    = "\033[34m"
const ResetColor   = "\033[0m"
const YellowColor  = "\033[33m"
const GreenColor   = "\033[32m"
const MagentaColor = "\033[35m"
