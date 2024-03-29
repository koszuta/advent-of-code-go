package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

const expectedResult = 0

/*
 *   --- Day 19:  ---
 *      --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/19#part2
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	x := doPart2()
	log.Println(x)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		log.Println(line)
	}

	return 0
}
