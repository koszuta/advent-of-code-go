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
 *                --- Part One ---
 *
 *   https://adventofcode.com/2016/day/7
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

	nIPsSupportTLS := 0
	scanner := bufio.NewScanner(file)
OUTER:
	for scanner.Scan() {
		line := scanner.Text()

		// Check if any of the hypernet sequences in the line contains an ABBA
		hypernetSeqs := hypernetSeqRex.FindAllStringSubmatch(line, -1)
		for _, hns := range hypernetSeqs {
			if containsABBA(hns[1]) {
				continue OUTER
			}
			// Not sure whether it's faster to remove the hypernet sequence from the line
			line = strings.ReplaceAll(line, hns[0], "~~~")
		}

		// Since no hypernet sequence contains an ABBA, check if a supernet sequence does
		if containsABBA(line) {
			nIPsSupportTLS++
		}
	}

	log.Printf("the number of IPs that support TLS is %d\n", nIPsSupportTLS)
}

func containsABBA(s string) bool {
	for i := 3; i < len(s); i++ {
		if s[i] != s[i-1] && s[i] == s[i-3] && s[i-1] == s[i-2] {
			return true
		}
	}
	return false
}
