package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CreateCommit(message string) error {
	author, err := ReadFromFile(configPath)
	if err != nil {
		return fmt.Errorf("ERROR: author not configured. Use '--config' command first")
	}
	
	trackedFiles, err := GetTrackedFiles()
	if err != nil {
		return fmt.Errorf("ERROR: Failed to get tracked files")
	}
	
	if len(trackedFiles) == 0 {
		return fmt.Errorf("ERROR: no files are currently being tracked")
	}
	
	commitID := GenerateCommitID(author)
	commitDir := filepath.Join(mainVCSPath, "commits", commitID)
	
	err = os.MkdirAll(commitDir, 0777)
	if err != nil {
		return fmt.Errorf("ERROR: failed to create commit directory: %v", err)
	}
	
	err = createCommitMetadata(commitDir, author, message)
	if err != nil {
		return err
	}
	
	err = saveTrackedFiles(commitDir, trackedFiles)
	if err != nil {
		return err
	}
	
	err = updateCommitLog(commitID, author, message)
	if err != nil {
		return fmt.Errorf("ERROR: failed to update commit log: %v", err)
	}
	
	fmt.Printf("Committed as: %s\n", commitID)
	return nil
}

func createCommitMetadata(commitDir, author, message string) error {
	metadata := fmt.Sprintf("author: %s\ndate: %s\nmessage: %s\n",
		author, GetCurrentTime(), message)
	
	err := WriteToFile(filepath.Join(commitDir, "metadata"), metadata)
	if err != nil {
		return fmt.Errorf("ERROR: failed to write commit metadata: %v", err)
	}
	
	return nil
}

func saveTrackedFiles(commitDir string, trackedFiles []string) error {
	for _, file := range trackedFiles {
		content, err := ReadFromFile(file)
		if err != nil {
			fmt.Printf("WARNING: could not read from file: %s - %v\n", file, err)
			continue
		}
		
		fileDir := filepath.Dir(filepath.Join(commitDir, file))
		if fileDir != commitDir {
			err = os.MkdirAll(fileDir, 0777)
			if err != nil {
				fmt.Printf("WARNING: could not create directories for file %s: %v\n", file, err)
				continue
			}
		}
		
		err = WriteToFile(filepath.Join(commitDir, file), content)
		if err != nil {
			fmt.Printf("WARNING: could not save file %s in commit: %v\n", file, err)
			continue
		}
	}
	
	return nil
}

func updateCommitLog(commitID, author, message string) error {
	logPath := filepath.Join(mainVCSPath, "commit_log")
	
	timestamp := GetCurrentTime()
	commitEntry := fmt.Sprintf("%s|%s|%s|%s\n", commitID, author, timestamp, message)
	
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	
	_, err = file.WriteString(commitEntry)
	return err
}

// func LogAllCommits() {
// 	logPath := filepath.Join(mainVCSPath, "commit_log")
// 	
// 	if _, err := os.Stat(logPath); os.IsNotExist(err) {
// 		fmt.Println("No commits yet")
// 		return
// 	}
// 	
// 	content, err := ReadFromFile(logPath)
// 	if err != nil {
// 		fmt.Printf("ERROR: failed to read commit log: %v\n", err)
// 		return
// 	}
// 	
// 	if content == "" {
// 		fmt.Println("No commits yet")
// 		return
// 	}
// 	
// 	lines := strings.Split(content, "\n")
// 	for i := len(lines) - 1; i >= 0; i-- {
// 		line := lines[i]
// 		if line == "" {
// 			continue
// 		}
// 		
// 		parts := strings.Split(line, "|")
// 		if len(parts) != 4 {
// 			continue
// 		}
// 		
// 		fmt.Printf("commit %s\n", parts[0])
// 		fmt.Printf("Author: %s\n", parts[1])
// 		fmt.Printf("Date: %s\n", parts[2])
// 		fmt.Printf("\n    %s\n\n", parts[3])
// 	}
// }

func GenerateCommitID(authorName string) string {
    return fmt.Sprintf("%s%d", authorName, GetUnixTimestamp())
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
