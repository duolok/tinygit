package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LogAllCommits() {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		fmt.Println("No commits yet")
		return
	}

	content, err := ReadFromFile(logPath)
	if err != nil {
		fmt.Printf("ERROR: failed to read commit log: %v\n", err)
		return
	}

	if content == "" {
		fmt.Println("No commits yet")
		return
	}

	hasCommits := false
	lines := strings.Split(content, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		hasCommits = true
		fmt.Printf("commit %s\n", parts[0])
		fmt.Printf("Author: %s\n", parts[1])
		fmt.Printf("Date: %s\n", parts[2])

		if len(parts) >= 5 {
			fmt.Printf("Changed files :%s \n", parts[4])
		}

		fmt.Printf("\n		%s\n\n", parts[3])
	}

	if !hasCommits {
		fmt.Println("No valid commits found in log")
	}
}

func DisplayCommitDetails(commitID string) error {
	commitDir := filepath.Join(mainVCSPath, "commits", commitID)
	
	if _, err := os.Stat(commitDir); os.IsNotExist(err) {
		return fmt.Errorf("ERROR: commit %s does not exist", commitID)
	}
	
	metadataPath := filepath.Join(commitDir, "metadata")
	metadata, err := ReadFromFile(metadataPath)
	if err != nil {
		return fmt.Errorf("ERROR: failed to read commit metadata: %v", err)
	}
	
	fmt.Printf("Commit: %s\n\n", commitID)
	fmt.Println("Metadata:")
	fmt.Println(metadata)
	
	fmt.Println("\nFiles in commit:")
	walkErr := filepath.Walk(commitDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if path == commitDir || path == metadataPath {
			return nil
		}
		
		if info.IsDir() {
			return nil
		}
		
		relPath, err := filepath.Rel(commitDir, path)
		if err != nil {
			return err
		}
		
		fmt.Printf("  %s\n", relPath)
		return nil
	})
	
	if walkErr != nil {
		return fmt.Errorf("ERROR: failed to list files in commit: %v", walkErr)
	}
	
	return nil
}
