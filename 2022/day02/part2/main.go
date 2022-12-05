package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(b), "\n")

	var score int
	for _, line := range lines {
		shape := int(line[0] - 'A')
		switch line[2] {
		case 'X': // lose
			score += (shape+2)%3 + 1
		case 'Y': // draw
			score += 3
			score += shape + 1
		case 'Z': // win
			score += 6
			score += (shape+1)%3 + 1
		}
	}

	fmt.Println("score:", score)
}
