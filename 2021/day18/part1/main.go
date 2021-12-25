package main

import (
	"advent-of-code-go/2021/day18/sfn"
	"bufio"
	"log"
	"os"
	"time"
)

const expectedResult = 3675

/*
 *   --- Day 18: Snailfish ---
 *       --- Part One ---
 *
 *   https://adventofcode.com/2021/day/18
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	mag := doPart1()
	log.Println("the magnitude of the final sum is", mag)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sum *sfn.SnailfishNumber
	for scanner.Scan() {
		sfn := sfn.ParseSnailfishNumber(scanner.Text(), nil)
		if sum != nil {
			sum = sum.Add(sfn)
		} else {
			sum = sfn
		}
	}

	return sum.Magnitude()
}
