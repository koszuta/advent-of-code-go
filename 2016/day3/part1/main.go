package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 *   --- Day 3: Squares With Three Sides ---
 *              --- Part One ---
 *
 *   https://adventofcode.com/2016/day/3
 */

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	nValid := 0
	spaces := regexp.MustCompile(`\s+`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sides := spaces.Split(strings.TrimSpace(line), -1)
		s1, _ := strconv.Atoi(sides[0])
		s2, _ := strconv.Atoi(sides[1])
		s3, _ := strconv.Atoi(sides[2])
		if isTriangleValid(s1, s2, s3) {
			nValid++
		}
	}

	log.Printf("the number of valid triangles is %d\n", nValid)
}

func isTriangleValid(s1, s2, s3 int) bool {
	return s1+s2 > s3 && s1+s3 > s2 && s2+s3 > s1
}
