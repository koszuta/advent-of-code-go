package main

import (
	"log"
	"time"
)

const expectedResult = 49975322685009

/*
 *   --- Day 21: Dirac Dice ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/21#part2
 */

const (
	p1StartingPos = 2
	p2StartingPos = 10
	winningScore  = 21
)

var (
	rolls map[int]int
	games map[game]pair
)

type pair struct {
	p1, p2 int
}

type game struct {
	p1Position, p2Position int
	p1Score, p2Score       int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nWins := doPart2()
	log.Println("the more winning player won in", nWins, "universes")
}

func doPart2() int {
	games = make(map[game]pair)
	rolls = map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	wins := playGame(p1StartingPos, p2StartingPos, 0, 0)

	if wins.p1 > wins.p2 {
		return wins.p1
	}
	return wins.p2
}

func playGame(p1, p2, s1, s2 int) pair {
	if s2 >= winningScore {
		return pair{0, 1}
	}
	if wins, found := games[game{p1, p2, s1, s2}]; found {
		return wins
	}

	wins := pair{}
	for roll, times := range rolls {
		newPosition := p1 + roll
		if newPosition > 10 {
			newPosition -= 10
		}
		newScore := newPosition + s1
		w := playGame(p2, newPosition, s2, newScore)

		wins.p1 += w.p2 * times
		wins.p2 += w.p1 * times
	}

	games[game{p1, p2, s1, s2}] = wins
	return wins
}
