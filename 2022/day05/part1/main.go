package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	N        int
	From, To int
}

func main() {
	var stackLines, stepLines []string
	{ // separate stack and step lines
		b, _ := os.ReadFile("./input.txt")
		lines := strings.Split(string(b), "\n\n")

		stackLines = strings.Split(lines[0], "\n")
		stackLines = stackLines[:len(stackLines)-1]

		stepLines = strings.Split(lines[1], "\n")
	}

	var stacks [][]string
	var steps []Step
	{ // parse stacks and steps
		stacks = make([][]string, len(stackLines[0])/4)
		for _, line := range stackLines {
			for i := 0; len(line) >= 4; i++ {
				crate := line[1:2]
				line = line[4:]
				if crate == " " {
					continue
				}
				stacks[i] = append(stacks[i], crate)
			}
		}

		steps = make([]Step, 0, len(stepLines))
		for _, line := range stepLines {
			step := parseStep(line)
			steps = append(steps, step)
		}
	}

	for _, step := range steps {
		for i := 0; i < step.N; i++ {
			stack := stacks[step.From]
			crate := stack[0]
			stacks[step.From] = stacks[step.From][1:]
			stacks[step.To] = append([]string{crate}[:], stacks[step.To]...)
		}
	}

	var result string
	for _, stack := range stacks {
		result += stack[0]
	}

	fmt.Println("top crates:", result)
}

func parseStep(line string) Step {
	splits := strings.Split(line, " ")
	n, _ := strconv.Atoi(splits[1])
	from, _ := strconv.Atoi(splits[3])
	to, _ := strconv.Atoi(splits[5])
	return Step{
		N:    n,
		From: from - 1,
		To:   to - 1,
	}
}
