package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

var input string = "3113322113"

func lookAndSay(s string) string {
	result := ""
	current := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == current {
			count++
		} else {
			result = result + strconv.Itoa(count) + string(current)
			current = s[i]
			count = 1
		}
	}
	result = result + strconv.Itoa(count) + string(current)
	return result
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("%v\n", time.Since(start))
	}(time.Now())

	n := 50
	for i := 0; i < n; i++ {
		input = lookAndSay(input)
	}
	fmt.Printf("lookAndSay(%d) len=%d\n", n, len(input))

	n = 50
	conwaysConstant := 1.30357726903429639125709911215255189073070250465940487575486139062855088785246155712681576686442522555
	fmt.Printf("conwaysConstant=%f\n", conwaysConstant)
	c := conwaysConstant
	for i := 0; i < n-1; i++ {
		c = math.Pow(c, conwaysConstant)
	}
	fmt.Printf("c=%f\n", c)

	fmt.Printf("lookAndSay(%d) len=%f\n", n, math.Pow(float64(len(input)), c))
}
