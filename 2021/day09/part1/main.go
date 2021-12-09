package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

/*
 *   --- Day 9: Smoke Basin ---
 *        --- Part One ---
 *
 *   https://adventofcode.com/2021/day/9
 */

var heatMap [][]int

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	sum := doPart1()
	log.Println("the sum of the risk levels of all low points on the heightmap is", sum)
}

func doPart1() int {
	heatMap = make([][]int, 0)

	file, _ := os.Open("../input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for j := 0; scanner.Scan(); j++ {
		heatMap = append(heatMap, make([]int, 0))
		for _, r := range scanner.Text() {
			val, _ := strconv.Atoi(string(r))
			heatMap[j] = append(heatMap[j], val)
		}
	}

	sum := 0
	for y, row := range heatMap {
		for x, val := range row {
			if isLow(x, y) {
				sum += val + 1
			}
		}
	}
	return sum
}

func isLow(x, y int) bool {
	val := heatMap[y][x]
	if x > 0 && heatMap[y][x-1] <= val {
		return false
	}
	if x < len(heatMap[y])-1 && heatMap[y][x+1] <= val {
		return false
	}
	if y > 0 && heatMap[y-1][x] <= val {
		return false
	}
	if y < len(heatMap)-1 && heatMap[y+1][x] <= val {
		return false
	}
	return true
}
