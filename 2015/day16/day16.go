package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sue struct {
	number, children, cats, samoyeds, pomeranians, akitas, vizslas, goldfish, tree, cars, perfumes, matches int
}

func compare(s1, s2 sue) int {
	matches := 0
	if s2.children != -1 && s1.children == s2.children {
		matches++
	}
	if s2.cats != -1 && s1.cats < s2.cats {
		matches++
	}
	if s2.samoyeds != -1 && s1.samoyeds == s2.samoyeds {
		matches++
	}
	if s2.pomeranians != -1 && s1.pomeranians > s2.pomeranians {
		matches++
	}
	if s2.akitas != -1 && s1.akitas == s2.akitas {
		matches++
	}
	if s2.vizslas != -1 && s1.vizslas == s2.vizslas {
		matches++
	}
	if s2.goldfish != -1 && s1.goldfish > s2.goldfish {
		matches++
	}
	if s2.tree != -1 && s1.tree < s2.tree {
		matches++
	}
	if s2.cars != -1 && s1.cars == s2.cars {
		matches++
	}
	if s2.perfumes != -1 && s1.perfumes == s2.perfumes {
		matches++
	}
	return matches
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	giftSue := sue { 0, 3, 7, 2, 3, 0, 0, 5, 3, 2, 1, 0 }

	sues := make([]sue, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ":", "")
		parts := strings.Split(line, " ")
		sue := sue { -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0 }
		for i := 0; i < len(parts); i += 2 {
			switch parts[i] {
			case "Sue":
				number, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.number = number
			case "children":
				children, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.children = children
			case "cats":
				cats, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.cats = cats
			case "samoyeds":
				samoyeds, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.samoyeds = samoyeds
			case "pomeranians":
				pomeranians, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.pomeranians = pomeranians
			case "akitas":
				akitas, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.akitas = akitas
			case "vizslas":
				vizslas, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.vizslas = vizslas
			case "goldfish":
				goldfish, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.goldfish = goldfish
			case "tree":
				tree, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.tree = tree
			case "cars":
				cars, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.cars = cars
			case "perfumes":
				perfumes, err := strconv.Atoi(parts[i+1])
				if err != nil {
					panic(err)
				}
				sue.perfumes = perfumes
			}
		}
		sue.matches = compare(giftSue, sue)
		sues = append(sues, sue)
	}

	highestMatch := 0
	for _, sue := range sues {
		if sue.matches >= highestMatch {
			highestMatch = sue.matches
		}
	}
	for _, sue := range sues {
		if sue.matches == highestMatch {
			fmt.Printf("sue=%d\n", sue)
		}
	}
}
