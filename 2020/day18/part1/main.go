package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 18: Operation Order ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/18
 */

var expr = "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"

func eval(expr string) int {
	parts := strings.Split(expr, " ")

	total, _ := strconv.Atoi(parts[0])

	for i := 1; i < len(parts)-1; i += 2 {
		operator := parts[i]
		num, _ := strconv.Atoi(parts[i+1])

		switch operator {
		case "+":
			total += num
		case "*":
			total *= num
		}
	}

	return total
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		expr := scanner.Text()
		i := 0
		for strings.Contains(expr, "(") {
			for j, c := range expr {
				if c == '(' {
					i = j
				}
				if c == ')' {
					subExpr := expr[i+1 : j]
					// log.Println(subExpr)
					total := eval(subExpr)
					strTotal := strconv.Itoa(total)
					expr = expr[:i] + strTotal + expr[j+1:]
					// log.Println(expr)
					// log.Println(total)
					break
				}
			}
		}
		total := eval(expr)
		// log.Println(total)
		sum += total
	}
	log.Println("when + and * have the same precedence, the sum of the results is", sum)
}
