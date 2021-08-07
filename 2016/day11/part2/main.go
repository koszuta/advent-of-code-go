package main

import (
	"advent-of-code-go/2016/day11/part2/rtg"
	"fmt"
	"time"
)

/*
 *   --- Day 11: Radioisotope Thermoelectric Generators ---
 *                      --- Part One ---
 *
 *   https://adventofcode.com/2016/day/11
 */

var (
	hashGraph     map[rtg.Hash]map[rtg.Hash]struct{}
	startingState rtg.State
)

func init() {
	hashGraph = make(map[rtg.Hash]map[rtg.Hash]struct{})

	// startingState = rtg.NewState(0, [rtg.NElements]int{1, 2}, [rtg.NElements]int{0, 0})
	// example(startingState)
	startingState = rtg.NewState(0, [rtg.NElements]int{0, 0, 0, 0, 0, 0, 0}, [rtg.NElements]int{1, 1, 0, 0, 0, 0, 0})

	fmt.Println("starting state:")
	fmt.Println(startingState.ToString())
	fmt.Println()
}

func main() {
	defer (func(start time.Time) {
		fmt.Printf("... took %v\n", time.Since(start))
	})(time.Now())

	doMove(0, &startingState)

	dist := bfs(startingState.Hash())
	fmt.Println("depth:", dist)
}

type hashNode struct {
	hash rtg.Hash
	dist int
}

func bfs(root rtg.Hash) int {
	visited := make(map[rtg.Hash]struct{})
	queue := make([]hashNode, 0)

	visited[root] = struct{}{}
	queue = append(queue, hashNode{root, 0})

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v.hash == rtg.FinalStateHash {
			return v.dist
		}
		for w := range hashGraph[v.hash] {
			_, seen := visited[w]
			if !seen {
				queue = append(queue, hashNode{w, v.dist + 1})
				visited[w] = struct{}{}
			}
		}
	}
	return -1
}

func doMove(depth int, currentState *rtg.State) {
	if currentState.Hash() == rtg.FinalStateHash {
		return
	}

	var nextState *rtg.State
	for i := 0; i < rtg.NElements; i++ {
		nextState = currentState.MoveChipUp(i)
		if storeMove(currentState, nextState, fmt.Sprintf("move microchip %d up:", i)) {
			doMove(depth+1, nextState)
		}

		nextState = currentState.MoveChipDown(i)
		if storeMove(currentState, nextState, fmt.Sprintf("move microchip %d down:", i)) {
			doMove(depth+1, nextState)
		}

		nextState = currentState.MoveGenUp(i)
		if storeMove(currentState, nextState, fmt.Sprintf("move generator %d up:", i)) {
			doMove(depth+1, nextState)
		}

		nextState = currentState.MoveGenDown(i)
		if storeMove(currentState, nextState, fmt.Sprintf("move generator %d down:", i)) {
			doMove(depth+1, nextState)
		}

		nextState = currentState.MovePairUp(i)
		if storeMove(currentState, nextState, fmt.Sprintf("move gen and chip %d up:", i)) {
			doMove(depth+1, nextState)
		}

		nextState = currentState.MovePairDown(i)
		if storeMove(currentState, nextState, fmt.Sprintf("move gen and chip %d down:", i)) {
			doMove(depth+1, nextState)
		}

		for j := i + 1; j < rtg.NElements; j++ {
			nextState = currentState.MoveChipsUp(i, j)
			if storeMove(currentState, nextState, fmt.Sprintf("move microchips %d and %d up:", i, j)) {
				doMove(depth+1, nextState)
			}

			nextState = currentState.MoveChipsDown(i, j)
			if storeMove(currentState, nextState, fmt.Sprintf("move microchips %d and %d down:", i, j)) {
				doMove(depth+1, nextState)
			}

			nextState = currentState.MoveGensUp(i, j)
			if storeMove(currentState, nextState, fmt.Sprintf("move generators %d and %d up:", i, j)) {
				doMove(depth+1, nextState)
			}

			nextState = currentState.MoveGensDown(i, j)
			if storeMove(currentState, nextState, fmt.Sprintf("move generators %d and %d down:", i, j)) {
				doMove(depth+1, nextState)
			}
		}
	}
}

func storeMove(currentState, nextState *rtg.State, message string) bool {
	if nextState != nil {
		currentHash := currentState.Hash()
		_, found := hashGraph[currentHash]
		if !found {
			hashGraph[currentHash] = make(map[rtg.Hash]struct{})
		}
		nextHash := nextState.Hash()
		hashGraph[currentHash][nextHash] = struct{}{}
		_, found = hashGraph[nextHash]
		return !found
	}
	return false
}

func example(startingState rtg.State) {
	fmt.Println("example solution:")
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipUp(0)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MovePairUp(0)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipDown(0)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipDown(0)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipsUp(0, 1)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipsUp(0, 1)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipsUp(0, 1)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipDown(0)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveGensUp(0, 1)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipDown(1)
	fmt.Println(startingState.ToString())
	fmt.Println()
	startingState = *startingState.MoveChipsUp(0, 1)
	fmt.Println(startingState.ToString())
	fmt.Println()
}
