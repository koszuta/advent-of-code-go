package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
 *   --- Day 6: Custom Customs ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/6#part2
 */

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	sum := 0
	lines := make([]string, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()

		// Groups are separated by an empty line
		if line == "" {
			count := len(lines[0])

			// Iterate the chars in the first line
			for _, c := range lines[0] {
				// Check that each other line contains the char
				for i := 1; i < len(lines); i++ {
					if !strings.Contains(lines[i], string(c)) {
						// Reduce the common char count if the char isn't present in a line
						count--
						break
					}
				}
			}

			// Add the common char count to the sum
			sum += count
			lines = make([]string, 0, 0)

		} else {
			lines = append(lines, line)
		}
	}

	log.Printf("the sum of the counts is %d\n", sum)
}
