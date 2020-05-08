package main

import (
	"fmt"
)

const noun1, verb1 = 12, 2
const part2 = 19690720

var program = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 2, 19, 13, 23, 1, 23, 10, 27, 1, 13, 27, 31, 2, 31, 10, 35, 1, 35, 9, 39, 1, 39, 13, 43, 1, 13, 43, 47, 1, 47, 13, 51, 1, 13, 51, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 9, 63, 67, 1, 6, 67, 71, 2, 71, 13, 75, 2, 75, 13, 79, 1, 79, 9, 83, 2, 83, 10, 87, 1, 9, 87, 91, 1, 6, 91, 95, 1, 95, 10, 99, 1, 99, 13, 103, 1, 13, 103, 107, 2, 13, 107, 111, 1, 111, 9, 115, 2, 115, 10, 119, 1, 119, 5, 123, 1, 123, 2, 127, 1, 127, 5, 0, 99, 2, 14, 0, 0}

func main() {
	fmt.Printf("Go...\n")

	part1 := intcode(noun1, verb1)
	fmt.Printf("noun1=%d verb1=%d\n", noun1, verb1)
	fmt.Printf("part1=%d\n", part1)

OUT_OF_PART2:
	for noun := range program {
		for verb := range program {
			if intcode(noun, verb) == part2 {
				fmt.Printf("noun2=%d verb2=%d\n", noun, verb)
				fmt.Printf("part2=%d\n", part2)
				break OUT_OF_PART2
			}
		}
	}
}

func intcode(noun, verb int) int {
	prog := make([]int, len(program))
	copy(prog, program)

	prog[1] = noun
	prog[2] = verb

	for i := 0; prog[i] != 99; i += 4 {
		switch prog[i] {
		case 1:
			prog[prog[i+3]] = prog[prog[i+1]] + prog[prog[i+2]]
			break
		case 2:
			prog[prog[i+3]] = prog[prog[i+1]] * prog[prog[i+2]]
			break
		}
	}

	return prog[0]
}
