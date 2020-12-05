package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 5: Binary Boarding ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2020/day/5
 */

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	maxSeatID := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())

		// Binary search the rows and cols to find the seat
		var row, col int

		// F(ront) -> lower half
		// B(ack)  -> upper half
		l, h := 0, 127
		for i := 0; i < 7; i++ {
			switch line[i] {
			case 'F':
				h = (l + h) / 2
			case 'B':
				l = (l + h + 1) / 2
			}
		}
		// Actual row is the lower bound
		row = l

		// L(eft)  -> lower half
		// R(ight) -> upper half
		l, h = 0, 7
		for i := 7; i < 10; i++ {
			switch line[i] {
			case 'L':
				h = (l + h) / 2
			case 'R':
				l = (l + h + 1) / 2
			}
		}
		// Actual column is the upper bound
		col = h

		// The Seat ID is 8(row) + col
		seatID := row*8 + col

		// We're looking for the highest seat ID
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	log.Printf("the highest seat ID is %d\n", maxSeatID)
}
