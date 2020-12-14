package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 *   --- Day 13: Shuttle Search ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/13#part2
 */

type bus struct {
	id, offset int
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// The first line of the input isn't used in Part Two
	scanner.Scan()

	// The second line is a list of bus IDs
	// Store both their ID and index in the list
	buses := make([]bus, 0, 0)
	{
		scanner.Scan()
		for i, id := range strings.Split(scanner.Text(), ",") {
			if id != "x" {
				busID, _ := strconv.Atoi(id)
				buses = append(buses, bus{busID, i})
			}
		}
	}

	// Sort the buses by ID in descending order
	sort.Slice(buses, func(i, j int) bool {
		return buses[i].id > buses[j].id
	})

	// Use the Chinese Remainder Theorem
	time, jump := 0, 1
	for i := 0; i < len(buses); i++ {
		for (time+buses[i].offset)%buses[i].id != 0 {
			time += jump
		}
		jump *= buses[i].id
	}

	log.Printf("the earliest time where all of the listed bus IDs depart at offsets matching their positions in the list is %d\n", time)
}
