package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 3: Toboggan Trajectory ---
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/3#part2
 */

func doToboggan(x, y int, file *os.File) int {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	row, col, treeCount := -1, 0, 0
	for scanner.Scan() {
		// Go down y rows
		row++
		if row%y != 0 {
			continue
		}

		// Rows consist of open squares (.) and trees (#)
		rowSquares := []rune(scanner.Text())

		// Check if you encounter a tree
		if rowSquares[col] == '#' {
			treeCount++
		}

		// Move right x cols
		// The tree pattern repeats once you're past the edge of the puzzle input
		col = (col + x) % len(rowSquares)
	}

	log.Printf("right %d, down %d: %d trees encountered\n", x, y, treeCount)
	return treeCount
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	defer file.Close()

	prod := doToboggan(1, 1, file) // Right 1, down 1
	prod *= doToboggan(3, 1, file) // Right 3, down 1
	prod *= doToboggan(5, 1, file) // Right 5, down 1
	prod *= doToboggan(7, 1, file) // Right 7, down 1
	prod *= doToboggan(1, 2, file) // Right 1, down 2

	log.Printf("puzzle answer: %d\n", prod)
}
