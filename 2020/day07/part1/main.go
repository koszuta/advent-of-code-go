package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 *   --- Day 7: Handy Haversacks ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2020/day/7
 */

const shinyGold = "shiny gold"

var bags map[string]map[string]int

func containsBag(bag, target string) bool {
	for child := range bags[bag] {
		if bag == target || containsBag(child, target) {
			return true
		}
	}
	return false
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	bags = make(map[string]map[string]int)
	for scanner.Scan() {
		line := scanner.Text()

		// Split each line into a bag and its child bags
		parts := strings.Split(line, " bags contain ")

		// Init the map of children for this bag
		bag := parts[0]
		children := parts[1]
		bags[bag] = make(map[string]int)

		// Check if the bag actually has any children
		if children != "no other bags." {
			// Normalize the pluralization of "bag" for simpler parsing
			re := regexp.MustCompile("bags?")
			children = re.ReplaceAllString(children, "bags")

			// Remove the ending and split up the rest of the line
			children = strings.ReplaceAll(children, " bags.", "")
			parts = strings.Split(children, " bags, ")

			for _, child := range parts {
				// Child bags have a count and color
				childParts := strings.SplitN(child, " ", 2)
				count, _ := strconv.Atoi(childParts[0])
				childColor := childParts[1]
				bags[bag][childColor] = count
			}
		}
	}

	count := 0
	for b := range bags {
		if b != shinyGold {
			if containsBag(b, shinyGold) {
				count++
			}
		}
	}
	log.Printf("%d bag colors contain at least one shiny gold bag\n", count)
}
