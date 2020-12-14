package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 14: Docking Data ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/14
 */

func applyMask(val int, mask string) int {
	newVal := 0
	for i, maskBit := range mask {
		newVal <<= 1
		switch maskBit {
		case '0': // push a 0
		case '1': // push a 1
			newVal++
		case 'X': // leave the bit unchanged
			bit := (val >> (35 - i)) & 1
			newVal += bit
		}
	}
	return newVal
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	var mask string
	mem := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "mask = ") {
			mask = line[7:]

		} else {
			parts := strings.Split(line, " = ")
			a := parts[0]
			address, _ := strconv.Atoi(a[4 : len(a)-1])
			value, _ := strconv.Atoi(parts[1])

			// Apply the mask to the value and store it in memory at the address
			mem[address] = applyMask(value, mask)
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}
	log.Printf("the sum of all values in memory is %d\n", sum)
}
