package main

import (
	"advent-of-code-go/2016/day22/disc"
	"log"
)

/*
 *   --- Day 22: Grid Computing ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2016/day/22
 */

var startingState disc.GridState

func init() {
	startingState = disc.Parse("../input.txt")
}

func main() {
	nViablePairs := 0
	for _, nodeRow1 := range startingState.Nodes {
		for _, nodeA := range nodeRow1 {
			for _, nodeRow2 := range startingState.Nodes {
				for _, nodeB := range nodeRow2 {
					if nodeA != nodeB && nodeA.Used > 0 && nodeB.Available >= nodeA.Used {
						nViablePairs++
					}
				}
			}
		}
	}
	log.Printf("the number of viable pairs is %d\n", nViablePairs)
}
