package main

import (
	"advent-of-code-go/2016/day23/assembunny"
	"bufio"
	"log"
	"os"
)

/*
 *   --- Day 23: Safe Cracking ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/23#part2
 */

var (
	instructions []assembunny.Instruction
	registers    map[string]int
)

func init() {
	registers = make(map[string]int)
	for _, reg := range []string{"a", "b", "c", "d"} {
		registers[reg] = 0
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instrStr := scanner.Text()
		instruction, err := assembunny.Parse(instrStr, registers)
		if err != nil {
			log.Panicln(err)
		}
		instructions = append(instructions, *instruction)
	}
	for _, instruction := range instructions {
		log.Printf("%#v\n", instruction)
	}
	log.Println()
}

func main() {
	// Init register a to be the number of eggs
	registers["a"] = 12
	log.Println("initial registers", registers)

	i := 0
	for i >= 0 && i < len(instructions) {
		instructions, i = instructions[i].Execute(i, registers, instructions)
	}
	log.Printf("the value in register a to send to the safe is %d\n", registers["a"])
}
