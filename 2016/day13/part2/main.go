package main

import (
	"log"
	"math/bits"
)

/*
 *   --- Day 13: A Maze of Twisty Little Cubicles ---
 *                   --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/13#part2
 */

const (
	maxSteps       = 50
	startX         = 1
	startY         = 1
	favoriteNumber = 1358
)

type vec2i struct {
	x, y int
}

var visited map[vec2i]int

func init() {
	visited = make(map[vec2i]int)
}

func main() {
	move(startX, startY, 0)
	count := 0
	for _, steps := range visited {
		if steps <= maxSteps {
			count++
		}
	}
	log.Printf("the number of locations you can reach in at most %d steps is %d\n", maxSteps, count)
}

func move(x, y, steps int) {
	visited[vec2i{x, y}] = steps
	if steps < maxSteps && x <= startX+maxSteps && y <= startY+maxSteps {
		doMove(x-1, y, steps) // left
		doMove(x+1, y, steps) // right
		doMove(x, y-1, steps) // up
		doMove(x, y+1, steps) // down
	}
}

func doMove(x, y, steps int) {
	prevSteps, found := visited[vec2i{x, y}]
	if (!found || prevSteps > steps) && spaceIsOpen(x, y) {
		move(x, y, steps+1)
	}
}

func spaceIsOpen(x, y int) bool {
	q := uint(x*x + 3*x + 2*x*y + y + y*y + favoriteNumber)
	return x >= 0 && y >= 0 && bits.OnesCount(q)%2 == 0
}
