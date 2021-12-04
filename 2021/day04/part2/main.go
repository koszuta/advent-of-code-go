package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 *   --- Day 4: Giant Squid ---
 *        --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/4#part2
 */

type bingoBoard [5][5]*bingoSpace

type bingoSpace struct {
	number int
	marked bool
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	calledNums := make([]int, 0)
	scanner.Scan()
	line := scanner.Text()
	for _, n := range strings.Split(line, ",") {
		num, _ := strconv.ParseInt(n, 10, 64)
		calledNums = append(calledNums, int(num))
	}

	scanner.Scan() // skip blank line between

	whitespace := regexp.MustCompile(`\s+`)
	boards := make([]bingoBoard, 0)
	var board bingoBoard
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, board)
			board = bingoBoard{}
			j = 0
		} else {
			i := 0
			for _, n := range whitespace.Split(line, -1) {
				if n == "" {
					continue // leading whitespace causes blank value we can skip
				}
				num, _ := strconv.ParseInt(n, 10, 64)
				space := bingoSpace{number: int(num)}
				board[j][i] = &space
				i++
			}
			j++
		}
	}

	for _, calledNum := range calledNums {
		for i := 0; i < len(boards); {
			if marked := markNumber(calledNum, boards[i]); marked {
				if boardWon(boards[i]) {
					if len(boards) == 1 {
						log.Println("score:", calledNum*scoreBoard(boards[0]))
						os.Exit(0)
					}
					boards = append(boards[:i], boards[i+1:]...)
					continue
				}
			}
			i++
		}
	}
	log.Panicln("couldn't find a winning board \u200d")
}

func markNumber(calledNum int, board bingoBoard) bool {
	for j, row := range board {
		for i, space := range row {
			if space.number == calledNum {
				board[j][i].marked = true
				return true
			}
		}
	}
	return false
}

func boardWon(board bingoBoard) bool {
ACROSS:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if !board[j][i].marked {
				continue ACROSS
			}
		}
		log.Println("won across on row", j)
		printBoard(board)
		return true
	}
DOWN:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board[j][i].marked {
				continue DOWN
			}
		}
		log.Println("won down on column", i)
		printBoard(board)
		return true
	}
	return false
}

func scoreBoard(board bingoBoard) (score int) {
	for _, row := range board {
		for _, space := range row {
			if !space.marked {
				score += space.number
			}
		}
	}
	return
}

func printBoard(board bingoBoard) {
	for _, row := range board {
		for _, space := range row {
			fmt.Printf("%v ", *space)
		}
		fmt.Println()
	}
	fmt.Println()
}
