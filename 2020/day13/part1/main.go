package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 13: Shuttle Search ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2020/day/13
 */

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// The first line of the input is the earliest time at which you could depart
	var time int
	{
		scanner.Scan()
		time, _ = strconv.Atoi(scanner.Text())
	}

	// The second line is a list of bus IDs
	// Entries that show x are not used
	buses := make([]int, 0, 0)
	{
		scanner.Scan()
		for _, id := range strings.Split(scanner.Text(), ",") {
			if id != "x" {
				busID, _ := strconv.Atoi(id)
				buses = append(buses, busID)
			}
		}
	}

	// Buses depart at multiples of their ID
	// Find the bus which leaves the soonest after the given time
	minWait, minBusID := buses[0], buses[0]
	for _, busID := range buses {
		wait := busID - time%busID
		if wait < minWait {
			minWait = wait
			minBusID = busID
		}
	}

	log.Printf("the ID of the earliest bus multiplied by the time needed to wait for that bus is %d\n", minWait*minBusID)
}
