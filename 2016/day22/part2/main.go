package main

import (
	"advent-of-code-go/2016/day22/disc"
	"fmt"
	"log"
)

/*
 *   --- Day 22: Grid Computing ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/22#part2
 */

var (
	startingState disc.GridState
	states        map[disc.GridState]struct{}
)

func init() {
	states = make(map[disc.GridState]struct{})
	startingState = disc.Parse("../input.txt")
	fmt.Println(startingState.String())
}

func main() {
	states[startingState] = struct{}{}
	// steps := moveData(0, math.MaxInt64, startingState)

	// I manually found the route through the printed grid
	// Just use the same technique as the example; it's pretty straight forward
	steps := 220
	log.Printf("the fewest number of steps required to move the goal data to node-x0-y0 is %d\n", steps)
}

func moveData(steps, minSteps int, state disc.GridState) int {
	if steps >= minSteps {
		return minSteps
	}
	if state.DataIsAtTarget() {
		if steps < minSteps {
			minSteps = steps
			log.Println("new least steps is", minSteps)
		} else {
			log.Println("steps:", steps, "minSteps:", minSteps)
		}
		return minSteps
	}

	for y := 0; y < disc.Height; y++ {
		for x := 0; x < disc.Width; x++ {
			doMoveData := func(nextPos disc.Vec2i) {
				nextState, moved := state.MoveData(disc.Vec2i{X: x, Y: y}, nextPos)
				if moved {
					_, seen := states[nextState]
					if !seen {
						states[nextState] = struct{}{}
						minSteps = moveData(steps+1, minSteps, nextState)
						delete(states, nextState)
					}
				}
			}
			// Up
			doMoveData(disc.Vec2i{X: x, Y: y - 1})
			// Down
			doMoveData(disc.Vec2i{X: x, Y: y + 1})
			// Left
			doMoveData(disc.Vec2i{X: x - 1, Y: y})
			// Right
			doMoveData(disc.Vec2i{X: x + 1, Y: y})
		}
	}

	return minSteps
}
