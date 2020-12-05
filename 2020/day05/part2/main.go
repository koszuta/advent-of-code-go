package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

/*
 *   --- Day 5: Binary Boarding ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/5#part2
 */

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	seatIDs := make([]int, 0, 0)
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
		seatIDs = append(seatIDs, seatID)
	}

	// Seat IDs should be sequential
	// Therefore, we can compare the IDs of adjacent seats
	// and find where there's a gap
	sort.Ints(seatIDs)
	prevSeatID, i := -1, 0
	for prevSeatID < 0 || seatIDs[i]-prevSeatID == 1 {
		prevSeatID = seatIDs[i]
		i++
	}

	log.Printf("your seat ID is %d\n", prevSeatID+1)
}
