package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 1211172281877240

/*
 *   --- Day 22: Reactor Reboot ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/22#part2
 */

var cuboids map[cuboid]int

type cuboid struct {
	x0, x1 int
	y0, y1 int
	z0, z1 int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nLitCubes := doPart2()
	log.Println("the number of on cubes is", nLitCubes)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cuboids = make(map[cuboid]int)

	rebootRegex := regexp.MustCompile(`-?\d+`)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		on := parts[0] == "on"
		nums := rebootRegex.FindAllString(parts[1], -1)

		x0, _ := strconv.Atoi(nums[0])
		x1, _ := strconv.Atoi(nums[1])
		if x0 > x1 {
			x0, x1 = x1, x0
		}
		y0, _ := strconv.Atoi(nums[2])
		y1, _ := strconv.Atoi(nums[3])
		if y0 > y1 {
			y0, y1 = y1, y0
		}
		z0, _ := strconv.Atoi(nums[4])
		z1, _ := strconv.Atoi(nums[5])
		if z0 > z1 {
			z0, z1 = z1, z0
		}

		newCuboid := cuboid{x0, x1, y0, y1, z0, z1}

		change := make(map[cuboid]int)
		for c, sign := range cuboids {
			if intersection, intersects := intersection(c, newCuboid); intersects {
				change[intersection] -= sign
			}
		}
		if on {
			change[newCuboid]++
		}
		for c, sign := range change {
			cuboids[c] += sign
		}
	}

	nOnCubes := 0
	for c, sign := range cuboids {
		nOnCubes += (c.volume() * sign)
	}
	return nOnCubes
}

func intersection(c1, c2 cuboid) (intersection cuboid, intersects bool) {
	intersection = cuboid{
		max(c1.x0, c2.x0),
		min(c1.x1, c2.x1),
		max(c1.y0, c2.y0),
		min(c1.y1, c2.y1),
		max(c1.z0, c2.z0),
		min(c1.z1, c2.z1),
	}
	if intersection.x0 > intersection.x1 || intersection.y0 > intersection.y1 || intersection.z0 > intersection.z1 {
		return cuboid{}, false
	}
	return intersection, true
}

func (c *cuboid) volume() int {
	return (c.x1 - c.x0 + 1) * (c.y1 - c.y0 + 1) * (c.z1 - c.z0 + 1)
}

func min(p, q int) int {
	if p < q {
		return p
	}
	return q
}

func max(p, q int) int {
	if p > q {
		return p
	}
	return q
}
