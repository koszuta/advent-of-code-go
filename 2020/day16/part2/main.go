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
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/16#part2
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

func getPossibleIndexFields(values []int) []string {
	possibleFields := make([]string, 0, 0)
OUTER:
	for field, rule := range fieldRules {
		for _, v := range values {
			if !rule.isValid(v) {
				continue OUTER
			}
		}
		possibleFields = append(possibleFields, field)
	}
	return possibleFields
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

	validTickets := make([][]int, 0, 0)
	{
	TO_NEXT_TICKET:
		for _, ticket := range nearbyTickets {
			for _, value := range ticket {
				if !validForSomeRule(value) {
					continue TO_NEXT_TICKET
				}
			}
			validTickets = append(validTickets, ticket)
		}
	}

	possibleIndexFields := make(map[int]map[string]struct{})
	{
		for j := 0; j < len(validTickets[0]); j++ {
			possibleIndexFields[j] = make(map[string]struct{})

			values := make([]int, 0, 0)
			for _, validTicket := range validTickets {
				values = append(values, validTicket[j])
			}

			for _, field := range getPossibleIndexFields(values) {
				possibleIndexFields[j][field] = struct{}{}
			}
		}
	}

	hasAmbiguity := true
	for hasAmbiguity {
		for i, fields := range possibleIndexFields {
			if len(fields) == 1 {
				for field := range fields {
					hasAmbiguity = false
					for j, fields := range possibleIndexFields {
						if i != j {
							delete(fields, field)
						}
						if len(fields) > 1 {
							hasAmbiguity = true
						}
					}
				}
			}
		}
	}

	// Calculate the product for the final puzzle solution
	prod := 1
	for i, rule := range possibleIndexFields {
		for ruleName := range rule {
			if strings.HasPrefix(ruleName, "departure") {
				prod *= yourTicket[i]
			}
		}
	}
	log.Println("the product of fields that start with the word departure is", prod)
}
