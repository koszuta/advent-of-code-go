package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

/*
 *   --- Day 9: Smoke Basin ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/9#part2
 */

var heatMap [][]int

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	prod := doPart2()
	log.Println("the sizes of the three largest basins multiplied together is", prod)
}

func doPart2() int {
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

	basins := make([]int, 0)
	for y, row := range heatMap {
		for x := range row {
			size := findBasin(x, y)
			if size > 0 {
				basins = append(basins, size)
			}
		}
	}
	sort.Slice(basins, func(i, j int) bool {
		return basins[j] < basins[i] // reverse order
	})
	return basins[0] * basins[1] * basins[2]
}

func findBasin(x, y int) int {
	return doFindBasin(x, y, 0)
}

func doFindBasin(x, y int, basin int) int {
	if heatMap[y][x] == 9 {
		return basin
	}

	basin++
	heatMap[y][x] = 9

	if x > 0 {
		basin = doFindBasin(x-1, y, basin)
	}
	if x < len(heatMap[y])-1 {
		basin = doFindBasin(x+1, y, basin)
	}
	if y > 0 {
		basin = doFindBasin(x, y-1, basin)
	}
	if y < len(heatMap)-1 {
		basin = doFindBasin(x, y+1, basin)
	}
	return basin
}
