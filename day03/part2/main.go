package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(b), "\n")

	var sum int
	for i := 0; i < len(lines); i += 3 {
		itemCounts := make(map[rune]int)

		for j := i; j < i+3; j++ {
			unique := make(map[rune]struct{})
			for _, item := range lines[j] {
				unique[item] = struct{}{}
			}
			for item := range unique {
				itemCounts[item]++
			}
		}

		var badge rune
		for item, count := range itemCounts {
			if count == 3 {
				badge = item
				break
			}
		}

		if strings.ToLower(string(badge)) == string(badge) {
			sum += int(badge-'a') + 1
		} else {
			sum += int(badge-'A') + 27
		}
	}

	fmt.Println("badge priorities sum:", sum)
}
