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
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/14#part2
 */

var mem map[int]int

func applyMask(val int, mask string) []rune {
	newVal := []rune(mask)
	for i, c := range mask {
		if c == '0' {
			bit := (val >> (35 - i)) & 1
			newVal[i] = '0' + rune(bit)
		}
	}
	return newVal
}

func write(address []rune, val, depth int) {
	if depth == 36 {
		a, _ := strconv.ParseUint(string(address), 2, 36)
		mem[int(a)] = val

	} else if address[depth] == 'X' {
		address[depth] = '0'
		write(address, val, depth+1)
		address[depth] = '1'
		write(address, val, depth+1)
		address[depth] = 'X'

	} else {
		write(address, val, depth+1)
	}
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	var mask string
	mem = make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "mask = ") {
			mask = line[7:]

		} else {
			parts := strings.Split(line, " = ")

			address, _ := strconv.Atoi(parts[0][4 : len(parts[0])-1])
			value, _ := strconv.Atoi(parts[1])

			maskedAddr := applyMask(address, mask)
			write(maskedAddr, value, 0)
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}
	log.Printf("the sum of all values in memory is %d\n", sum)
}
