package main

import (
	"log"
	"math/bits"
)

/*
 *   --- Day 13: A Maze of Twisty Little Cubicles ---
 *                   --- Part One ---
 *
 *   https://adventofcode.com/2016/day/13
 */

const (
	startX         = 1
	startY         = 1
	destX          = 31
	destY          = 39
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
	distances := move(1, 1, 0, make([]int, 0))
	if len(distances) == 0 {
		log.Panicf("couldn't reach %d,%d\n", destX, destY)
	}

	minDist := distances[0]
	for _, distance := range distances {
		if distance < minDist {
			minDist = distance
		}
	}
	log.Printf("the fewest number of steps required for you to reach %d,%d is %d\n", destX, destY, minDist)
}

func move(x, y, steps int, distances []int) []int {
	if x == destX && y == destY {
		return append(distances, steps)
	}
	visited[vec2i{x, y}] = steps

	distances = doMove(x-1, y, steps, distances) // left
	distances = doMove(x+1, y, steps, distances) // right
	distances = doMove(x, y-1, steps, distances) // up
	distances = doMove(x, y+1, steps, distances) // down
	return distances
}

func doMove(x, y, steps int, distances []int) []int {
	prevSteps, found := visited[vec2i{x, y}]
	if (!found || prevSteps > steps) && spaceIsOpen(x, y) {
		return move(x, y, steps+1, distances)
	}
	return distances
}

func spaceIsOpen(x, y int) bool {
	q := uint(x*x + 3*x + 2*x*y + y + y*y + favoriteNumber)
	return x >= 0 && y >= 0 && bits.OnesCount(q)%2 == 0
}
