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
