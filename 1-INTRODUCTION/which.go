package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func which() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide an argument")
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
}
