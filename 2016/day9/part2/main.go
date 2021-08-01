package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
 *   --- Day 9: Explosives in Cyberspace ---
 *              --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/9#part2
 */

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	compressedData := string(file)

	log.Printf("the decompressed length is %d\n", decompressedLength(compressedData))
}

func decompressedLength(compressedData string) (decompressedLen int) {
	curr, prev := 0, 0
	// Move past opening parenthesis
	for curr < len(compressedData) && compressedData[curr] != '(' {
		curr++
	}
	uncompressedSection := string(compressedData[:curr])

	// If there's only uncompressed data, just return its length
	if curr == len(compressedData) {
		return len(uncompressedSection)
	}
	curr++
	prev = curr

	// Find index of closing parenthesis to get marker
	for curr < len(compressedData) && compressedData[curr] != ')' {
		curr++
	}
	marker := string(compressedData[prev:curr])
	curr++
	prev = curr

	// Parse marker to know how much data the marker affects
	a, b := parseMarker(marker)
	curr += a

	return len(uncompressedSection) + (decompressedLength(compressedData[prev:curr]) * b) + decompressedLength(compressedData[curr:])
}

func parseMarker(marker string) (a, b int) {
	parts := strings.Split(marker, "x")
	a, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Panicln(err)
	}
	b, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Panicln(err)
	}
	return a, b
}
