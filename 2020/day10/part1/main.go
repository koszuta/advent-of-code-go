package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// Build a list of joltage adapters
	jolts := make([]int, 0, 0)
	for scanner.Scan() {
		jolt, _ := strconv.Atoi(scanner.Text())
		jolts = append(jolts, jolt)
	}

	// Sort the jolts
	sort.Ints(jolts)

	// Add the built-in joltage adapter
	// which is 3 jolts higher than the highest-rated adapter
	maxJolts := jolts[len(jolts)-1]
	jolts = append(jolts, maxJolts+3)

	// Count the number of 1 and 3 jolt differences between consecutive adapters
	prev, oneDiff, threeDiff := 0, 0, 0
	for _, jolt := range jolts {
		switch jolt - prev {
		case 1:
			oneDiff++
		case 3:
			threeDiff++
		}
		prev = jolt
	}

	log.Printf("there are %d differences of 1 jolt and %d differences of 3 jolts\n", oneDiff, threeDiff)
	log.Printf("answer: %d\n", oneDiff*threeDiff)
}
