package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
 *   --- Day 24: Lobby Layout ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/24#part2
 */

const days = 100

type hexTile struct {
	ns, ew int
}

func (t *hexTile) toString() string {
	return fmt.Sprintf("(%f,%f)", float64(t.ns)/2.0, float64(t.ew)/2.0)
}

func getAdjacentTiles(t hexTile) (adjacent []hexTile) {
	adjacent = append(adjacent,
		hexTile{t.ns, t.ew + 2},     // E
		hexTile{t.ns, t.ew - 2},     // W
		hexTile{t.ns + 1, t.ew + 1}, // NE
		hexTile{t.ns + 1, t.ew - 1}, // NW
		hexTile{t.ns - 1, t.ew + 1}, // SE
		hexTile{t.ns - 1, t.ew - 1}, // SW
	)
	return
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	blackTiles := make(map[hexTile]struct{})
	for scanner.Scan() {
		line := scanner.Text()

		ns, ew := 0, 0
		for i := 0; i < len(line); i++ {
			isDiagonal := false
			switch line[i] {
			case 'e':
				ew += 2
			case 'w':
				ew -= 2
			case 'n':
				ns++
				isDiagonal = true
			case 's':
				ns--
				isDiagonal = true
			}
			if isDiagonal {
				i++
				switch line[i] {
				case 'e':
					ew++
				case 'w':
					ew--
				}
			}
		}

		tile := hexTile{ns, ew}
		// log.Println(tile.toString())

		_, exists := blackTiles[tile]
		if exists {
			delete(blackTiles, tile)
		} else {
			blackTiles[tile] = struct{}{}
		}
	}

	for day := 0; day < days; day++ {
		adjacentBlackCounts := make(map[hexTile]int)
		for t := range blackTiles {
			for _, adjacent := range getAdjacentTiles(t) {
				count, exists := adjacentBlackCounts[adjacent]
				if !exists {
					count = 0
				}
				adjacentBlackCounts[adjacent] = count + 1
			}
		}

		blackBuf := make(map[hexTile]struct{})
		for t, count := range adjacentBlackCounts {
			_, isBlack := blackTiles[t]
			if isBlack && (count == 1 || count == 2) {
				blackBuf[t] = struct{}{}
			}
			if !isBlack && count == 2 {
				blackBuf[t] = struct{}{}
			}
		}
		blackTiles = blackBuf
		// log.Printf("Day %d: %d\n", day+1, len(blackTiles))
	}

	log.Println("after", days, "days,", len(blackTiles), "tiles are black")
}
