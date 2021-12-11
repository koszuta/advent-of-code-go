package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

/*
 *   --- Day 10: Syntax Scoring ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2021/day/10
 */

var (
	illegalCharScores map[rune]int

	openingChars = "([{<"
	closingChars = ")]}>"
)

func init() {
	illegalCharScores = make(map[rune]int)
	illegalCharScores[')'] = 3
	illegalCharScores[']'] = 57
	illegalCharScores['}'] = 1197
	illegalCharScores['>'] = 25137
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	score := doPart1()
	log.Println("the total syntax error score is", score)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		char := checkChunks(scanner.Text())
		if char != 0 {
			sum += illegalCharScores[char]
		}
	}
	return sum
}

func checkChunks(line string) rune {
	var charToMatch rune
	chars := make([]rune, 0)
	for _, r := range line {
		if strings.ContainsRune(openingChars, r) {
			chars = append(chars, r)

		} else if strings.ContainsRune(closingChars, r) {
			charToMatch, chars = chars[len(chars)-1], chars[:len(chars)-1]
			expectedOpening := getOpeningChar(r)
			if charToMatch != expectedOpening {
				return r
			}
		} else {
			log.Panicln("unhandled char", string(r))
		}
	}
	return 0
}

func getOpeningChar(r rune) rune {
	return rune(openingChars[strings.IndexRune(closingChars, r)])
}
