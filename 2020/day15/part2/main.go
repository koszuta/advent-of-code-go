package main

import "log"

/*
 *   --- Day 15: Rambunctious Recitation ---
 *              --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/15#part2
 */

var startingNums = [...]int{2, 0, 1, 7, 4, 14, 18}
var maxTurn = 30000000

func main() {
	prev, turnSpoken := 0, make(map[int]int)
	for turn := 0; turn < maxTurn; turn++ {
		if turn < len(startingNums) {
			// Begin by saying the starting numbers
			prev = startingNums[turn]
			turnSpoken[prev] = turn

		} else {
			prevTurn := turn - 1
			lastTurnSpoken, hasBeenSpoken := turnSpoken[prev]

			// Store which turn the previous number was spoken on
			turnSpoken[prev] = prevTurn

			if hasBeenSpoken {
				// Since the previous number has already been spoken, say the age of the previous number
				// i.e. how many turns since is was previously spoken
				prev = prevTurn - lastTurnSpoken

			} else {
				// Since the previous number hasn't been spoken yet, say 0
				prev = 0
			}
		}
	}

	log.Printf("the %dth number spoken is %d\n", maxTurn, prev)
}
