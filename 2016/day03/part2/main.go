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
 *              --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/3#part2
 */

var spaces *regexp.Regexp

func init() {
	spaces = regexp.MustCompile(`\s+`)
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	nValid := 0
	hasLines := true
	scanner := bufio.NewScanner(file)
OUT:
	for hasLines {
		var triangles [3][3]int
		for i := 0; i < 3; i++ {
			if !scanner.Scan() {
				break OUT
			}
			line := scanner.Text()
			sides := parseLine(line)
			for j, side := range sides {
				triangles[j][i] = side
			}
		}
		for _, t := range triangles {
			if isTriangleValid(t[0], t[1], t[2]) {
				nValid++
			}
		}
	}

	log.Printf("the number of valid triangles is %d\n", nValid)
}

func parseLine(line string) [3]int {
	sides := spaces.Split(strings.TrimSpace(line), -1)
	s11, _ := strconv.Atoi(sides[0])
	s12, _ := strconv.Atoi(sides[1])
	s13, _ := strconv.Atoi(sides[2])
	return [3]int{s11, s12, s13}
}

func isTriangleValid(s1, s2, s3 int) bool {
	return s1+s2 > s3 && s1+s3 > s2 && s2+s3 > s1
}
