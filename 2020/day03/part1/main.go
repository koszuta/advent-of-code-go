package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 3: Toboggan Trajectory ---
 *            --- Part One ---
 *
 *   https://adventofcode.com/2020/day/3
 */

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	col, treeCount := 0, 0
	for scanner.Scan() {
		// Rows consist of open squares (.) and trees (#)
		rowSquares := []rune(scanner.Text())

		// Check if you encounter a tree
		if rowSquares[col] == '#' {
			treeCount++
		}

		// Move right 3 cols
		// The tree pattern repeats once you're past the edge of the puzzle input
		col = (col + 3) % len(rowSquares)
	}

	log.Printf("%d trees encountered\n", treeCount)
}
