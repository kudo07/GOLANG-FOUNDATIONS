package main

import (
	"fmt"
	"os"
	"strconv"
)

func process() {
	arguments := os.Args
	fmt.Println(arguments)
	// ["filpath", 10(arguments),12,12,...]
	if len(arguments) == 1 {
		return
	}
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		invalid = append(invalid, k)
	}
	fmt.Println("#read", total, "#ints", nInts, "floats", nFloats)
	if len(invalid) > total {
		fmt.Println("too much invalid")
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
