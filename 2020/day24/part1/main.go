package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
 *   --- Day 24: Lobby Layout ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/24
 */

type hexTile struct {
	ns, ew int
}

func (t *hexTile) toString() string {
	return fmt.Sprintf("(%f,%f)", float64(t.ns)/2.0, float64(t.ew)/2.0)
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

	log.Println("after all of the instructions have been followed,", len(blackTiles), "tiles are left with the black side up")
}
