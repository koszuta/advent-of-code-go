package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 6: Signals and Noise ---
 *           --- Part One ---
 *
 *   https://adventofcode.com/2016/day/6
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
		var maxChar rune
		maxCount := 0
		for c, count := range counts {
			if count > maxCount {
				maxCount = count
				maxChar = c
			}
		}
		message += string(maxChar)
	}

	log.Printf("the error-corrected version of the message is %s\n", message)
}
