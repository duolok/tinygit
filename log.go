package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func LogAllCommits() {
    entries, err := os.ReadDir(path.Join(mainVCSPath, "commits"))
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(entry)
		}
	}
}
