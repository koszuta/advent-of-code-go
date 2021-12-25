package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

const expectedResult = 424

/*
 *   --- Day 25: Sea Cucumber ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2021/day/25
 */

var m [][]rune

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nSteps := doPart1()
	log.Println("the first step in which no sea cucumbers move is", nSteps)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	m = make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0, len(line))
		for _, r := range line {
			row = append(row, r)
		}
		m = append(m, row)
	}

	var nSteps int
	var before [][]rune
	for nSteps = 0; !equal(before, m); nSteps++ {
		before = m
		moveEast()
		moveSouth()
	}
	return nSteps
}

func moveEast() {
	newM := make([][]rune, 0, len(m))
	for j, row := range m {
		newRow := make([]rune, len(row))
		newRow[0] = m[j][0]
		for i := len(row) - 1; i > 0; i-- {
			next := (i + 1) % len(row)
			if row[i] == '>' && row[next] == '.' {
				newRow[i] = '.'
				newRow[next] = '>'
			} else {
				newRow[i] = row[i]
			}
		}
		if row[0] == '>' && row[1] == '.' {
			newRow[0] = '.'
			newRow[1] = '>'
		}
		newM = append(newM, newRow)
	}
	m = newM
}

func moveSouth() {
	newM := make([][]rune, len(m))
	for j := range newM {
		newM[j] = make([]rune, len(m[0]))
		if j == 0 {
			for i := range newM[j] {
				newM[j][i] = m[j][i]
			}
		}
	}
	for i := range m[0] {
		for j := len(m) - 1; j > 0; j-- {
			next := (j + 1) % len(m)
			if m[j][i] == 'v' && m[next][i] == '.' {
				newM[j][i] = '.'
				newM[next][i] = 'v'
			} else {
				newM[j][i] = m[j][i]
			}
		}
		if m[0][i] == 'v' && m[1][i] == '.' {
			newM[0][i] = '.'
			newM[1][i] = 'v'
		}
	}
	m = newM
}

func equal(s1, s2 [][]rune) bool {
	if len(s1) != len(s2) || len(s1[0]) != len(s2[0]) {
		return false
	}
	for j, row := range s1 {
		for i := range row {
			if s1[j][i] != s2[j][i] {
				return false
			}
		}
	}
	return true
}
