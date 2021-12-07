package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

/*
 *   --- Day 1: Sonar Sweep ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/1#part2
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nIncreases := doPart2()
	log.Println("the number of measurements larger than the previous measurement is", nIncreases)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	depths := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Panicln(err)
		}
		depths = append(depths, depth)
	}

	nIncreases := 0
	prevSum := depths[0] + depths[1] + depths[2]
	for i := 3; i < len(depths); i++ {
		newSum := prevSum - depths[i-3] + depths[i]
		if newSum > prevSum {
			nIncreases++
		}
		prevSum = newSum
	}

	return nIncreases
}
