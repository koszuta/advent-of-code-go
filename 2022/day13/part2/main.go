package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	Divider1 = "[[2]]"
	Divider2 = "[[6]]"
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
	// Add the extra divider packets to the input
	lines = append(lines, Divider1)
	lines = append(lines, Divider2)

	packets := make([][]any, 0, (len(lines)+2)*2/3)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var packet []any
		_ = json.Unmarshal([]byte(line), &packet)
		packets = append(packets, packet)
	}

	sort.Slice(packets, func(i, j int) bool {
		return Compare(packets[i], packets[j]) < 0
	})

	decoderKey := 1
	for i, packet := range packets {
		p, _ := json.Marshal(packet)
		switch string(p) {
		case Divider1:
			decoderKey *= i + 1
		case Divider2:
			decoderKey *= i + 1
		}
	}
	log.Println("decoder key:", decoderKey)
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
