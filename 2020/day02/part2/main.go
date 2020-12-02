package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

/*
 *   --- Day 2: Password Philosophy ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/2#part2
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

		// Get the 1st position, 2nd position, letter, and password from the line
		idx1, _ := strconv.Atoi(parts[0])
		idx2, _ := strconv.Atoi(parts[1])
		letter := []rune(parts[2])[0]
		password := parts[3]

		// Check if the letter shows up at exactly one of the positions
		b1 := []rune(password)[idx1-1] == letter
		b2 := []rune(password)[idx2-1] == letter
		if b1 != b2 {
			goodPasswords++
		}
	}

	log.Printf("the number of valid passwords is %d\n", goodPasswords)
}
