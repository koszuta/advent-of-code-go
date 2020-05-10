package main

import (
	"fmt"
	"math"
	"time"
)

var n int = 36000000
var factor1 int = 10
var factor2 int = 11
var deliveryLimit int = 50

func part1() int {
	var house int
	for house = 1; ; house++ {
		factorsSum := 0
		limit := int(math.Sqrt(float64(house)))
		for i := 1; i <= limit; i++ {
			if house%i == 0 {
				factorsSum += i
				factorsSum += (house / i)
			}
		}
		if factorsSum*factor1 >= n {
			break
		}
	}
	return house
}

func part2() int {
	factorCount := make(map[int]int)
	var house int
	for house = 1; ; house++ {
		factorsSum := 0
		sqrt := int(math.Sqrt(float64(house)))
		for i := 1; i <= sqrt; i++ {
			if house%i == 0 {
				if factorCount[i] < deliveryLimit {
					factorsSum += i
					factorCount[i]++
				}
				f2 := house / i
				if factorCount[f2] < deliveryLimit {
					factorsSum += f2
					factorCount[f2]++
				}
			}
		}
		if factorsSum*factor2 >= n {
			break
		}
	}
	return house
}

func main() {
	var start time.Time
	var house int

	start = time.Now()
	house = part1()
	fmt.Printf("Part 1 took %v\n", time.Since(start))
	fmt.Printf("House %d\n", house)

	start = time.Now()
	house = part2()
	fmt.Printf("Part 2 took %v\n", time.Since(start))
	fmt.Printf("House %d\n", house)
}
