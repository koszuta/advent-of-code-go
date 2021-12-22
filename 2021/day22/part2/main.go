package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 0

/*
 *   --- Day 22:  ---
 *      --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/22#part2
 */

var (
	onCuboids  []cuboid
	offCuboids []cuboid
)

type cuboid struct {
	x0, x1 int
	y0, y1 int
	z0, z1 int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	x := doPart2()
	log.Println(x)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		on := parts[0] == "on"

		parts = strings.Split(parts[1], ",")

		xParts := strings.Split(parts[0][2:], "..")
		x0, _ := strconv.Atoi(xParts[0])
		x1, _ := strconv.Atoi(xParts[1])
		if x0 > x1 {
			x0, x1 = x1, x0
		}

		yParts := strings.Split(parts[1][2:], "..")
		y0, _ := strconv.Atoi(yParts[0])
		y1, _ := strconv.Atoi(yParts[1])
		if y0 > y1 {
			y0, y1 = y1, y0
		}

		zParts := strings.Split(parts[2][2:], "..")
		z0, _ := strconv.Atoi(zParts[0])
		z1, _ := strconv.Atoi(zParts[1])
		if z0 > z1 {
			z0, z1 = z1, z0
		}

		if on {
			onCuboids = append(onCuboids, cuboid{x0, x1, y0, y1, z0, z1})
		} else {
			offCuboids = append(offCuboids, cuboid{x0, x1, y0, y1, z0, z1})
		}
	}

	nOnCubes := 0

	for j := 1; j < len(onCuboids); j++ {
		for i := j + 1; i < len(onCuboids); i++ {

		}
	}

	return nOnCubes
}

func overlap(c1, c2 cuboid) []cuboid {
    newCuboids = 
}

func volume(c cuboid) int {
	return c.x0*c.x1 + c.y0*c.y1 + c.z0*c.z1
}
