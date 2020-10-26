package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var registers map[string]uint = make(map[string]uint)

func loadInstructions(fileName string) [][]string {
	instructions := make([][]string, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Panicln(err)
	}

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		line = strings.ReplaceAll(line, ",", "")
		instrParts := strings.Split(line, " ")
		instructions = append(instructions, instrParts)
	}

	return instructions
}

func execute(instructions [][]string) {
	instrPointer := 0
	for instrPointer < len(instructions) {
		instr := instructions[instrPointer]
		switch instr[0] {
		case "hlf":
			registers[instr[1]] /= 2
			instrPointer++
		case "tpl":
			registers[instr[1]] *= 3
			instrPointer++
		case "inc":
			registers[instr[1]]++
			instrPointer++
		case "jmp":
			jmp, err := strconv.Atoi(instr[1])
			if err != nil {
				log.Panicln(err)
			}
			instrPointer += jmp
		case "jie":
			if registers[instr[1]]%2 == 0 {
				jmp, err := strconv.Atoi(instr[2])
				if err != nil {
					log.Panicln(err)
				}
				instrPointer += jmp
			} else {
				instrPointer++
			}
		case "jio":
			if registers[instr[1]] == 1 {
				jmp, err := strconv.Atoi(instr[2])
				if err != nil {
					log.Panicln(err)
				}
				instrPointer += jmp
			} else {
				instrPointer++
			}
		}
	}
}

func main() {
	example := loadInstructions("example.txt")
	registers["a"] = 0
	registers["b"] = 0
	execute(example)
	log.Printf("Example registers=%v\n", registers)

	instructions := loadInstructions("input.txt")
	registers["a"] = 0
	registers["b"] = 0
	execute(instructions)
	log.Printf("Part 1 registers=%v\n", registers)
	registers["a"] = 1
	registers["b"] = 0
	execute(instructions)
	log.Printf("Part 2 registers=%v\n", registers)
}
