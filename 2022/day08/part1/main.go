package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	trees := make([][]int, 0)

	for _, line := range lines {
		row := make([]int, 0)
		for _, r := range line {
			height, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatalln(err)
			}
			row = append(row, height)
		}
		trees = append(trees, row)
	}
	nRows, nCols := len(trees), len(trees[0])

	var visibilityMaps [4][][]bool
	for dir := 0; dir < 4; dir++ {
		visibilityMaps[dir] = make([][]bool, nRows)
		for col := 0; col < nCols; col++ {
			visibilityMaps[dir][col] = make([]bool, nCols)
		}
	}

	var dir int
	// Look right
	for row := 0; row < nRows; row++ {
		maxHeight := -1
		for col := 0; col < nCols; col++ {
			treeHeight := trees[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibilityMaps[dir][row][col] = true
			}
		}
	}

	dir++
	// Look left
	for row := 0; row < nRows; row++ {
		maxHeight := -1
		for col := nCols - 1; col >= 0; col-- {
			treeHeight := trees[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibilityMaps[dir][row][col] = true
			}
		}
	}

	dir++
	// Look down
	for col := 0; col < nCols; col++ {
		maxHeight := -1
		for row := 0; row < nRows; row++ {
			treeHeight := trees[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibilityMaps[dir][row][col] = true
			}
		}
	}

	dir++
	// Look up
	for col := 0; col < nCols; col++ {
		maxHeight := -1
		for row := nRows - 1; row >= 0; row-- {
			treeHeight := trees[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibilityMaps[dir][row][col] = true
			}
		}
	}

	var count int
	for row := 0; row < nRows; row++ {
		for col := 0; col < nCols; col++ {
			var visible bool
			for dir := 0; dir < 4; dir++ {
				visible = visible || visibilityMaps[dir][row][col]
			}
			if visible {
				count++
			}
		}
	}
	log.Println("visible trees:", count)
}
