package main

import (
	"fmt"
	"time"
)

var row, col int = 2947, 3029

// var row, col int = 6, 6

func main() {
	r := row + col - 1
	s := r*(r-1)/2 + col
	fmt.Printf("row %d col %d = %d\n", row, col, s)

	start := time.Now()
	var code, prev int
	prev = 20151125
	for i := 1; i < s; i++ {
		code = (prev * 252533) % 33554393
		prev = code
	}
	fmt.Printf("Part 1 took %v\n", time.Since(start))
	fmt.Printf("code=%d\n", code)
}
