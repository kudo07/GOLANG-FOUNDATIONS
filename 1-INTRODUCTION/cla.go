package main

import (
	"fmt"
	"os"
	"strconv"
)

func cla() {
	arguments := os.Args

	// os.Args always contains at least 1 element
	if len(arguments) == 1 {
		// If length = 1 → only program name exists → no real user input
		fmt.Println("need one or more arguments!")
		return
	}
	var min, max float64
	initialized := 0

	for i := 1; i < len(arguments); i++ {
		// string -> float64
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min ", min)
	fmt.Println("Max ", max)
}
