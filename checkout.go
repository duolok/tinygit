package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Checkout(commitID string) error {
	commitDir := filepath.Join(mainVCSPath, "commits", commitID)
	metadataPath := filepath.Join(commitDir, "metadata")

	if _, err := os.Stat(commitDir); os.IsNotExist(err) {
		return fmt.Errorf("ERROR: commit %s does not exist", commitID)
	}

	walkErr := filepath.Walk(commitDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == commitDir || path == metadataPath {
			return nil
		}

		relPath, err := filepath.Rel(commitDir, path)
		if err != nil {
			return fmt.Errorf("ERROR: couldn't get relative path: %v", err)
		}

		destPath := filepath.Join(".", relPath)
		destDir := filepath.Dir(destPath)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return fmt.Errorf("ERROR: couldn't create directory %s: %v", destDir, err)
		}

		content, err := ReadFromFile(path)
		if err != nil {
			fmt.Printf("ERROR: couldn't read from file %s - %v", path, err)
			return err
		}

		err = WriteToFile(destPath, content)
		if err != nil {
			fmt.Printf("ERROR: couldn't restore file %s - %v", path, err)
			return err
		}

		return nil
	})

	if walkErr != nil {
		return fmt.Errorf("ERROR: failed to checkout to a commit: %v", walkErr)
	}

	fmt.Printf("%sSuccessfuly checked out at commit%s: %s", GreenColor, ResetColor, commitID)
	return nil
}
