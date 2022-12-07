package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	var score int
	for _, line := range lines {
		// Points for the shape
		switch line[2] {
		case 'X':
			score += 1
		case 'Y':
			score += 2
		case 'Z':
			score += 3
		}

		// Points for the outcome
		switch {
		case line[0]-'A' == line[2]-'X': // draw
			score += 3
		case line[0] == 'A' && line[2] == 'Y', // win
			line[0] == 'B' && line[2] == 'Z',
			line[0] == 'C' && line[2] == 'X':
			score += 6
		}
	}

	fmt.Println("score:", score)
}
