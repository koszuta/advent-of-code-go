package main

import (
	"bufio"
	"log"
	"os"
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

var (
	cuboids map[coord3D]struct{}
	steps   []rebootStep
)

type coord3D struct {
	x, y, z int
}

type cuboid struct {
	x0, x1 int
	y0, y1 int
	z0, z1 int
}

type rebootStep struct {
	c  cuboid
	on bool
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	x := doPart1()
	log.Println(x)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cuboids = make(map[coord3D]struct{})
	steps = make([]rebootStep, 0)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		on := parts[0] == "on"

		parts = strings.Split(parts[1], ",")

		xParts := strings.Split(parts[0][2:], "..")
		x0, _ := strconv.Atoi(xParts[0])
		x1, _ := strconv.Atoi(xParts[1])

		yParts := strings.Split(parts[1][2:], "..")
		y0, _ := strconv.Atoi(yParts[0])
		y1, _ := strconv.Atoi(yParts[1])

		zParts := strings.Split(parts[2][2:], "..")
		z0, _ := strconv.Atoi(zParts[0])
		z1, _ := strconv.Atoi(zParts[1])

		step := rebootStep{cuboid{x0, x1, y0, y1, z0, z1}, on}
		steps = append(steps, step)
	}

	for _, step := range steps {
		if step.c.x0 >= -limit && step.c.x1 <= limit && step.c.y0 >= -limit && step.c.y1 <= limit && step.c.z0 >= -limit && step.c.z1 <= limit {
			if step.on {
				turnOn(step.c)
			} else {
				turnOff(step.c)
			}
		}
	}

	nOnCubes := 0
	for c := range cuboids {
		if c.x >= -limit && c.x <= limit && c.y >= -limit && c.y <= limit && c.z >= -limit && c.z <= limit {
			nOnCubes++
		}
	}
	return nOnCubes
}

func turnOn(c cuboid) {
	for z := c.z0; z <= c.z1; z++ {
		for y := c.y0; y <= c.y1; y++ {
			for x := c.x0; x <= c.x1; x++ {
				cuboids[coord3D{x, y, z}] = struct{}{}
			}
		}
	}
}

func turnOff(c cuboid) {
	for z := c.z0; z <= c.z1; z++ {
		for y := c.y0; y <= c.y1; y++ {
			for x := c.x0; x <= c.x1; x++ {
				delete(cuboids, coord3D{x, y, z})
			}
		}
	}
}
