package main

import (
	"fmt"
	"os"
	"strconv"
)

func control() {
	if len(os.Args) != 2 {
		fmt.Println(("Please provide a command line argument"))
		return
	}
	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println(("Zero!"))
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("value: ", argument)
	}
	// convert string into an integer
	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to integer", err)
		return
	}

	switch {
	case value == 0:
		fmt.Println("positive integer")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen")
	}
}
