package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	var sum int
	for _, line := range lines {

		nHalf := len(line) / 2
		c1, c2 := make(map[rune]int, nHalf), make(map[rune]int, nHalf)

		firstHalf, secondHalf := line[:len(line)/2], line[len(line)/2:]
		for i := 0; i < nHalf; i++ {
			c1[rune(firstHalf[i])]++
		}
		for i := 0; i < nHalf; i++ {
			c2[rune(secondHalf[i])]++
		}

		for k := range c1 {
			if _, found := c2[k]; found {
				if strings.ToLower(string(k)) == string(k) {
					sum += int(k-'a') + 1
				} else {
					sum += int(k-'A') + 27
				}
			}
		}
	}

	fmt.Println("item priorities sum:", sum)
}
