package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 17: Conway Cubes ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/17
 */

const cycles = 6

var activeCubes map[coord3D]struct{}

type coord3D struct {
	x, y, z int
}

func getNeighbors(cube coord3D) []coord3D {
	neighbors := make([]coord3D, 0, 0)
	for z := cube.z + 1; z >= cube.z-1; z-- {
		for y := cube.y + 1; y >= cube.y-1; y-- {
			for x := cube.x + 1; x >= cube.x-1; x-- {
				neighbor := coord3D{x, y, z}
				if neighbor != cube {
					neighbors = append(neighbors, neighbor)
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

	activeCubes = make(map[coord3D]struct{})
	for y := 0; scanner.Scan(); y++ {
		x := 0
		for _, c := range scanner.Text() {
			if c == '#' {
				cube := coord3D{x, y, 0}
				activeCubes[cube] = struct{}{}
			}
			x++
		}
	}

	for i := 0; i < cycles; i++ {
		activeNeighborCounts := make(map[coord3D]int)
		for cube := range activeCubes {
			for _, neighbor := range getNeighbors(cube) {
				count, exists := activeNeighborCounts[neighbor]
				if !exists {
					count = 0
				}
				activeNeighborCounts[neighbor] = count + 1
			}
		}

		activeBuf := make(map[coord3D]struct{})
		for cube, count := range activeNeighborCounts {
			_, active := activeCubes[cube]
			if active && (count == 2 || count == 3) {
				activeBuf[cube] = struct{}{}
			}
			if !active && count == 3 {
				activeBuf[cube] = struct{}{}
			}
		}
		activeCubes = activeBuf
	}

	log.Println("the number of active cubes after the sixth cycle is", len(activeCubes))
}
