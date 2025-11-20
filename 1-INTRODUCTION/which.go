package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func which() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a command name as an argument")
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		extensions := []string{""}
		if filepath.Ext(file) == "" {
			extensions = []string{"", ".exe", ".cmd", ".bat", ".com"}
		}
		for _, ext := range extensions {
			testPath := fullPath + ext
			fileInfo, err := os.Stat(testPath)
			if err != nil {
				continue
			}
			mode := fileInfo.Mode()
			if !mode.IsRegular() {
				continue
			}
			if mode&0111 != 0 || ext != "" {
				fmt.Println(testPath)
				return
			}
		}
	}
	fmt.Printf("%s not found in PATH\n", file)
}
