package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteToFile(filePath, content string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}

}

func CreateConfigDirectory() {
	if err := ensureDir("./.tinygit"); err != nil {
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
	fullPath := filepath.Join("./.tinygit/", fileName)

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
