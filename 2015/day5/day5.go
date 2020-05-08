package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numNiceStrings := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numVowels := 0
		containsDoubleLetter := false
		var lastLetter rune = -1
		if !strings.Contains(line, "ab") && !strings.Contains(line, "cd") && !strings.Contains(line, "pq") && !strings.Contains(line, "xy") {
			for _, c := range line {
				if strings.ContainsRune("aeiou", c) {
					numVowels++
				}
				if lastLetter != -1 && c == lastLetter {
					containsDoubleLetter = true
				}
				lastLetter = c
			}
			if containsDoubleLetter && numVowels >= 3 {
				numNiceStrings++
			}
		}
	}

	fmt.Printf("numNiceStrings=%d\n", numNiceStrings)

	numNiceStringsNew := 0
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
SCANNING:
	for scanner.Scan() {
		line := scanner.Text()
		containsDoublePair := false
	DOUBLE_PAIR_CHECK:
		for i := 0; i+3 < len(line); i++ {
			pair := line[i:i+2]
			for j := i+2; j+1 < len(line); j++ {
				if reflect.DeepEqual(pair, line[j:j+2]) {
					containsDoublePair = true
					break DOUBLE_PAIR_CHECK
				}
			}
		}
		if !containsDoublePair {
			continue SCANNING
		}
		for i := 0; i+2 < len(line); i++ {
			if line[i] == line[i+2] {
				numNiceStringsNew++
				continue SCANNING
			}
		}
	}

	fmt.Printf("numNiceStringsNew=%d\n", numNiceStringsNew)
}
