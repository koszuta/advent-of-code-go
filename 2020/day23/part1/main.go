package main

import (
	"fmt"
	"log"
)

/*
 *   --- Day 23: Crab Cups ---
 *       --- Part One ---
 *
 *   https://adventofcode.com/2020/day/23
 */

const moves = 100

var input = [...]int{1, 5, 7, 6, 2, 3, 9, 8, 4}

// var input = [...]int{3, 8, 9, 1, 2, 5, 4, 6, 7} // example

func main() {
	// Puzzle input
	cups := input[:]
	nCups := len(cups)

	for q := 1; q <= moves; q++ {
		// fmt.Printf("\n-- move %d --\n", q)
		// fmt.Println("cups:", cups)

		currentCup := cups[0]
		cup1, cup2, cup3 := cups[1], cups[2], cups[3]
		// fmt.Printf("pick: %d, %d, %d\n", cup1, cup2, cup3)

		destCup := currentCup - 1
		if destCup < 1 {
			destCup = nCups
		}
		for destCup == cup1 || destCup == cup2 || destCup == cup3 {
			if destCup == 1 {
				destCup = nCups
			} else {
				destCup--
			}
		}
		// fmt.Println("destination:", destCup)

		for dest, cup := range cups {
			if cup == destCup {
				newCups := make([]int, 0, nCups)
				newCups = append(newCups, cups[4:dest+1]...)
				newCups = append(newCups, cup1, cup2, cup3)
				newCups = append(newCups, cups[dest+1:]...)
				newCups = append(newCups, currentCup)
				cups = newCups
				break
			}
		}
	}
	// fmt.Print("\n-- final --\n")
	// fmt.Printf("cups: %d\n\n", cups)

	labels := ""
	for i := 0; i < nCups; i++ {
		if cups[i] == 1 {
			for j := (i + 1) % nCups; j != i; j = (j + 1) % nCups {
				labels += fmt.Sprint(cups[j])
			}
			break
		}
	}
	log.Printf("after %d moves, the labels on the cups after '1' are %s\n", moves, labels)
}
