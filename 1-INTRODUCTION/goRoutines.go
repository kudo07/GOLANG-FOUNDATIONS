package main

import (
	"fmt"
	"time"
)

func callGoRoutines(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Println(i, "-----")
	}
}

func goRoutines() {
	for i := 0; i < 4; i++ {
		go callGoRoutines(i, 5)
	}
	time.Sleep(time.Second)
}
