package main

import (
	"advent-of-code-go/2021/day20"
	"log"
	"time"
)

const expectedResult = 5619

/*
 *   --- Day 20: Trench Map ---
 *        --- Part One ---
 *
 *   https://adventofcode.com/2021/day/20
 */

const iterations = 2

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nLitPixels := doPart1()
	log.Println("the number of lit pixels after", iterations, "iterations is", nLitPixels)
}

func doPart1() int {
	return day20.Do(iterations)
}
