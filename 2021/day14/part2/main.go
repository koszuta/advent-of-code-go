package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

const expectedResult = 3353146900153

/*
 *   --- Day 14: Extended Polymerization ---
 *              --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/14#part2
 */

const steps = 40

type pair struct {
	p1, p2 string
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	diff := doPart2()
	log.Println("the quantity of the most common element subtracted from the quantity of the least common element is", diff)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	polymerTemplate := scanner.Text()

	scanner.Scan() // skip blank line

	rules := make(map[string]pair)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		pair1 := string(parts[0][0]) + parts[1]
		pair2 := parts[1] + string(parts[0][1])
		rules[parts[0]] = pair{
			pair1,
			pair2,
		}
	}

	pairCounts := make(map[string]int)
	for i := 1; i < len(polymerTemplate); i++ {
		pair := polymerTemplate[i-1 : i+1]
		pairCounts[pair]++
	}

	newPairCounts := make(map[string]int)
	for step := 0; step < steps; step++ {
		for pair, count := range pairCounts {
			newPairs := rules[pair]
			newPairCounts[newPairs.p1] += count
			newPairCounts[newPairs.p2] += count
			// Zero counts so you can just swap old and new instead of allocating every iteration
			pairCounts[pair] = 0
		}
		pairCounts, newPairCounts = newPairCounts, pairCounts
	}

	nLeastCommon, nMostCommon := math.MaxInt64, 0
	elementCounts := make(map[byte]int)
	for pair, count := range pairCounts {
		elementCounts[pair[0]] += count
		elementCounts[pair[1]] += count
	}
	for _, count := range elementCounts {
		if count < nLeastCommon {
			nLeastCommon = count
		}
		if count > nMostCommon {
			nMostCommon = count
		}
	}
	return (nMostCommon - nLeastCommon + 1) / 2
}
