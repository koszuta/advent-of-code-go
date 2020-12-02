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
 *   https://adventofcode.com/2020/day/2#part2
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
		letter := []rune(parts[1])[0]

		// Get the positions
		parts = strings.Split(parts[0], "-")
		idx1, _ := strconv.Atoi(parts[0])
		idx2, _ := strconv.Atoi(parts[1])

		// Check if the letter shows up at exactly one of the positions
		p1, p2 := 0, 0
		if []rune(password)[idx1-1] == letter {
			p1 = 1
		}
		if []rune(password)[idx2-1] == letter {
			p2 = 1
		}
		if p1^p2 == 1 {
			goodPasswords++
		}
	}

	log.Printf("the number of valid passwords is %d\n", goodPasswords)
}
