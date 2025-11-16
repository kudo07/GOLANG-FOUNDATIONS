package main

import (
	"fmt"
	"math"
)

var (
	Global        int = 1234
	AnotherGlobal     = -32453
)

func variables() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Initial j value: ", j)
	j = Global
	// math.Abs() requires a float64
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d,j=%d, k=%.2f\n", Global, i, j, k)
}
