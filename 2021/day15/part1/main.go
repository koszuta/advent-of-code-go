package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const expectedResult = 656

/*
 *   --- Day 15: Chiton ---
 *      --- Part One ---
 *
 *   https://adventofcode.com/2021/day/15
 */

type coord2D struct {
	x, y int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	totalRisk := doPart1()
	log.Println("the lowest total risk of any path from the top left to the bottom right", totalRisk)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var size int
	cave := make(map[coord2D]int)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		size = len(line)
		for j, r := range line {
			riskLevel, _ := strconv.Atoi(string(r))
			cave[coord2D{i, j}] = riskLevel
		}
	}

	source := coord2D{0, 0}
	target := coord2D{size - 1, size - 1}

	unvisited := make(map[coord2D]struct{})
	risk := make(map[coord2D]int)
	prev := make(map[coord2D]coord2D)
	for v := range cave {
		risk[v] = math.MaxInt64
		prev[v] = coord2D{-1, -1}
		unvisited[v] = struct{}{}
	}
	risk[source] = 0

	for len(unvisited) > 0 {
		minRisk := math.MaxInt64
		var u coord2D
		for v := range unvisited {
			if risk[v] < minRisk {
				minRisk = risk[v]
				u = v
			}
		}

		delete(unvisited, u)

		doNeighbor := func(v coord2D) {
			_, ok := unvisited[v]
			if ok {
				alt := risk[u] + cave[v]
				if alt < risk[v] {
					risk[v] = alt
					prev[v] = u
				}
			}
		}

		doNeighbor(coord2D{u.x - 1, u.y}) // left
		doNeighbor(coord2D{u.x, u.y - 1}) // up
		doNeighbor(coord2D{u.x + 1, u.y}) // right
		doNeighbor(coord2D{u.x, u.y + 1}) // down
	}

	return risk[target]
}
