package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
)

/*
 *   --- Day 11: Seating System ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2020/day/11
 */

var rows, cols int
var spaces []rune

func occupiedSeen(r, c int) int {
	if r >= 0 && r < rows && c >= 0 && c < cols && spaces[r*cols+c] == '#' {
		return 1
	}
	return 0
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// Init the list of spaces, # rows, and # cols
	spaces = make([]rune, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		for _, r := range []rune(line) {
			spaces = append(spaces, r)
		}
	}
	rows = len(spaces) / cols

	buf := make([]rune, len(spaces), len(spaces))
	for {
		// Perform a simultaneous update of the spaces by storing to the buffer and swapping
		for i := 0; i < len(spaces); i++ {
			space := spaces[i]
			row := i / cols
			col := i % cols

			// Count the occupied seats directly adjacent to this one
			occupiedCount := 0
			occupiedCount += occupiedSeen(row, col-1)   // left
			occupiedCount += occupiedSeen(row, col+1)   // right
			occupiedCount += occupiedSeen(row-1, col)   // up
			occupiedCount += occupiedSeen(row+1, col)   // down
			occupiedCount += occupiedSeen(row-1, col-1) // up and left
			occupiedCount += occupiedSeen(row-1, col+1) // up and right
			occupiedCount += occupiedSeen(row+1, col-1) // down and left
			occupiedCount += occupiedSeen(row+1, col+1) // down and right

			// If an empty seat sees no occupied seats, it becomes occupied
			// If an occupied seat sees 4 or more occupied seats, it becomes empty
			// Otherwise, the space doesn't change
			buf[i] = space
			switch space {
			case 'L':
				if occupiedCount == 0 {
					buf[i] = '#'
				}
			case '#':
				if occupiedCount >= 4 {
					buf[i] = 'L'
				}
			}
		}

		// Stop if no seats change
		if reflect.DeepEqual(spaces, buf) {
			break
		}

		// Swap buffers
		spaces, buf = buf, spaces
	}

	occupiedCount := 0
	for _, r := range spaces {
		if r == '#' {
			occupiedCount++
		}
	}
	log.Printf("%d seats end up occupied\n", occupiedCount)
}
