package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
 *   --- Day 9: Explosives in Cyberspace ---
 *              --- Part One ---
 *
 *   https://adventofcode.com/2016/day/9
 */

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	compressedData := string(file)

	decompressedData := ""
	curr, prev := 0, 0
	for curr < len(compressedData) {
		// Move past opening parenthesis
		curr++
		prev = curr

		// Find index of closing parenthesis to get marker
		for curr < len(compressedData) && compressedData[curr] != ')' {
			curr++
		}
		marker := compressedData[prev:curr]
		curr++
		prev = curr

		// Parse marker and add repeated data
		a, b := parseMarker(marker)
		data := compressedData[curr : curr+a]
		decompressedData += strings.Repeat(data, b)
		curr += a
		prev = curr
	}

	// log.Printf("the decompressed data is %s\n", decompressedData)
	log.Printf("the decompressed length is %d\n", len(decompressedData))
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
