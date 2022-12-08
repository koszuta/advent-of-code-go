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

	var maxScore int
	for row := 1; row < nRows-1; row++ {
		for col := 1; col < nCols-1; col++ {
			var seeLeft, seeRight, seeUp, seeDown int
			for left := col - 1; left >= 0; left-- {
				seeLeft++
				if trees[row][left] >= trees[row][col] {
					break
				}
			}
			for right := col + 1; right < nCols; right++ {
				seeRight++
				if trees[row][right] >= trees[row][col] {
					break
				}
			}
			for up := row - 1; up >= 0; up-- {
				seeUp++
				if trees[up][col] >= trees[row][col] {
					break
				}
			}
			for down := row + 1; down < nRows; down++ {
				seeDown++
				if trees[down][col] >= trees[row][col] {
					break
				}
			}

			scenicScore := seeLeft * seeRight * seeUp * seeDown
			if scenicScore > maxScore {
				maxScore = scenicScore
			}
		}
	}
	log.Println("highest scenic score:", maxScore)
}
