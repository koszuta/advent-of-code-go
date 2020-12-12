package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	arg int
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// Parse the instructions
	instructions := make([]instruction, 0, 0)
	for scanner.Scan() {
		// Instructions are composed of an operation and an argument
		// Operations: acc, jmp, nop
		// Arguments: signed integers
		line := scanner.Text()
		parts := strings.Split(line, " ")
		op := parts[0]
		arg, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, instruction{op, arg})
	}

}
