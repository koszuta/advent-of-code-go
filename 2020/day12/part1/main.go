package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type instruction struct {
	action string
	val    int
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func rotCW(x, y int) (int, int) {
	return y, -x
}

func rotCCW(x, y int) (int, int) {
	return -y, x
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// Parse the navigation instructions
	// Action N means to move north by the given value.
	// Action S means to move south by the given value.
	// Action E means to move east by the given value.
	// Action W means to move west by the given value.
	// Action L means to turn left the given number of degrees.
	// Action R means to turn right the given number of degrees.
	// Action F means to move forward by the given value in the direction the ship is currently facing.
	instructions := make([]instruction, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		action := line[:1]
		val, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, instruction{action, val})
	}

	shipX, shipY, x, y := 0, 0, 1, 0
	for _, c := range instructions {
		switch c.action {
		case "N":
			shipY += c.val
		case "S":
			shipY -= c.val
		case "E":
			shipX += c.val
		case "W":
			shipX -= c.val
		case "L":
			for i := 0; i < c.val/90; i++ {
				x, y = rotCCW(x, y)
			}
		case "R":
			for i := 0; i < c.val/90; i++ {
				x, y = rotCW(x, y)
			}
		case "F":
			shipX += x * c.val
			shipY += y * c.val
		}
	}

	log.Printf("the ship's final position is (%d, %d)\n", shipX, shipY)
	log.Printf("the Manhattan distance from the start is %d\n", abs(shipX)+abs(shipY))
}
