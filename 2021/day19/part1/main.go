package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 0

/*
 *   --- Day 19:  ---
 *      --- Part One ---
 *
 *   https://adventofcode.com/2021/day/19
 */

type Scanner []Beacon

type Beacon struct {
	x, y, z int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	x := doPart1()
	log.Println(x)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scan := bufio.NewScanner(file)

	scanners := make([]Scanner, 0)

	var scanner Scanner
	headerRegex := regexp.MustCompile(`--- scanner \d+ ---`)
	for scan.Scan() {
		line := scan.Text()
		if headerRegex.Match([]byte(line)) {
			scanner = make(Scanner, 0)
		} else if line != "" {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			scanner = append(scanner, Beacon{x, y, z})
		} else {
			scanners = append(scanners, scanner)
		}
	}
	scanners = append(scanners, scanner)

	for _, scanner := range scanners {
		log.Println(scanner)
	}

	return 0
}
