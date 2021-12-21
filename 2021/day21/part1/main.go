package main

import (
	"log"
	"time"
)

const expectedResult = 571032

/*
 *   --- Day 21: Dirac Dice ---
 *        --- Part One ---
 *
 *   https://adventofcode.com/2021/day/21
 */

const (
	p1StartingPos = 2
	p2StartingPos = 10
	winningScore  = 1000
)

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	prod := doPart1()
	log.Println("the score of the losing player multiplied by the number of times the die was rolled is", prod)
}

func doPart1() int {
	p1Position, p2Position := p1StartingPos-1, p2StartingPos-1
	p1Score, p2Score := 0, 0

	nRolls, die := 0, 1
	for {
		roll := 3 * (die + 1)
		die += 3
		nRolls += 3
		p1Position = (p1Position + roll) % 10
		p1Score += p1Position + 1
		if p1Score >= winningScore {
			break
		}

		roll = 3 * (die + 1)
		die += 3
		nRolls += 3
		p2Position = (p2Position + roll) % 10
		p2Score += p2Position + 1
		if p2Score >= winningScore {
			break
		}
	}

	if p1Score < winningScore {
		return p1Score * nRolls
	}
	return p2Score * nRolls
}
