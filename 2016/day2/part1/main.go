package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
 *   --- Day 2: Bathroom Security ---
 *           --- Part One ---
 *
 *   https://adventofcode.com/2016/day/2
 */

var keypad = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var pos vec2

type vec2 struct {
	x, y int
}

func init() {
	pos.x, pos.y = 1, 1
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	code := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, dir := range scanner.Text() {
			updatePosition(dir)
		}

		key := keypad[pos.y][pos.x]
		code += strconv.Itoa(key)
	}

	log.Printf("the bathroom code is %s\n", code)
}

func updatePosition(dir rune) {
	switch dir {
	case 'U':
		if pos.y > 0 {
			pos.y--
		}
	case 'D':
		if pos.y < 2 {
			pos.y++
		}
	case 'L':
		if pos.x > 0 {
			pos.x--
		}
	case 'R':
		if pos.x < 2 {
			pos.x++
		}
	}
}
