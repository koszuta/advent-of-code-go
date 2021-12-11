package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

/*
 *   --- Day 11: Dumbo Octopus ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2021/day/11
 */

const (
	size   = 10
	nSteps = 100
)

var octopuses [][]int

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nFlashes := doPart1()
	log.Println("the number of total flashes after", nSteps, "steps is", nFlashes)
}

func doPart1() int {
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

	nFlashes := 0
	for i := 0; i < nSteps; i++ {
		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {
				octopuses[j][i]++
			}
		}
		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {
				nFlashes += flash(i, j)
			}
		}
	}
	return nFlashes
}

func flash(x, y int) (flashes int) {
	if octopuses[y][x] > 9 {
		octopuses[y][x] = 0
		flashes++

		doFlash := func(x, y int) {
			if x >= 0 && y >= 0 && y < len(octopuses) && x < len(octopuses[y]) && octopuses[y][x] != 0 {
				octopuses[y][x]++
				flashes += flash(x, y)
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
	return
}
