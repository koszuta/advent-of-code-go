package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

const expectedResult = 418

/*
 *   --- Day 11: Dumbo Octopus ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/11#part2
 */

const size = 10

var octopuses [][]int

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nSteps := doPart2()
	log.Println("the first step during which all octopuses flash is", nSteps)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	octopuses = make([][]int, 0, size)
	for scanner.Scan() {
		row := make([]int, 0, size)
		for _, r := range scanner.Text() {
			energy, _ := strconv.Atoi(string(r))
			row = append(row, energy)
		}
		octopuses = append(octopuses, row)
	}

	var nSteps int
	for nSteps = 0; !allFlashed(); nSteps++ {
		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {
				octopuses[j][i]++
			}
		}
		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {
				flash(i, j)
			}
		}
	}
	return nSteps
}

func allFlashed() bool {
	for _, row := range octopuses {
		for _, energy := range row {
			if energy != 0 {
				return false
			}
		}
	}
	return true
}

func flash(x, y int) {
	if octopuses[y][x] > 9 {
		octopuses[y][x] = 0

		doFlash := func(x, y int) {
			if x >= 0 && y >= 0 && y < size && x < size && octopuses[y][x] != 0 {
				octopuses[y][x]++
				flash(x, y)
			}
		}

		doFlash(x-1, y)   // left
		doFlash(x, y-1)   // up
		doFlash(x+1, y)   // right
		doFlash(x, y+1)   // down
		doFlash(x-1, y-1) // upper left
		doFlash(x-1, y+1) // lower left
		doFlash(x+1, y-1) // upper right
		doFlash(x+1, y+1) // lower right
	}
}
