package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
 *   --- Day 7: Internet Protocol Version 7 ---
 *                --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/7#part2
 */

var hypernetSeqRex *regexp.Regexp

func init() {
	hypernetSeqRex = regexp.MustCompile(`\[([^\[\]]+)\]`)
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	nIPsSupportSSL := 0
	scanner := bufio.NewScanner(file)
OUTER:
	for scanner.Scan() {
		line := scanner.Text()

		// Find all BABs in each hypernet sequences and invert them
		invertedHypernetSeqBABs := make([]string, 0)
		for _, hns := range hypernetSeqRex.FindAllStringSubmatch(line, -1) {
			for _, bab := range findABAs(hns[1]) {
				invertedHypernetSeqBABs = append(invertedHypernetSeqBABs, invertABA(bab))
			}
			// Remove hypernet sequence from the line
			line = strings.ReplaceAll(line, hns[0], "~~")
		}

		// Find all ABAs in supernet sequences
		supernetSeqABAs := findABAs(line)

		// Count how many ABAs match inverted BABs
		for _, aba := range supernetSeqABAs {
			for _, invertedBAB := range invertedHypernetSeqBABs {
				if aba == invertedBAB {
					nIPsSupportSSL++
					continue OUTER
				}
			}
		}
	}

	log.Printf("the number of IPs that support SSL is %d\n", nIPsSupportSSL)
}

func findABAs(s string) []string {
	abas := make([]string, 0)
	for i := 2; i < len(s); i++ {
		if s[i] != s[i-1] && s[i] == s[i-2] {
			abas = append(abas, s[i-2:i+1])
		}
	}
	return abas
}

func invertABA(aba string) (bab string) {
	return string(aba[1]) + string(aba[0]) + string(aba[1])
}
