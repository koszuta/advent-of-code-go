package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
 *   --- Day 1: No Time for a Taxicab ---
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/1#part2
 */

var position, heading vec2
var prevPositions map[vec2]struct{}

type vec2 struct {
	x, y int
}

func init() {
	heading.y = 1
	prevPositions = make(map[vec2]struct{})
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Panicln(err)
	}

OUTER:
	for _, instruction := range strings.Split(string(file), ", ") {
		instr := []rune(instruction)
		dir := instr[0]
		updateHeading(dir)
		blocks, err := strconv.Atoi(string(instr[1:]))
		if err != nil {
			log.Panicln(err)
		}

		for i := 0; i < blocks; i++ {
			position.x += heading.x
			position.y += heading.y

			_, found := prevPositions[position]
			if found {
				break OUTER
			}
			prevPositions[position] = struct{}{}
		}
	}

	log.Printf("position of first location visited twice is %v\n", position)
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
