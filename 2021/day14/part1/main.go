package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

const expectedResult = 2915

/*
 *   --- Day 14: Extended Polymerization ---
 *              --- Part One ---
 *
 *   https://adventofcode.com/2021/day/14
 */

const steps = 10

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	diff := doPart1()
	log.Println("the quantity of the most common element subtracted from the quantity of the least common element is", diff)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	polymerTemplate := scanner.Text()

	scanner.Scan() // skip blank line

	rules := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	for step := 0; step < steps; step++ {
		nInsertions := 0
		insertions := make([]string, len(polymerTemplate))
		for i := 1; i < len(polymerTemplate); i++ {
			pair := polymerTemplate[i-1 : i+1]
			element, ok := rules[pair]
			if ok {
				insertions[i] = element
				nInsertions++
			}
		}

		elements := make([]byte, len(polymerTemplate)+nInsertions)
		for i, j := 0, 0; i < len(insertions); i++ {
			if insertions[i] != "" {
				elements[i+j] = byte(insertions[i][0])
				j++
			}
			elements[i+j] = polymerTemplate[i]
		}
		polymerTemplate = string(elements)
	}

	nLeastCommon, nMostCommon := math.MaxInt64, 0
	elementCounts := make(map[rune]int)
	for _, r := range polymerTemplate {
		elementCounts[r]++
	}
	for _, count := range elementCounts {
		if count < nLeastCommon {
			nLeastCommon = count
		}
		if count > nMostCommon {
			nMostCommon = count
		}
	}
	return nMostCommon - nLeastCommon
}
