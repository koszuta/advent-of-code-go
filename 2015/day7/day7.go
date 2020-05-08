package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var instructions map[string]string = make(map[string]string)
var digitsRegex *regexp.Regexp = regexp.MustCompile(`\d+`)

func parse(val string) int {
	if !digitsRegex.MatchString(val) {
		val = execute(val)
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return intVal
}

func execute(wire string) string {
	parts := strings.Split(instructions[wire], " ")

	if len(parts) == 1 {
		if !digitsRegex.MatchString(instructions[wire]) {
			instructions[wire] = execute(instructions[wire])
		}
	} else {
		if strings.Compare(parts[0], "NOT") == 0 {
			wire1 := parse(parts[1])
			instructions[wire] = strconv.Itoa(int(^uint16(wire1)))

		} else {
			wire1 := parse(parts[0])
			wire2 := parse(parts[2])
			if strings.Compare(parts[1], "AND") == 0 {
				instructions[wire] = strconv.Itoa(int(uint16(wire1) & uint16(wire2)))
				
			} else if strings.Compare(parts[1], "OR") == 0 {
				instructions[wire] = strconv.Itoa(int(uint16(wire1) | uint16(wire2)))
				
			} else if strings.Compare(parts[1], "LSHIFT") == 0 {
				instructions[wire] = strconv.Itoa(int(uint16(wire1) << uint16(wire2)))
				
			} else if strings.Compare(parts[1], "RSHIFT") == 0 {
				instructions[wire] = strconv.Itoa(int(uint16(wire1) >> uint16(wire2)))
			}
		}
	}

	return instructions[wire]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		instructions[parts[1]] = parts[0]
	}

	signalA := execute("a")
	fmt.Printf("signal of a=%s\n", signalA)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		instructions[parts[1]] = parts[0]
	}
	instructions["b"] = signalA

	signalA = execute("a")
	fmt.Printf("signal of a=%s\n", signalA)
}
