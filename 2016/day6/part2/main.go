package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

/*
 *   --- Day 6: Signals and Noise ---
 *           --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/6#part2
 */

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	columnCounts := make([]map[rune]int, 0)
	for i := range columnCounts {
		columnCounts[i] = make(map[rune]int)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// init number of columns based on input
		if len(columnCounts) == 0 {
			for i := 0; i < len(line); i++ {
				columnCounts = append(columnCounts, make(map[rune]int))
			}
		}

		// increment the character count for each column
		for i, c := range line {
			columnCounts[i][c]++
		}
	}

	message := ""
	for _, counts := range columnCounts {
		var minChar rune
		minCount := math.MaxInt64
		for c, count := range counts {
			if count < minCount {
				minCount = count
				minChar = c
			}
		}
		message += string(minChar)
	}

	log.Printf("the original message is %s\n", message)
}
