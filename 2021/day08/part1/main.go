package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

/*
 *   --- Day 8: Seven Segment Search ---
 *            --- Part One ---
 *
 *   https://adventofcode.com/2021/day/8
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	n := doPart1()
	log.Println("the digits 1, 4, 7, or 8 appear", n, "times in the output")
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	n := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		outputs := strings.Split(parts[1], " ")
		for _, segments := range outputs {
			l := len(segments)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				n++
			}
		}
	}

	return n
}
