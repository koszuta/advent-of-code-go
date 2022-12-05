package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(b), "\n")

	sum, max := 0, 0
	for _, line := range lines {
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		}

		calories, _ := strconv.Atoi(line)
		sum += calories
	}

	fmt.Println("most calories:", max)
}
