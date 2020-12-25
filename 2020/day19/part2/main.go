package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type rule struct {
	symbol    string
	terminals []string
	subRules  []subRule
}

type subRule struct {
	symbols []string
}

func index(x, y, z, n int) int {
	return z*n*n + y*n + x
}

// CYK Algorithm: https://en.wikipedia.org/wiki/CYK_algorithm#As_pseudocode
func cyk(I string, R []rule, indexOfSymbol map[string]int) bool {
	n := len(I)
	r := len(R)

	size := n * n * r
	P := make([]bool, size, size)

	for s, as := range I {
		for v, Rv := range R {
			for _, terminal := range Rv.terminals {
				if terminal == string(as) {
					P[index(0, s, v, n)] = true
				}
			}
		}
	}

	for l := 2; l <= n; l++ {
		for s := 1; s <= n-l+1; s++ {
			for p := 1; p <= l-1; p++ {
				for a, Ra := range R {
					for _, sub := range Ra.subRules {
						Rb := sub.symbols[0]
						Rc := sub.symbols[1]
						b := indexOfSymbol[Rb]
						c := indexOfSymbol[Rc]

						if P[index(p-1, s-1, b, n)] && P[index(l-p-1, s+p-1, c, n)] {
							P[index(l-1, s-1, a, n)] = true
						}
					}
				}
			}
		}
	}

	return P[index(n-1, 0, 0, n)]
}

func parseRule(line string) rule {
	parts := strings.Split(line, " ")

	symbol := parts[0][:len(parts[0])-1]
	newRule := rule{}
	newRule.symbol = symbol

	sub := subRule{}
	sub.symbols = make([]string, 0, 0)
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		if part == "|" {
			if len(sub.symbols) != 2 && len(sub.symbols) != 0 {
				log.Panicln("a production must have zero or two variables", sub.symbols)
			}
			if len(sub.symbols) > 0 {
				newRule.subRules = append(newRule.subRules, sub)
				sub = subRule{}
				sub.symbols = make([]string, 0, 0)
			}

		} else if part[0] == '"' {
			t := part[1 : len(part)-1]
			newRule.terminals = append(newRule.terminals, t)

		} else {
			symbol := part
			sub.symbols = append(sub.symbols, symbol)
		}
	}
	if len(sub.symbols) > 0 {
		newRule.subRules = append(newRule.subRules, sub)
	}
	return newRule
}

func main() {
	// Puzzle input
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// Parse grammar; must be in Chomsky Normal Form
	rules := make([]rule, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rules = append(rules, parseRule(line))
	}
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].symbol < rules[j].symbol
	})

	// Build map of symbol indexes
	indexOfSymbol := make(map[string]int)
	for i, r := range rules {
		indexOfSymbol[r.symbol] = i
	}

	// Check if each message can be generated from the grammar
	validMessageCount := 0
	for scanner.Scan() {
		message := scanner.Text()

		inGrammar := cyk(message, rules, indexOfSymbol)

		if inGrammar {
			validMessageCount++
		}
	}

	log.Println("after replace rules 8 and 11 with:")
	log.Println("8: 42 | 42 8")
	log.Println("11: 42 31 | 42 11 31")
	log.Println(validMessageCount, "messages match rule 0")
}
