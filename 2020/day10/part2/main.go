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

	// Add the outlet
	jolts = append(jolts, 0)

	// Sort the jolts
	sort.Ints(jolts)

	// Add the built-in joltage adapter
	// which is 3 jolts higher than the highest-rated adapter
	builtInAdapter := jolts[len(jolts)-1] + 3
	jolts = append(jolts, builtInAdapter)

	// Find the number of unique paths from the outlet to the built-in adapter
	pathCounts := make(map[int]int)
	pathCounts[0] = 1
	for i := 1; i < len(jolts); i++ {
		sum := 0
		for j := i - 1; j >= 0 && jolts[i]-jolts[j] <= 3; j-- {
			sum += pathCounts[jolts[j]]
		}
		pathCounts[jolts[i]] = sum
	}

	paths := pathCounts[builtInAdapter]
	log.Printf("the total number of distinct ways you can arrange the adapters is %d\n", paths)
}
