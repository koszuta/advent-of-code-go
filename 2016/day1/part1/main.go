package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
 *   --- Day 1: No Time for a Taxicab ---
 *            --- Part One ---
 *
 *   https://adventofcode.com/2016/day/1
 */

var position, heading vec2

type vec2 struct {
	x, y int
}

func init() {
	heading.y = 1
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Panicln(err)
	}

	for _, instruction := range strings.Split(string(file), ", ") {
		instr := []rune(instruction)
		dir := instr[0]
		updateHeading(dir)
		blocks, err := strconv.Atoi(string(instr[1:]))
		if err != nil {
			log.Panicln(err)
		}
		position.x += heading.x * blocks
		position.y += heading.y * blocks
	}

	log.Printf("final position is %v\n", position)
	log.Printf("manhattan distance from start is %d\n", abs(position.x)+abs(position.y))
}

func updateHeading(dir rune) {
	switch dir {
	case 'L':
		if heading.x == 0 {
			heading.x, heading.y = -heading.y, 0
		} else {
			heading.x, heading.y = 0, heading.x
		}
	case 'R':
		if heading.x == 0 {
			heading.x, heading.y = heading.y, 0
		} else {
			heading.x, heading.y = 0, -heading.x
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
