package main

import (
	"advent-of-code-go/2021/day20"
	"log"
	"time"
)

const expectedResult = 20122

/*
 *   --- Day 20: Trench Map ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/20#part2
 */

const iterations = 50

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nLitPixels := doPart2()
	log.Println("the number of lit pixels after", iterations, "iterations is", nLitPixels)
}

func doPart2() int {
	return day20.Do(iterations)
}
