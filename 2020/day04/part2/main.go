package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
 *   --- Day 4: Passport Processing ---
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2020/day/4#part2
 */

/*
 * Field requirements:
 *   byr (Birth Year) - four digits; at least 1920 and at most 2002.
 *   iyr (Issue Year) - four digits; at least 2010 and at most 2020.
 *   eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
 *   hgt (Height) - a number followed by either cm or in:
 *     If cm, the number must be at least 150 and at most 193.
 *     If in, the number must be at least 59 and at most 76.
 *   hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
 *   ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
 *   pid (Passport ID) - a nine-digit number, including leading zeroes.
 *   cid (Country ID) - ignored, missing or not.
 */
func fieldsAreValid(fields map[string]string) bool {
	return regexp.MustCompile("^(19[2-9]\\d|200[0-2])$").MatchString(fields["byr"]) &&
		regexp.MustCompile("^20(1\\d|20)$").MatchString(fields["iyr"]) &&
		regexp.MustCompile("^20(2\\d|30)$").MatchString(fields["eyr"]) &&
		regexp.MustCompile("^(1([5-8]\\d|9[0-3])cm|(59|6\\d|7[0-6])in)$").MatchString(fields["hgt"]) &&
		regexp.MustCompile("^#[\\da-f]{6}$").MatchString(fields["hcl"]) &&
		regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$").MatchString(fields["ecl"]) &&
		regexp.MustCompile("^\\d{9}$").MatchString(fields["pid"])
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	validPassports := 0
	fields := make(map[string]string)
	for scanner.Scan() {
		text := scanner.Text()

		// Passport fields can span multiple lines
		// Individual passports are separated by an empty line
		if text == "" {
			// Check if all the required fields are valid
			if fieldsAreValid(fields) {
				validPassports++
			}

			// Reset the fields for the next passport
			fields = make(map[string]string)

		} else {
			// Fields are separated by spaces
			parts := strings.Split(text, " ")

			// Parse the fields as key:value pairs
			for _, part := range parts {
				field := strings.Split(part, ":")
				fields[field[0]] = field[1]
			}
		}
	}

	log.Printf("%d passports are valid\n", validPassports)
}
