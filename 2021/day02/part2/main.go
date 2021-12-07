package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 *   --- Day 2: Dive! ---
 *     --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/2#part2
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	pos, depth := doPart2()
	log.Println("pos:", pos)
	log.Println("depth:", depth)
	log.Println("prod:", pos*depth)
}

func doPart2() (int, int) {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	pos, depth, aim := 0, 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			pos += val
			depth += val * aim
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	return pos, depth
}
