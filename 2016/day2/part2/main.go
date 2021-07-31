package main

import (
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 2: Bathroom Security ---
 *           --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/2#part2
 */

const keys = "123456789ABCD"

var keypad = [5][5]int{{0, 0, 1, 0, 0}, {0, 2, 3, 4, 0}, {5, 6, 7, 8, 9}, {0, 10, 11, 12, 0}, {0, 0, 13, 0, 0}}
var pos vec2

type vec2 struct {
	x, y int
}

func init() {
	pos.x, pos.y = 0, 2
}

func main() {
	code := ""

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, dir := range scanner.Text() {
			updatePosition(dir)
		}

		key := keypad[pos.y][pos.x]
		code += string(keys[key-1])
	}

	log.Printf("the bathroom code is %s\n", code)
}

func updatePosition(dir rune) {
	switch dir {
	case 'U':
		if pos.y > 0 && keypad[pos.y-1][pos.x] != 0 {
			pos.y--
		}
	case 'D':
		if pos.y < 4 && keypad[pos.y+1][pos.x] != 0 {
			pos.y++
		}
	case 'L':
		if pos.x > 0 && keypad[pos.y][pos.x-1] != 0 {
			pos.x--
		}
	case 'R':
		if pos.x < 4 && keypad[pos.y][pos.x+1] != 0 {
			pos.x++
		}
	}
}
