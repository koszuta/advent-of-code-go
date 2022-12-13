package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var lines []string

func init() {
	b, _ := os.ReadFile("../input.txt")
	lines = strings.Split(string(b), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line) // sanitize CRLF
	}
}

func main() {
	var sum int
	for i := 0; i < len(lines); i += 3 {
		var p1, p2 []any
		_ = json.Unmarshal([]byte(lines[i]), &p1)
		_ = json.Unmarshal([]byte(lines[i+1]), &p2)

		if Compare(p1, p2) <= 0 {
			sum += i/3 + 1
		}
	}

	log.Println("sum of the indexes of in order pairs:", sum)
}

func Compare(left, right any) int {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64: // [number, number]
			return int(left.(float64) - right.(float64))

		default: // [number, array]
			return Compare([]any{left}, right)
		}
	default:
		switch right.(type) {
		case float64: // [array, number]
			return Compare(left, []any{right})

		default: // [array, array]
			left, right := left.([]any), right.([]any)
			for i := 0; i < len(left) && i < len(right); i++ {
				if diff := Compare(left[i], right[i]); diff != 0 {
					return diff
				}
			}
			return len(left) - len(right)
		}
	}
}
