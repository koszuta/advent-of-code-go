package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

/*
 *   --- Day 10: Syntax Scoring ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/10#part2
 */

var (
	unmatchedCharScores map[rune]int

	openingChars = "([{<"
	closingChars = ")]}>"
)

func init() {
	unmatchedCharScores = make(map[rune]int)
	unmatchedCharScores[')'] = 1
	unmatchedCharScores[']'] = 2
	unmatchedCharScores['}'] = 3
	unmatchedCharScores['>'] = 4
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	score := doPart2()
	log.Println("the middle score is", score)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scores := make([]int, 0)
	for scanner.Scan() {
		char, unmatchedChars := checkChunks(scanner.Text())
		if char == 0 {
			score := 0
			for i := len(unmatchedChars) - 1; i >= 0; i-- {
				score *= 5
				openingChar := unmatchedChars[i]
				score += unmatchedCharScores[getClosingChar(openingChar)]
			}
			scores = append(scores, score)
		}
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	return scores[len(scores)/2]
}

func checkChunks(line string) (rune, []rune) {
	var charToMatch rune
	chars := make([]rune, 0)
	for _, r := range line {
		if strings.ContainsRune(openingChars, r) {
			chars = append(chars, r)

		} else if strings.ContainsRune(closingChars, r) {
			charToMatch, chars = chars[len(chars)-1], chars[:len(chars)-1]
			expectedOpening := getOpeningChar(r)
			if charToMatch != expectedOpening {
				return r, nil
			}
		} else {
			log.Panicln("unhandled char", string(r))
		}
	}
	return 0, chars
}

func getOpeningChar(r rune) rune {
	return rune(openingChars[strings.IndexRune(closingChars, r)])
}

func getClosingChar(r rune) rune {
	return rune(closingChars[strings.IndexRune(openingChars, r)])
}
