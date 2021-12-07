package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

/*
 *   --- Day 3: Binary Diagnostic ---
 *           --- Part One ---
 *
 *   https://adventofcode.com/2021/day/3
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	g, e := doPart1()
	log.Println("gamma rate:", g)
	log.Println("epsilon rate:", e)
	log.Println("power consumption:", g*e)
}

func doPart1() (int64, int64) {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	zeros := make(map[int]int)
	ones := make(map[int]int)

	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count = len(line)
		for i, c := range line {
			if c == '0' {
				zeros[i]++
			} else {
				ones[i]++
			}
		}
	}

	gamma, eplison := "", ""
	for i := 0; i < count; i++ {
		if zeros[i] > ones[i] {
			gamma += "0"
			eplison += "1"
		} else {
			gamma += "1"
			eplison += "0"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(eplison, 2, 64)
	return g, e
}
