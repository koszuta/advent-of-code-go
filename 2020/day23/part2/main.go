package main

import (
	"fmt"
)

/*
 *   --- Day 23: Crab Cups ---
 *       --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/23#part2
 */

const moves = 10_000_000

// var input = [...]int{1, 5, 7, 6, 2, 3, 9, 8, 4}

var input = [...]int{3, 8, 9, 1, 2, 5, 4, 6, 7} // example

func main() {
	numCups := 1_000_000
	// numCups := len(input)

	cups := make([]int, numCups, numCups)
	for i, cup := range input {
		cups[i] = cup
		fmt.Printf("cups[%d]=%d\n", i, cup)
	}
	for i := len(input); i < numCups; i++ {
		cups[i] = i + 1
		if i < 20 || i > numCups-20 {
			fmt.Printf("cups[%d]=%d\n", i, cups[i])
		}
	}

	for i, q := 0, 0; q < 10*numCups; i, q = (i+1)%len(cups), q+1 {
		if q%99_999 == 0 {
			fmt.Printf("\n-- move %d --\n", q+1)
		}
		// fmt.Println("cups:", cups)

		nextCup := cups[0]
		threeCups := make([]int, 3, 3)
		copy(threeCups, cups[1:4])

		cups = append(cups[1:1], append(cups[4:], cups[0])...)
		// fmt.Println("pick:", threeCups)

		var dest int
	OUT:
		for {
			nextCup--
			if nextCup == 0 {
				nextCup = numCups
			}
			// fmt.Println("checking cup", nextCup)
			var cup int
			for dest, cup = range cups {
				if cup == nextCup {
					dest++
					break OUT
				}
			}
		}
		// fmt.Println("destination:", dest)
		// fmt.Println("destination:", nextCup)

		cups = append(cups[:dest], append(threeCups, cups[dest:]...)...)
		// fmt.Println("new cups", cups)
	}
	// fmt.Println("cups:", cups)
	fmt.Printf("\n-- final --\n")
	for i := 0; i < len(cups); i++ {
		if cups[i] == 1 {
			fmt.Println("cup[i+1]:", cups[(i+1)%len(cups)], "cup[i+2]:", cups[(i+2)%len(cups)])
			fmt.Println("prod:", cups[(i+1)%len(cups)]*cups[(i+2)%len(cups)])
			break
		}
	}
}
