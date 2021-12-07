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
 *     --- Part One ---
 *
 *   https://adventofcode.com/2021/day/2
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	pos, depth := doPart1()
	log.Println("pos:", pos)
	log.Println("depth:", depth)
	log.Println("prod:", pos*depth)
}

func doPart1() (int, int) {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	pos, depth := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			pos += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}

	return pos, depth
}
