package main

import (
	"advent-of-code-go/2016/day21/operation"
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

/*
 *   --- Day 21: Scrambled Letters and Hash ---
 *                --- Part One ---
 *
 *   https://adventofcode.com/2016/day/21
 */

const password = "abcdefgh"

var (
	operations []operation.Operation
)

func init() {
	swapPosRex := regexp.MustCompile(`swap position (\d+) with position (\d+)`)
	swapLetRex := regexp.MustCompile(`swap letter (\w+) with letter (\w+)`)
	rotLeftRex := regexp.MustCompile(`rotate left (\d+) step`)
	rotRightRex := regexp.MustCompile(`rotate right (\d+) step`)
	rotLetPosRex := regexp.MustCompile(`rotate based on position of letter (\w+)`)
	revSpanRex := regexp.MustCompile(`reverse positions (\d+) through (\d+)`)
	moveRex := regexp.MustCompile(`move position (\d+) to position (\d+)`)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var op operation.Operation
		match := swapPosRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			x, err := strconv.Atoi(match[0][1])
			if err != nil {
				log.Panicln(err)
			}
			y, err := strconv.Atoi(match[0][2])
			if err != nil {
				log.Panicln(err)
			}
			op = &operation.SwapPositionsOp{X: x, Y: y}
		}
		match = swapLetRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			op = &operation.SwapLettersOp{X: rune(match[0][1][0]), Y: rune(match[0][2][0])}
		}
		match = rotLeftRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			x, err := strconv.Atoi(match[0][1])
			if err != nil {
				log.Panicln(err)
			}
			op = &operation.RotateLeftOp{X: x}
		}
		match = rotRightRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			x, err := strconv.Atoi(match[0][1])
			if err != nil {
				log.Panicln(err)
			}
			op = &operation.RotateRightOp{X: x}
		}
		match = rotLetPosRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			op = &operation.RotateOnLetterOp{X: rune(match[0][1][0])}
		}
		match = revSpanRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			x, err := strconv.Atoi(match[0][1])
			if err != nil {
				log.Panicln(err)
			}
			y, err := strconv.Atoi(match[0][2])
			if err != nil {
				log.Panicln(err)
			}
			op = &operation.ReverseSpanOp{X: x, Y: y}
		}
		match = moveRex.FindAllStringSubmatch(line, -1)
		if match != nil {
			x, err := strconv.Atoi(match[0][1])
			if err != nil {
				log.Panicln(err)
			}
			y, err := strconv.Atoi(match[0][2])
			if err != nil {
				log.Panicln(err)
			}
			op = &operation.MoveOp{X: x, Y: y}
		}
		operations = append(operations, op)
	}
}

func main() {
	scrambledPW := []rune(password)
	for _, op := range operations {
		scrambledPW = op.Execute(scrambledPW)
	}
	log.Printf("scrambling %s gives you %s\n", password, string(scrambledPW))
}
