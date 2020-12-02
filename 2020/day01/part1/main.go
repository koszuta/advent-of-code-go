package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
 *   --- Day 1: Report Repair ---
 *         --- Part One ---
 *
 * Find 2 numbers from the puzzle input that sum to 2020.
 * What's the product of those 2 numbers?
 *
 * https://adventofcode.com/2020/day/1
 */

const currentYear = 2020

func main() {
	// Open the puzzle input file
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	// Keep track of the values in a map to quickly check the difference
	vals := make(map[int]struct{})

	// Go through the puzzle input line-by-line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert the puzzle input line to an int
		var val int
		{
			line := scanner.Text()
			val, err = strconv.Atoi(line)
			if err != nil {
				log.Panicln(err)
			}
		}

		// Check if the difference of 2020-val has already been seen
		_, exists := vals[currentYear-val]
		if exists {
			// Print the product of the two numbers that sum to 2020
			log.Printf("%d\n", val*(currentYear-val))
			break
		}

		// Store the value we've seen
		vals[val] = struct{}{}
	}
}
