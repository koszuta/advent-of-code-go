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
	// Action N means to move the waypoint north by the given value.
	// Action S means to move the waypoint south by the given value.
	// Action E means to move the waypoint east by the given value.
	// Action W means to move the waypoint west by the given value.
	// Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
	// Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
	// Action F means to move forward to the waypoint a number of times equal to the given value.
	instructions := make([]instruction, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		action := line[:1]
		val, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, instruction{action, val})
	}

	shipX, shipY, x, y := 0, 0, 10, 1
	for _, c := range instructions {
		switch c.action {
		case "N":
			y += c.val
		case "S":
			y -= c.val
		case "E":
			x += c.val
		case "W":
			x -= c.val
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
