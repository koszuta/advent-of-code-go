package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 16: Ticket Translation ---
 *            --- Part One ---
 *
 *   https://adventofcode.com/2020/day/16
 */

var fieldRules map[string]ruleRange
var yourTicket []int
var nearbyTickets [][]int

type intRange struct {
	min, max int
}

type ruleRange struct {
	r1, r2 intRange
}

func (r *ruleRange) isValid(v int) bool {
	return (v >= r.r1.min && v <= r.r1.max) || (v >= r.r2.min && v <= r.r2.max)
}

func parseRule(line string) (string, ruleRange) {
	parts := strings.Split(line, ": ")

	field := parts[0]

	rule := ruleRange{}
	parts = strings.Split(parts[1], " or ")

	nums := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[1])
	rule.r1 = intRange{min, max}

	nums = strings.Split(parts[1], "-")
	min, _ = strconv.Atoi(nums[0])
	max, _ = strconv.Atoi(nums[1])
	rule.r2 = intRange{min, max}

	return field, rule
}

func parseTicket(line string) []int {
	ticket := make([]int, 0, 0)
	for _, n := range strings.Split(line, ",") {
		value, _ := strconv.Atoi(n)
		ticket = append(ticket, value)
	}
	return ticket
}

func validForSomeRule(v int) bool {
	for _, rule := range fieldRules {
		if rule.isValid(v) {
			return true
		}
	}
	return false
}

func init() {
	fieldRules = make(map[string]ruleRange)
	nearbyTickets = make([][]int, 0, 0)
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	parseMode := 0
	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "":
			continue
		case "your ticket:":
			parseMode = 1
			continue
		case "nearby tickets:":
			parseMode = 2
			continue
		}

		switch parseMode {
		case 0:
			field, ranges := parseRule(line)
			fieldRules[field] = ranges
		case 1:
			yourTicket = parseTicket(line)
		case 2:
			ticket := parseTicket(line)
			nearbyTickets = append(nearbyTickets, ticket)
		}
	}

	invalidCount := 0
	{
	CHECKING_TICKETS:
		for _, ticket := range nearbyTickets {
			for _, value := range ticket {
				if !validForSomeRule(value) {
					invalidCount += value
					continue CHECKING_TICKETS
				}
			}
		}
	}

	log.Println("the ticket scanning error rate is", invalidCount)
}
