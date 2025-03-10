package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func CreateCommit(message string) error {
	author, err := ReadFromFile(configPath)
	if err != nil {
		return fmt.Errorf("ERROR: author not configured. Use '--config' command first")
	}

	trackedFiles, err := GetTrackedFiles()
	if err != nil {
		return fmt.Errorf("ERROR: failed to get tracked files")
	} 

	if len(trackedFiles) == 0 {
		return fmt.Errorf("ERROR: no files are currently being tracked")
	}

	lastCommitID, err := getLastCommitID() 
	if err != nil {
		fmt.Printf("WARNING: Could not get last commit ID: %v\n", err)
	}

	changedFiles := []string{}
	for _, file := range trackedFiles {
		changed, err := hasFileChanged(file, lastCommitID)
		if err != nil {
			fmt.Printf("WARNING: could not check if file changed: %s - %v\n", file, err)
			changedFiles = append(changedFiles, file)
			continue
		}

		if changed {
			changedFiles = append(changedFiles, file)
		}
	}

	if len(changedFiles) == 0 && lastCommitID != "" {
		return fmt.Errorf("ERROR: no changes detected in tracked files since last commit")
	}

	commitID := GenerateCommitID(author)
	commitDir := filepath.Join(mainVCSPath, "commits", commitID)

	err = os.MkdirAll(commitDir, 0777)
	if err != nil {
		return fmt.Errorf("ERROR: failed to create commit directory: %v", err)
	}

	err = createCommitMetadata(commitDir, author, message, changedFiles)
	if err != nil {
		return err
	}

	err = saveTrackedFiles(commitDir, trackedFiles)
	if err != nil {
		return err
	}

	err = updateCommitLog(commitID, author, message, len(changedFiles))
	if err != nil {
		return fmt.Errorf("ERROR: failed to update commit log: %v", err)
	}

	if len(changedFiles) > 0 {
		fmt.Printf("Files changed: %d\n", len(changedFiles))
	}

	return nil
}

func hasFileChanged(filePath string, lastCommitID string) (bool, error) {
	currentHash, err := calculateFileHash(filePath)
	if err != nil {
		return false, fmt.Errorf("Failed to calculate hash for current file: %v", err)
	}

	if lastCommitID == "" {
		return true, nil
	}

	lastCommitFilePath := filepath.Join(mainVCSPath, "commits", lastCommitID, filePath)
	if _, err := os.Stat(lastCommitID); os.IsNotExist(err) {
		return true, nil
	}

	previousHash, err := calculateFileHash(lastCommitFilePath)
	if err != nil {
		return false, fmt.Errorf("Failed to calculate hash for previous version: %v", err)
	}

	return currentHash != previousHash, nil

}


func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil 
}


func createCommitMetadata(commitDir, author, message string, changedFiles []string) error {
	sort.Strings(changedFiles)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("author: %s\n", author))
	sb.WriteString(fmt.Sprintf("date: %s\n", GetCurrentTime()))
	sb.WriteString(fmt.Sprintf("message: %s\n", message))
	sb.WriteString(fmt.Sprintf("changed_files: %d\n", len(changedFiles)))

	if len(changedFiles) > 0 {
		sb.WriteString("files: \n")
		for _, file := range changedFiles {
			sb.WriteString(fmt.Sprintf(" - %s \n", file))
		}
	}

	err := WriteToFile(filepath.Join(commitDir, "metadata"), sb.String())
	if err != nil {
		return fmt.Errorf("ERROR: failed to write commit metadata: %v", err)
	}

	return nil
}

func saveTrackedFiles(commitDir string, trackedFiles []string) error {
	for _, file := range trackedFiles {
		content, err := ReadFromFile(file)
		if err != nil {
			fmt.Printf("WARNING: could not read from file %s - %v\n", file, err)
			continue
		}

		err = WriteToFile(filepath.Join(commitDir, file), content)
		if err != nil {
			fmt.Printf("WARNING: could not save file %s in commit: %v\n", file, err)
			continue
		}
	}

	return nil
}

func updateCommitLog(commitID, author, message string, changedFilesCount int) error {
	timestamp := GetCurrentTime()
	commitEntry := fmt.Sprintf("%s|%s|%s|%s|%d\n", commitID, author, timestamp, message, changedFilesCount)

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(commitEntry)
	return err
}

func GenerateCommitID(authorName string) string {
    return fmt.Sprintf("%s%d", authorName, GetUnixTimestamp())
}

func getLastCommitID() (string, error) {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		return "", nil
	}

	content, err := ReadFromFile(logPath)
	if err != nil {
		return "", nil
	}

	lines := strings.Split(content, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) >= 1 {
			return parts[0], nil
		}
	}

	return "", nil
}

func GetCurrentTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d-%02d-%d %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
}

func GetUnixTimestamp() int64 {
    return time.Now().Unix()
}
