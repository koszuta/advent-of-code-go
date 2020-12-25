package main

import (
	"log"
)

/*
 *   --- Day 23: Crab Cups ---
 *       --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/23#part2
 */

const nCups, moves = 1_000_000, 10_000_000

var input = [...]int{1, 5, 7, 6, 2, 3, 9, 8, 4}

// var input = [...]int{3, 8, 9, 1, 2, 5, 4, 6, 7} // example

type node struct {
	label int
	next  *node
}

func main() {
	cups := make(map[int]*node)
	var currentCup, lastCup *node
	for i := nCups; i > 0; i-- {
		label := i
		if i <= len(input) {
			label = input[i-1]
		}
		cup := node{label, currentCup}
		cups[label] = &cup
		currentCup = &cup
		if lastCup == nil {
			lastCup = currentCup
		}
	}
	lastCup.next = currentCup

	// start := time.Now()
	for q := 1; q <= moves; q++ {
		cup1 := currentCup.next
		cup2 := cup1.next
		cup3 := cup2.next

		nextCup := currentCup.label - 1
		if nextCup < 1 {
			nextCup = nCups
		}
		for nextCup == cup1.label || nextCup == cup2.label || nextCup == cup3.label {
			if nextCup == 1 {
				nextCup = nCups
			} else {
				nextCup--
			}
		}

		currentCup.next = cup3.next
		destCup := cups[nextCup]
		cup3.next = destCup.next
		destCup.next = cup1
		currentCup = currentCup.next
	}
	// t := time.Since(start)
	// log.Println("took:", t)
	// log.Printf("avg. time: %fns\n", float64(t.Nanoseconds())/float64(moves))

	nextCup := cups[1].next
	nextNextCup := nextCup.next
	log.Printf(".., 1, %d, %d,..\n", nextCup.label, nextNextCup.label)
	log.Println("the product of the two cups immediately clockwise of cup '1' is", nextCup.label*nextNextCup.label)
}
