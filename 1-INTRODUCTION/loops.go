package main

import "fmt"

func loops() {
	for i := 0; i < 10; i++ {
		fmt.Println(i*i, " ")
	}
	// for in while , while is not support directly
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Println(i*i, " ")
		i++
	}
	aSlice := []int{-1, 2, 4, 5, 6}
	for i, v := range aSlice {
		fmt.Println("index: ", i, "value ", v)
	}
}
