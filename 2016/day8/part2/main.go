package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

/*
 *   --- Day 8: Two-Factor Authentication ---
 *               --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/8#part2
 */

const width = 50
const height = 6

var screen [height][width]bool

var rectRex, rotRowRex, rotColRex *regexp.Regexp

func init() {
	rectRex = regexp.MustCompile(`rect (\d+)x(\d+)`)
	rotRowRex = regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
	rotColRex = regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		rect := rectRex.FindAllStringSubmatch(line, -1)
		rotRow := rotRowRex.FindAllStringSubmatch(line, -1)
		rotCol := rotColRex.FindAllStringSubmatch(line, -1)

		if rect != nil {
			a, _ := strconv.Atoi(rect[0][1])
			b, _ := strconv.Atoi(rect[0][2])
			rectOp(a, b)
		}
		if rotRow != nil {
			a, _ := strconv.Atoi(rotRow[0][1])
			b, _ := strconv.Atoi(rotRow[0][2])
			rotRowOp(a, b)
		}
		if rotCol != nil {
			a, _ := strconv.Atoi(rotCol[0][1])
			b, _ := strconv.Atoi(rotCol[0][2])
			rotColOp(a, b)
		}
		// printScreen()
	}

	nLitPixels := 0
	for _, row := range screen {
		for _, pixel := range row {
			if pixel {
				nLitPixels++
			}
		}
	}

	printScreen()
}

func printScreen() {
	for _, row := range screen {
		for _, pixel := range row {
			if pixel {
				fmt.Print("\u2588")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func rectOp(a, b int) {
	for x := 0; x < a; x++ {
		for y := 0; y < b; y++ {
			screen[y][x] = true
		}
	}
}

func rotRowOp(a, b int) {
	row := make([]bool, 0, width)
	row = append(row, screen[a][width-b:]...)
	row = append(row, screen[a][:width-b]...)
	for x, v := range row {
		screen[a][x] = v
	}
}

func rotColOp(a, b int) {
	col := make([]bool, height)
	for y, row := range screen {
		col[y] = row[a]
	}
	for y := 0; y < height; y++ {
		screen[(y+b)%height][a] = col[y]
	}
}
