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

const expectedResult = 591365

/*
 *   --- Day 22: Reactor Reboot ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2021/day/22
 */

const limit = 50

var cuboids map[coord3D]struct{}

type coord3D struct {
	x, y, z int
}

type cuboid struct {
	x0, x1 int
	y0, y1 int
	z0, z1 int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nLitCubes := doPart1()
	log.Println("the number of on cubes is", nLitCubes)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cuboids = make(map[coord3D]struct{})

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
		x0 = clamp(x0, -limit, x0)
		x1 = clamp(x1, x1, limit)

		y0, _ := strconv.Atoi(nums[2])
		y1, _ := strconv.Atoi(nums[3])
		if y0 > y1 {
			y0, y1 = y1, y0
		}
		y0 = clamp(y0, -limit, y0)
		y1 = clamp(y1, y1, limit)

		z0, _ := strconv.Atoi(nums[4])
		z1, _ := strconv.Atoi(nums[5])
		if z0 > z1 {
			z0, z1 = z1, z0
		}
		z0 = clamp(z0, -limit, z0)
		z1 = clamp(z1, z1, limit)

		toggle(cuboid{x0, x1, y0, y1, z0, z1}, on)
	}

	return len(cuboids)
}

func toggle(c cuboid, on bool) {
	for z := c.z0; z <= c.z1; z++ {
		for y := c.y0; y <= c.y1; y++ {
			for x := c.x0; x <= c.x1; x++ {
				if on {
					cuboids[coord3D{x, y, z}] = struct{}{}
				} else {
					delete(cuboids, coord3D{x, y, z})
				}
			}
		}
	}
}

func clamp(x, lo, hi int) int {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}
