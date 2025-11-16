package main

import "fmt"

func readcli() {
	fmt.Println("Please give me your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("your name is ", name)
}
