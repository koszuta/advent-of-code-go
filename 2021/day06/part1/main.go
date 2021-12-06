package main

import (
	"log"
	"strconv"
	"strings"
)

/*
 *   --- Day 6: Lanternfish ---
 *        --- Part One ---
 *
 *   https://adventofcode.com/2021/day/6
 */

const (
	totalDays   = 80
	initialFish = "1,1,3,5,1,1,1,4,1,5,1,1,1,1,1,1,1,3,1,1,1,1,2,5,1,1,1,1,1,2,1,4,1,4,1,1,1,1,1,3,1,1,5,1,1,1,4,1,1,1,4,1,1,3,5,1,1,1,1,4,1,5,4,1,1,2,3,2,1,1,1,1,1,1,1,1,1,1,1,1,1,5,1,1,1,1,1,1,1,1,2,2,1,1,1,1,1,5,1,1,1,3,4,1,1,1,1,3,1,1,1,1,1,4,1,1,3,1,1,3,1,1,1,1,1,3,1,5,2,3,1,2,3,1,1,2,1,2,4,5,1,5,1,4,1,1,1,1,2,1,5,1,1,1,1,1,5,1,1,3,1,1,1,1,1,1,4,1,2,1,1,1,1,1,1,1,1,1,1,1,1,1,3,2,1,1,1,1,2,2,1,2,1,1,1,5,5,1,1,1,1,1,1,1,1,1,1,1,1,2,2,1,1,4,2,1,4,1,1,1,1,1,1,1,2,1,2,1,1,1,1,1,1,1,1,1,1,1,1,1,2,2,1,5,1,1,1,1,1,1,1,1,3,1,1,3,3,1,1,1,3,5,1,1,4,1,1,1,1,1,4,1,1,3,1,1,1,1,1,1,1,1,2,1,5,1,1,1,1,1,1,1,1,1,1,4,1,1,1,1"
)

func main() {
	fish := make([]int, 0)
	for _, dayStr := range strings.Split(initialFish, ",") {
		days, _ := strconv.Atoi(dayStr)
		fish = append(fish, days)
	}

	for i := 0; i < totalDays; i++ {
		for i, days := range fish {
			if days == 0 {
				fish[i] = 6
				fish = append(fish, 8)
			} else {
				fish[i]--
			}
		}
	}
	log.Println("the number of lanternfish after", totalDays, "days is", len(fish))
}
