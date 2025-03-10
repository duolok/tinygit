package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteToFile(filePath, content string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetTrackedFiles() ([]string, error) {
	data, err := os.ReadFile(filepath.Join(mainVCSPath, "index"))
	if err != nil {
		return nil, err
	}

	content := string(data)
	files := strings.Split(content, "\n")

	var result []string
	for _, file := range files {
		result = append(result, file)
	}

	return result, nil
}

func ShowTrackedFiles(files []string) {
	for _, el := range files[1:] {
		fmt.Println("-", el)
	}
}

func RemoveDuplicates(files []string) []string {
	seen := make(map[string]bool)

	var result []string
	for _, file := range files {
		if !seen[file] {
			seen[file] = true
			result = append(result, file)
		}
	}

	return result
}


func CreateConfigDirectory() {
	if err := ensureDir(mainVCSPath); err != nil {
		fmt.Println("Directory creation failed. Err: ", err.Error())
		os.Exit(1)
	}

	if err := ensureDir(filepath.Join(mainVCSPath, "commits")); err != nil {
		fmt.Println("Directory creation failed. Err: ", err.Error())
		os.Exit(1)
	}

	if err := ensureFile("config"); err != nil {
		fmt.Println("Config creation failed. Err: ", err.Error())
		os.Exit(1)
	}

	if err := ensureFile("index"); err != nil {
		fmt.Println("Index creation failed. Err: ", err.Error())
		os.Exit(1)
	}
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0777) 
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func ensureFile(fileName string) error {
	fullPath := filepath.Join(mainVCSPath, fileName)

	_, err := os.Stat(fullPath)
	if err == nil {
		return nil 
	} else if !os.IsNotExist(err) {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil{
		return err
	} 
	defer f.Close()
	return nil
}
