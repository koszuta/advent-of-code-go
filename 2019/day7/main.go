package main

import (
	"fmt"
)

const minPhase1 = 0
const maxPhase1 = 4
const feedbackMode1 = false

const minPhase2 = 5
const maxPhase2 = 9
const feedbackMode2 = true

// var program = []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 42, 51, 60, 77, 94, 175, 256, 337, 418, 99999, 3, 9, 1001, 9, 4, 9, 102, 5, 9, 9, 1001, 9, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 1002, 9, 4, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 3, 9, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}
var program = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

func main() {
	fmt.Printf("Go...\n")

	maxOutput := amplify(minPhase1, maxPhase1, feedbackMode1)
	fmt.Printf("maxOutput=%d feedbackMode=%v\n", maxOutput, feedbackMode1)

	maxOutput = amplify(minPhase2, maxPhase2, feedbackMode2)
	fmt.Printf("maxOutput=%d feedbackMode=%v\n", maxOutput, feedbackMode2)
}

func amplify(minPhase int, maxPhase int, feedbackMode bool) int {
	var ok bool
	output := 0
	maxOutput := 0
	phases := make([]int, 0)
	for i := minPhase; i <= maxPhase; i++ {
		phases = append(phases, i)
		for j := minPhase; j <= maxPhase; j++ {
			if contains(j, phases) {
				continue
			}
			phases = append(phases, j)
			for k := minPhase; k <= maxPhase; k++ {
				if contains(k, phases) {
					continue
				}
				phases = append(phases, k)
				for l := minPhase; l <= maxPhase; l++ {
					if contains(l, phases) {
						continue
					}
					phases = append(phases, l)
					for m := minPhase; m <= maxPhase; m++ {
						if contains(m, phases) {
							continue
						}
						phases = append(phases, m)

						// for output != prev {
						for q := 0; q < len(phases); q++ {
							output, ok = intcode(phases[q], output)
						}

						if ok {
							fmt.Printf("(%d, %d, %d, %d, %d)=%d\n", i, j, k, l, m, output)
							if output > maxOutput {
								maxOutput = output
							}
						}
						// }

						output = 0
						phases = phases[:len(phases)-1]
					}
					phases = phases[:len(phases)-1]
				}
				phases = phases[:len(phases)-1]
			}
			phases = phases[:len(phases)-1]
		}
		phases = phases[:len(phases)-1]
	}

	return maxOutput
}

func intcode(args ...int) (int, bool) {
	prog := make([]int, len(program))
	copy(prog, program)

	var output int
	hasOutput := false

	opSize := 0
	argsIndex := 0
	for i := 0; prog[i] != 99; i += opSize {
		// fmt.Printf("%d ", prog[i])

		opCode := prog[i] % 100
		modes := prog[i] / 100

		var numOpParams int
		allowImmediateMode := true
		switch opCode {
		case 1, 2, 7, 8:
			opSize = 4
			numOpParams = 2
		case 5, 6:
			opSize = 3
			numOpParams = 2
		case 3:
			allowImmediateMode = false
			fallthrough
		default:
			opSize = 2
			numOpParams = 1
		}

		opParams := make([]int, numOpParams)
		for j := 0; j < numOpParams; j, modes = j+1, modes/10 {
			if allowImmediateMode && modes%10 == 0 {
				opParams[j] = prog[prog[i+j+1]]
				// fmt.Printf("[%d]=%d ", prog[i+j+1], opParams[j])
			} else {
				opParams[j] = prog[i+j+1]
				// fmt.Printf("%d ", opParams[j])
			}
		}

		switch opCode {
		case 1:
			prog[prog[i+3]] = opParams[0] + opParams[1]
			// fmt.Printf("\n[%d]<-%d\n", prog[i+3], prog[prog[i+3]])
		case 2:
			prog[prog[i+3]] = opParams[0] * opParams[1]
			// fmt.Printf("\n[%d]<-%d\n", prog[i+3], prog[prog[i+3]])
		case 3:
			for argsIndex >= len(args) {}
			prog[opParams[0]] = args[argsIndex]
			argsIndex++
			// fmt.Printf("\n[%d]<-%d\n", opParams[0], input)
		case 4:
			output = opParams[0]
			hasOutput = true
			// fmt.Printf("\n>>> %d\n", output)
		case 5:
			// fmt.Printf("\n")
			if opParams[0] != 0 {
				i = opParams[1]
				opSize = 0
				// fmt.Printf("Go to %d\n", i)
			}
		case 6:
			// fmt.Printf("\n")
			if opParams[0] == 0 {
				i = opParams[1]
				opSize = 0
				// fmt.Printf("Go to %d\n", i)
			}
		case 7:
			if opParams[0] < opParams[1] {
				prog[prog[i+3]] = 1
			} else {
				prog[prog[i+3]] = 0
			}
			// fmt.Printf("\n[%d]<-%d\n", prog[i+3], prog[prog[i+3]])
		case 8:
			if opParams[0] == opParams[1] {
				prog[prog[i+3]] = 1
			} else {
				prog[prog[i+3]] = 0
			}
			// fmt.Printf("\n[%d]<-%d\n", prog[i+3], prog[prog[i+3]])
		}
	}

	return output, hasOutput
}

func contains(target int, arr []int) bool {
	for _, i := range arr {
		if i == target {
			return true
		}
	}
	return false
}
