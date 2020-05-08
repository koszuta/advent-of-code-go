package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	numChars := 0
	numCharsInMemory := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numChars += len(line)
		charsInMemory := 0
		for i := 1; i < len(line) - 1; i, charsInMemory = i+1, charsInMemory+1 {
			if line[i] == '\\' {
				switch line[i+1] {
				case '\\', '"':
					i++
				case 'x':
					i += 3
				}
			}
		}
		// fmt.Printf("line=%s %d %d\n", line, len(line), charsInMemory)
		numCharsInMemory += charsInMemory
	}

	fmt.Printf("difference=%d\n", numChars - numCharsInMemory)

	numChars = 0
	numEncodedChars := 0
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numChars += len(line)
		// fmt.Printf("line=%s -> ", line)
		line = strings.ReplaceAll(line, "\\", "\\\\")
		line = strings.ReplaceAll(line, "\"", "\\\"")
		line = "\"" + line + "\""
		// fmt.Printf("%s\n", line)
		numEncodedChars += len(line)
	}
	fmt.Printf("difference=%d\n", numEncodedChars - numChars)
}
