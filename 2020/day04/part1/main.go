package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
 *   --- Day 4: Passport Processing ---
 *            --- Part One ---
 *
 *   https://adventofcode.com/2020/day/4
 */

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	validPassports := 0
	fields := make(map[string]string)
	for scanner.Scan() {
		text := scanner.Text()

		// Passport fields can span multiple lines
		// Individual passports are separated by an empty line
		if text == "" {
			// Check if the 7 required fields are present
			_, hasCID := fields["cid"]
			if len(fields) == 8 || (!hasCID && len(fields) == 7) {
				validPassports++
			}

			// Reset the fields for the next passport
			fields = make(map[string]string)

		} else {
			// Fields are separated by spaces
			parts := strings.Split(text, " ")

			// Parse the fields as key:value pairs
			for _, part := range parts {
				field := strings.Split(part, ":")
				fields[field[0]] = field[1]
			}
		}
	}

	log.Printf("%d passports are valid\n", validPassports)
}
