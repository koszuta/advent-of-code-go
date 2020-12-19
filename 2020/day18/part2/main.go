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
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/18#part2
 */

var expression = "5 + (8 * 3 + 9 + 3 * 4 * 3)"

func evalExpression(expr string) int {
	parts := strings.Split(expr, " ")

	toMultiply := make([]int, 0, 0)
	lastWasAdd := false
	// log.Println()
	for i := 1; i < len(parts)-1; i += 2 {
		// log.Println(toMultiply)
		// log.Println(i, lastWasAdd, parts[i], len(parts), i == len(parts)-2)
		if parts[i] == "+" {
			lastWasAdd = true
			sum, _ := strconv.Atoi(parts[i-1])
			for ; i < len(parts)-1 && parts[i] == "+"; i += 2 {
				// log.Println(sum)
				num, _ := strconv.Atoi(parts[i+1])
				sum += num
			}
			i -= 2
			// log.Println(sum)
			toMultiply = append(toMultiply, sum)
		} else {
			if i == 1 || !lastWasAdd {
				num, _ := strconv.Atoi(parts[i-1])
				toMultiply = append(toMultiply, num)
			}
			if i == len(parts)-2 {
				num, _ := strconv.Atoi(parts[i+1])
				toMultiply = append(toMultiply, num)
			}
			lastWasAdd = false
		}
	}
	// log.Println(toMultiply)

	total := 1
	for _, num := range toMultiply {
		total *= num
	}

	// log.Println(total)
	return total
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		expression := scanner.Text()
		i := 0
		for strings.Contains(expression, "(") {
			for j, c := range expression {
				if c == '(' {
					i = j
				}
				if c == ')' {
					subExpr := expression[i+1 : j]
					// log.Println(subExpr)
					total := evalExpression(subExpr)
					strTotal := strconv.Itoa(total)
					expression = expression[:i] + strTotal + expression[j+1:]
					// log.Println(expression)
					// log.Println(total)
					break
				}
			}
		}
		total := evalExpression(expression)
		// log.Println(total)
		sum += total
	}
	log.Println("when + has higher precedence than *, the sum of the results is", sum)
}
