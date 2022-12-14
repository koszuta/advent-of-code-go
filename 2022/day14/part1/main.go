package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var SandSource = Vec2{500, 0}

const (
	Air = iota
	Rock
	Sand
)

type Vec2 struct {
	x, y int
}

var lines []string

func init() {
	b, _ := os.ReadFile("../input.txt")
	lines = strings.Split(string(b), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line) // sanitize CRLF
	}
}

func main() {
	cave := make(map[Vec2]int)
	var lowestY int

	// Add the rocks
	for _, line := range lines {
		var prev *Vec2
		for _, v := range strings.Split(line, " -> ") {
			parts := strings.Split(v, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			if y > lowestY {
				lowestY = y
			}
			if prev != nil {
				start, end := prev.x, x
				if start > end {
					start, end = end, start
				}
				for u := start; u <= end; u++ {
					cave[Vec2{u, y}] = Rock
				}
				start, end = prev.y, y
				if start > end {
					start, end = end, start
				}
				for v := start; v <= end; v++ {
					cave[Vec2{x, v}] = Rock
				}
			}
			prev = &Vec2{x, y}
		}
	}
	// PrintCave(cave)

OUT:
	for {
		sand := SandSource
		for {
			if sand.y == lowestY {
				break OUT
			}
			sand.y = sand.y + 1
			if _, found := cave[sand]; !found {
				continue
			}
			sand.x = sand.x - 1
			if _, found := cave[sand]; !found {
				continue
			}
			sand.x = sand.x + 2
			if _, found := cave[sand]; !found {
				continue
			}
			sand.x--
			sand.y--
			cave[sand] = Sand
			break
		}
	}
	PrintCave(cave)

	var count int
	for _, typ := range cave {
		if typ == Sand {
			count++
		}
	}
	log.Println("units of sand at rest:", count)
}

func PrintCave(cave map[Vec2]int) {
	minX, minY, maxX, maxY := SandSource.x, SandSource.y, SandSource.x, SandSource.y
	for pos := range cave {
		if pos.x < minX {
			minX = pos.x
		}
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	maxX, maxY = maxX+1, maxY+1

	grid := make([][]byte, maxY-minY)
	for y := range grid {
		grid[y] = make([]byte, maxX-minX)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	grid[SandSource.y-minY][SandSource.x-minX] = '+'
	for pos, typ := range cave {
		switch typ {
		case Rock:
			grid[pos.y-minY][pos.x-minX] = '#'
		case Sand:
			grid[pos.y-minY][pos.x-minX] = 'o'
		}
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}
