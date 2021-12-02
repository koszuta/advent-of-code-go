package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
 *   --- Day 1: Sonar Sweep ---
 *        --- Part One ---
 *
 *   https://adventofcode.com/2021/day/1
 */

func main() {
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
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			nIncreases++
		}
	}
	log.Println("the number of measurements larger than the previous measurement is", nIncreases)
}
