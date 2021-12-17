package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 5460

/*
 *   --- Day 17: Trick Shot ---
 *        --- Part One ---
 *
 *   https://adventofcode.com/2021/day/17
 */

const input = "target area: x=206..250, y=-105..-57"

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	maxY := doPart1()
	log.Println("the highest y position the probe reaches is", maxY)
}

func doPart1() int {
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

	maxY := 0
	for velX := 1; velX <= x1; velX++ {
		// We're only concerned with shots that go so high that
		// they end up falling straight down within the X range
		xDist := (velX + 1) * velX / 2
		if xDist < x0 || xDist > x1 {
			continue
		}
		for velY := 0; velY < -y1; velY++ {
			velX, velY := velX, velY
			var posX, posY int
			attemptMaxY := 0
			// While the probe isn't beyond the target range
			for posX <= x1 && posY >= y1 {
				if posY > attemptMaxY {
					attemptMaxY = posY
				}
				if posX >= x0 && posX <= x1 && posY <= y0 && posY >= y1 {
					if attemptMaxY > maxY {
						maxY = attemptMaxY
					}
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

	return maxY
}
