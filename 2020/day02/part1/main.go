package main

import (
	"bufio"
	"log"
	"os"
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

	goodPasswords := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Get the password
		parts := strings.Split(line, ": ")
		password := parts[1]

		// Get the letter
		parts = strings.Split(parts[0], " ")
		letter := parts[1]

		// Get the min/max count
		parts = strings.Split(parts[0], "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		// Check if the letter is within the valid range
		count := strings.Count(password, letter)
		if count >= min && count <= max {
			goodPasswords++
		}
	}

	log.Printf("the number of valid passwords is %d\n", goodPasswords)
}
