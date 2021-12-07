package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

/*
 *   --- Day 6: Lanternfish ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/6#part2
 */

const (
	nDays  = 256
	n0Fish = "1,1,3,5,1,1,1,4,1,5,1,1,1,1,1,1,1,3,1,1,1,1,2,5,1,1,1,1,1,2,1,4,1,4,1,1,1,1,1,3,1,1,5,1,1,1,4,1,1,1,4,1,1,3,5,1,1,1,1,4,1,5,4,1,1,2,3,2,1,1,1,1,1,1,1,1,1,1,1,1,1,5,1,1,1,1,1,1,1,1,2,2,1,1,1,1,1,5,1,1,1,3,4,1,1,1,1,3,1,1,1,1,1,4,1,1,3,1,1,3,1,1,1,1,1,3,1,5,2,3,1,2,3,1,1,2,1,2,4,5,1,5,1,4,1,1,1,1,2,1,5,1,1,1,1,1,5,1,1,3,1,1,1,1,1,1,4,1,2,1,1,1,1,1,1,1,1,1,1,1,1,1,3,2,1,1,1,1,2,2,1,2,1,1,1,5,5,1,1,1,1,1,1,1,1,1,1,1,1,2,2,1,1,4,2,1,4,1,1,1,1,1,1,1,2,1,2,1,1,1,1,1,1,1,1,1,1,1,1,1,2,2,1,5,1,1,1,1,1,1,1,1,3,1,1,3,3,1,1,1,3,5,1,1,4,1,1,1,1,1,4,1,1,3,1,1,1,1,1,1,1,1,2,1,5,1,1,1,1,1,1,1,1,1,1,4,1,1,1,1"
)

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nFish := doPart2()
	log.Println("the number of lanternfish after", nDays, "days is", nFish)
}

func doPart2() int {
	fish := make([]int, 0)
	for _, dayStr := range strings.Split(n0Fish, ",") {
		days, _ := strconv.Atoi(dayStr)
		fish = append(fish, days)
	}

	days := make([]int, nDays)
	for _, day := range fish {
		for d := day; d < len(days); d += 7 {
			days[d]++
		}
	}
	for day := 0; day < len(days)-9; day++ {
		for d := day + 9; d < len(days); d += 7 {
			days[d] += days[day]
		}
	}

	totalSpawns := 0
	for _, spawns := range days {
		totalSpawns += spawns
	}
	return len(fish) + totalSpawns
}
