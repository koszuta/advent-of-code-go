package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
 *   --- Day 1: Report Repair ---
 *         --- Part Two ---
 *
 * Find 3 numbers from the puzzle input that sum to 2020.
 * What's the product of those 3 numbers?
 *
 * https://adventofcode.com/2020/day/1#part2
 */

type sumPair struct {
	s1, s2 int
}

const currentYear = 2020

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	// Store all numbers from the puzzle input
	vals := make([]int, 0, 0)
	{
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			val, err := strconv.Atoi(line)
			if err != nil {
				log.Panicln(err)
			}
			vals = append(vals, val)
		}
	}

	sums := make(map[int]sumPair)
OUT:
	for i := 0; i < len(vals); i++ {
		for j := i + 1; j < len(vals); j++ {
			sum := vals[i] + vals[j]
			if sum < currentYear {

				// Check if the remainder of 2020-val was already seen as a sum of 2 previous numbers
				pair, exists := sums[currentYear-vals[j]]
				if exists {
					log.Printf("%d\n", pair.s1*pair.s2*vals[j])
					break OUT
				}

				// Keep track of 2 number sums
				sums[sum] = sumPair{vals[i], vals[j]}
			}
		}
	}
}
