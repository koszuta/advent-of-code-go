package main

import "log"

/*
 *   --- Day 18: Like a Rogue ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/18#part2
 */

const (
	firstRow = ".^.^..^......^^^^^...^^^...^...^....^^.^...^.^^^^....^...^^.^^^...^^^^.^^.^.^^..^.^^^..^^^^^^.^^^..^"
	nRows    = 400_000
)

func main() {
	log.Println(firstRow)
	nSafeTiles := countSafeTiles(firstRow)

	prevRow := firstRow
	for row := 1; row < nRows; row++ {
		thisRow := make([]byte, len(firstRow))
		for col := 0; col < len(firstRow); col++ {
			left, center, right := byte('.'), prevRow[col], byte('.')
			if col != 0 {
				left = prevRow[col-1]
			}
			if col != len(firstRow)-1 {
				right = prevRow[col+1]
			}
			leftAndCenter := left == '^' && center == '^' && right == '.'
			centerAndRight := left == '.' && center == '^' && right == '^'
			onlyLeft := left == '^' && center == '.' && right == '.'
			onlyRight := left == '.' && center == '.' && right == '^'
			if leftAndCenter || centerAndRight || onlyLeft || onlyRight {
				thisRow[col] = '^'
			} else {
				thisRow[col] = '.'
			}
		}

		prevRow = string(thisRow)
		nSafeTiles += countSafeTiles(prevRow)
	}
	log.Printf("%d rows...\n", nRows-2)
	log.Println(prevRow)

	log.Printf("in %d rows, there are %d safe tiles\n", nRows, nSafeTiles)
}

func countSafeTiles(row string) int {
	nSafeTiles := 0
	for _, b := range row {
		if b == '.' {
			nSafeTiles++
		}
	}
	return nSafeTiles
}
