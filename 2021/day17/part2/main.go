package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 3618

/*
 *   --- Day 17: Trick Shot ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/17#part2
 */

const input = "target area: x=206..250, y=-105..-57"

type coord2D struct {
	x, y int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nInitVels := doPart2()
	log.Println("the number of distinct initial velocity values which cause the probe to be within the target area is", nInitVels)
}

func doPart2() int {
	var x0, x1 int
	var y0, y1 int
	{
		parts := strings.Split(input, ", y=")
		xParts := strings.Split(parts[0][15:], "..")
		x0, _ = strconv.Atoi(xParts[0])
		x1, _ = strconv.Atoi(xParts[1])
		yParts := strings.Split(parts[1], "..")
		y1, _ = strconv.Atoi(yParts[0])
		y0, _ = strconv.Atoi(yParts[1])
	}

	initVels := make([]coord2D, 0)
	for velX := 1; velX <= x1; velX++ {
		for velY := y1; velY < -y1; velY++ {
			velX, velY := velX, velY
			var posX, posY int
			for posX <= x1 && posY >= y1 {
				if posX >= x0 && posX <= x1 && posY <= y0 && posY >= y1 {
					initVels = append(initVels, coord2D{velX, velY})
					break
				}
				// Update position and velocity
				posX += velX
				posY += velY
				if velX > 0 {
					velX--
				}
				velY--
			}
		}
	}

	return len(initVels)
}
