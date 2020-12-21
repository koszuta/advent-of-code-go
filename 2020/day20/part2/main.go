package main

import (
	"Advent-Of-Code-Go/2020/day20/part2/tile"
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 20: Jurassic Jigsaw ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/20#part2
 */

var tiles []tile.ImageTile

func getAbove(i, size int, ordered []tile.ImageTile) (above tile.ImageTile, found bool) {
	found = i >= size
	if found {
		above = ordered[i-size]
	}
	return
}

func getLeft(i, size int, ordered []tile.ImageTile) (left tile.ImageTile, found bool) {
	found = i%size != 0
	if found {
		left = ordered[i-1]
	}
	return
}

func setOrderedTile(t tile.ImageTile, i, size int, ordered []tile.ImageTile) bool {
	above, hasAbove := getAbove(i, size, ordered)
	left, hasLeft := getLeft(i, size, ordered)
	if (!hasAbove || above.BottomEdge == t.TopEdge) && (!hasLeft || left.RightEdge == t.LeftEdge) {
		ordered[i] = t
		return true
	}
	return false
}

func buildImage(idx, size int, ordered []tile.ImageTile, usedIndexes map[int]struct{}) bool {
	// We're done when all tiles have been placed in a valid orientation
	if idx == len(ordered) {
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
				if setOrderedTile(tile, idx, size, ordered) {
					usedIndexes[j] = struct{}{}
					if buildImage(idx+1, size, ordered, usedIndexes) {
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
				before := len(t.Image)
				t.Image = t.Image[:len(t.Image)-1]
				if len(t.Image) == before {
					log.Panicln("expected length", before-1)
				}
				tiles = append(tiles, t)

				t = tile.ImageTile{}

			} else if strings.HasPrefix(line, "Tile") {
				id, _ := strconv.Atoi(line[5 : len(line)-1])
				t.ID = id

			} else {
				if t.TopEdge == "" {
					t.TopEdge = line
				} else {
					t.Image = append(t.Image, line[1:len(line)-1])
				}
				t.LeftEdge += line[:1]
				t.RightEdge += line[len(line)-1:]
				t.BottomEdge = line
			}
		}
	}
	// log.Println(tiles)
	imageSize := int(math.Sqrt(float64(len(tiles))))
	// log.Println("image size", imageSize, "x", imageSize)

	ordered := make([]tile.ImageTile, imageSize*imageSize, imageSize*imageSize)
	buildImage(0, imageSize, ordered, make(map[int]struct{}))
	// log.Println(ordered[0].ID * ordered[imageSize-1].ID * ordered[len(ordered)-imageSize].ID * ordered[len(ordered)-1].ID)

	tileSize := len(tiles[0].Image[0])
	// log.Println("tile size", tileSize, "x", tileSize)
	fullImage := make([]string, imageSize*tileSize, imageSize*tileSize)
	for i, tile := range ordered {
		for q, line := range tile.Image {
			fullImage[i/imageSize*tileSize+q] += line
		}
	}

	// Look for sea monsters in the image
	//  ____________________
	// |                  # |
	// |#    ##    ##    ###|
	// | #  #  #  #  #  #   |
	// |____________________|

	line1Indexes := make([]int, 0, 0)
	line1Indexes = append(line1Indexes, 18)

	line2Indexes := make([]int, 0, 0)
	for _, i := range [...]int{0, 5, 6, 11, 12, 17, 18, 19} {
		line2Indexes = append(line2Indexes, i)
	}

	line3Indexes := make([]int, 0, 0)
	for _, i := range [...]int{1, 4, 7, 10, 13, 16} {
		line3Indexes = append(line3Indexes, i)
	}

	// Check each orientation of the image
	for q := 0; q < 8; q++ {
		if q == 4 {
			fullImage = tile.FlipStrings(fullImage)
		}
		fullImage = tile.RotStrings(fullImage)

		seaMonsterCount := 0
		for i := 0; i < len(fullImage)-2; i++ {
			for j := 0; j < len(fullImage[i])-20; j++ {
				line1, line2, line3 := []rune(fullImage[i]), []rune(fullImage[i+1]), []rune(fullImage[i+2])
				if couldHaveSeaMonster(j, line1, line1Indexes) &&
					couldHaveSeaMonster(j, line2, line2Indexes) &&
					couldHaveSeaMonster(j, line3, line3Indexes) {

					// Replace the sea monster '#' with 'O'
					fullImage[i] = drawSeaMonster(j, line1, line1Indexes)
					fullImage[i+1] = drawSeaMonster(j, line2, line2Indexes)
					fullImage[i+2] = drawSeaMonster(j, line3, line3Indexes)

					seaMonsterCount++
				}
			}
		}

		if seaMonsterCount > 0 {
			hashCount := 0
			for _, line := range fullImage {
				for _, c := range line {
					if c == '#' {
						hashCount++
					}
				}
			}
			tile.PrintStrings(fullImage)
			log.Println(seaMonsterCount, "sea monsters found")
			log.Println("the image has", hashCount, "'#' that are not part of a sea monster")
			break
		}
	}
}

func drawSeaMonster(start int, line []rune, indexes []int) string {
	for _, i := range indexes {
		line[start+i] = 'O'
	}
	return string(line)
}

func couldHaveSeaMonster(start int, line []rune, indexes []int) bool {
	for _, i := range indexes {
		if line[start+i] != '#' {
			return false
		}
	}
	return true
}
