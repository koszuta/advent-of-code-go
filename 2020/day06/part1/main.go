package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 6: Custom Customs ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/6
 */

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	sum := 0
	commonChars := make(map[rune]struct{})
	for scanner.Scan() {
		line := scanner.Text()

		// Groups are separated by an empty line
		if line == "" {
			// Add the common char count to the sum
			sum += len(commonChars)
			commonChars = make(map[rune]struct{})

		} else {
			// Add this line's chars to the set of common chars
			for _, c := range line {
				commonChars[c] = struct{}{}
			}
		}
	}

	log.Printf("the sum of the counts is %d\n", sum)
}
