package main

import (
	"Advent-Of-Code-Go/2020/day20/part1/tile"
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 20: Jurassic Jigsaw ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/20
 */

var tiles []tile.ImageTile

func getAboveTile(i, size int, ordered []tile.ImageTile) (above tile.ImageTile, found bool) {
	found = i >= size
	if found {
		above = ordered[i-size]
	}
	return
}

func getLeftTile(i, size int, ordered []tile.ImageTile) (left tile.ImageTile, found bool) {
	found = i%size != 0
	if found {
		left = ordered[i-1]
	}
	return
}

func setOrderedTile(t tile.ImageTile, i, size int, ordered []tile.ImageTile) bool {
	above, hasAbove := getAboveTile(i, size, ordered)
	left, hasLeft := getLeftTile(i, size, ordered)
	if (!hasAbove || above.BottomEdge == t.TopEdge) &&
		(!hasLeft || left.RightEdge == t.LeftEdge) {
		ordered[i] = t
		return true
	}
	return false
}

func buildImage(i, size int, ordered []tile.ImageTile, usedIndexes map[int]struct{}) bool {
	// We're done when all tiles have been placed in a valid orientation
	if i == len(ordered) {
		return true
	}
	// Backtrack
	for j := 0; j < len(tiles); j++ {
		// If this tile hasn't been used yet...
		_, used := usedIndexes[j]
		if !used {
			tile := tiles[j]

			// Try all orientations of the tile by rotating and/or flipping it
			for q := 0; q < 8; q++ {
				if q == 4 {
					tile.Flip()
				}
				tile.Rot()

				// Try this tile at the current position in the image
				if setOrderedTile(tile, i, size, ordered) {
					usedIndexes[j] = struct{}{}
					if buildImage(i+1, size, ordered, usedIndexes) {
						return true
					}
					delete(usedIndexes, j)
				}
			}
		}
	}
	return false
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	tiles = make([]tile.ImageTile, 0, 0)
	{
		t := tile.ImageTile{}
		for scanner.Scan() {
			line := scanner.Text()

			if line == "" {
				tiles = append(tiles, t)

				t = tile.ImageTile{}

			} else if strings.HasPrefix(line, "Tile") {
				id, _ := strconv.Atoi(line[5 : len(line)-1])
				t.ID = id

			} else {
				if t.TopEdge == "" {
					t.TopEdge = line
				}
				t.LeftEdge += line[:1]
				t.RightEdge += line[len(line)-1:]
				t.BottomEdge = line
			}
		}
	}
	// log.Println(tiles)
	size := int(math.Sqrt(float64(len(tiles))))
	// log.Println("image size", size, "x", size)

	ordered := make([]tile.ImageTile, size*size, size*size)
	buildImage(0, size, ordered, make(map[int]struct{}))

	// Multiply the IDs of the 4 corner tiles together
	prod := ordered[0].ID                 // upper left
	prod *= ordered[size-1].ID            // upper right
	prod *= ordered[len(ordered)-size].ID // lower left
	prod *= ordered[len(ordered)-1].ID    // lower right

	log.Println("the product of the corner tile IDs is", prod)
}
