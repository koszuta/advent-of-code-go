package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 17: Conway Cubes ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/17#part2
 */

const cycles = 6

var activeCubes map[coord4D]struct{}

type coord4D struct {
	w, x, y, z int
}

func getNeighbors(cube coord4D) []coord4D {
	neighbors := make([]coord4D, 0, 0)
	for z := cube.z + 1; z >= (cube.z - 1); z-- {
		for y := cube.y + 1; y >= (cube.y - 1); y-- {
			for x := cube.x + 1; x >= (cube.x - 1); x-- {
				for w := cube.w + 1; w >= (cube.w - 1); w-- {
					neighbor := coord4D{w, x, y, z}
					if neighbor != cube {
						neighbors = append(neighbors, neighbor)
					}
				}
			}
		}
	}
	return neighbors
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	activeCubes = make(map[coord4D]struct{})
	for y := 0; scanner.Scan(); y++ {
		x := 0
		for _, c := range scanner.Text() {
			if c == '#' {
				cube := coord4D{0, x, y, 0}
				activeCubes[cube] = struct{}{}
			}
			x++
		}
	}

	for i := 0; i < cycles; i++ {
		// start := time.Now()
		activeNeighborCounts := make(map[coord4D]int)
		for cube := range activeCubes {
			for _, neighbor := range getNeighbors(cube) {
				count, exists := activeNeighborCounts[neighbor]
				if !exists {
					count = 0
				}
				activeNeighborCounts[neighbor] = count + 1
			}
		}

		activeBuf := make(map[coord4D]struct{})
		for cube, count := range activeNeighborCounts {
			if count == 2 {
				_, active := activeCubes[cube]
				if active {
					activeBuf[cube] = struct{}{}
				}
			}
			if count == 3 {
				activeBuf[cube] = struct{}{}
			}
		}
		activeCubes = activeBuf
		// log.Println("update", i+1, "took", time.Since(start))
	}

	log.Println("the number of active hypercubes after the sixth cycle is", len(activeCubes))
}
