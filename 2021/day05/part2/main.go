package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 21577

/*
 *   --- Day 5: Hydrothermal Venture ---
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/5#part2
 */

type line2D struct {
	start, end coord2D
}

type coord2D struct {
	x, y int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nOverlaps := doPart2()
	log.Println("the number of points where at least two lines overlap is", nOverlaps)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	lines := make([]line2D, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " -> ")
		line := line2D{
			start: parseCoord(parts[0]),
			end:   parseCoord(parts[1]),
		}
		lines = append(lines, line)
	}

	points := make(map[coord2D]int)
	for _, line := range lines {
		start, end := line.start, line.end
		if start.x > end.x {
			start, end = end, start
		}
		if start.y == end.y { // horizontal
			for x := start.x; x <= end.x; x++ {
				p := coord2D{x: x, y: start.y}
				points[p]++
			}
		} else if start.x == end.x { // vertical
			if start.y > end.y {
				start, end = end, start
			}
			for y := start.y; y <= end.y; y++ {
				p := coord2D{x: start.x, y: y}
				points[p]++
			}
		} else { // diagonal
			if start.y < end.y { // positive slope
				for x, y := start.x, start.y; x <= end.x; x, y = x+1, y+1 {
					p := coord2D{x: x, y: y}
					points[p]++
				}
			} else { // negative slope
				for x, y := start.x, start.y; x <= end.x; x, y = x+1, y-1 {
					p := coord2D{x: x, y: y}
					points[p]++
				}
			}
		}
	}

	nOverlaps := 0
	for _, count := range points {
		if count > 1 {
			nOverlaps++
		}
	}
	return nOverlaps
}

func parseCoord(s string) coord2D {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coord2D{x: x, y: y}
}
