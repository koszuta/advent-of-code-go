package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const rounds = 10000

type Monkey struct {
	items           []int
	op, opVal       string
	test            int
	onTrue, onFalse int
}

func main() {
	var monkeys []*Monkey
	trick := 1
	{
		b, _ := os.ReadFile("../input.txt")
		lines := strings.Split(string(b), "\n")
		for i, line := range lines {
			lines[i] = strings.TrimSpace(line) // sanitize CRLF
		}

		monkeys = make([]*Monkey, 0, (len(lines)+1)/7)
		for i := 1; i < len(lines); i += 7 {
			var monkey Monkey

			startingItems := strings.Split(lines[i][16:], ", ")
			items := make([]int, len(startingItems))
			for i, v := range startingItems {
				worryLevel, _ := strconv.Atoi(v)
				items[i] = worryLevel
			}
			monkey.items = items

			operationParts := strings.Split(lines[i+1][21:], " ")
			monkey.op = operationParts[0]
			monkey.opVal = operationParts[1]

			mod, _ := strconv.Atoi(lines[i+2][19:])
			monkey1, _ := strconv.Atoi(lines[i+3][25:])
			monkey2, _ := strconv.Atoi(lines[i+4][26:])
			monkey.test = mod
			monkey.onTrue = monkey1
			monkey.onFalse = monkey2

			trick *= mod

			monkeys = append(monkeys, &monkey)
		}
	}

	inspections := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for i, monkey := range monkeys {
			// Increase the number of inspected items
			inspections[i] += len(monkey.items)

			for _, worryLevel := range monkey.items {
				// Execute operation
				val, err := strconv.Atoi(monkey.opVal)
				switch monkey.op {
				case "+":
					worryLevel += val
				case "*":
					if err != nil {
						val = worryLevel
					}
					worryLevel *= val
				}

				// Magic to prevent overflow
				worryLevel = worryLevel % trick

				// Throw the item base on the modulus test
				var throwTo *Monkey
				if worryLevel%monkey.test == 0 {
					throwTo = monkeys[monkey.onTrue]
				} else {
					throwTo = monkeys[monkey.onFalse]
				}
				throwTo.items = append(throwTo.items, worryLevel)
			}
			monkey.items = make([]int, 0)
		}
	}

	sort.Slice(inspections, func(i, j int) bool {
		return inspections[j] < inspections[i]
	})
	log.Println("level of monkey business:", inspections[0]*inspections[1])
}
