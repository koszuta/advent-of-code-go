package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 *   --- Day 2: Password Philosophy ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/2
 */

func main() {
	// Open the puzzle input file
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	goodPasswords := 0
	for scanner.Scan() {
		// Split the line on non-word chars
		re := regexp.MustCompile("\\W+")
		parts := re.Split(scanner.Text(), 4)

		// Get the min, max, letter, and password from the line
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		letter := parts[2]
		password := parts[3]

		// Check if the letter is within the valid range
		count := strings.Count(password, letter)
		if count >= min && count <= max {
			goodPasswords++
		}
	}

	log.Printf("the number of valid passwords is %d\n", goodPasswords)
}
