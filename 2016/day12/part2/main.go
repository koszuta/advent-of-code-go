package main

import (
	"advent-of-code-go/2016/day12/assembunny"
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 12: Leonardo's Monorail ---
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/12#part2
 */

var (
	instructions []assembunny.Instruction
	registers    map[string]int
	registerSet  map[string]struct{}
)

func init() {
	registers = make(map[string]int)
	registerSet = make(map[string]struct{})
	for _, reg := range []string{"a", "b", "c", "d"} {
		registerSet[reg] = struct{}{}
		registers[reg] = 0
	}

	// init register c to 1
	registers["c"] = 1
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instrStr := scanner.Text()
		instruction, err := assembunny.Parse(instrStr, registerSet)
		if err != nil {
			log.Panicln(err)
		}
		instructions = append(instructions, instruction)
	}
	log.Println("instructions:", instructions)

	i := 0
	for i < len(instructions) {
		i = instructions[i].Execute(i, registers)
	}
	log.Println("registers:", registers)
}
