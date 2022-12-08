package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	var count int
	for _, line := range lines {
		line = strings.TrimSpace(line) // sanitize CRLF

		var pair [2]Range
		for i, part := range strings.Split(line, ",") {
			sections := strings.Split(part, "-")
			start, _ := strconv.Atoi(sections[0])
			end, _ := strconv.Atoi(sections[1])
			pair[i] = Range{Start: start, End: end}
		}

		if pair[0].Overlaps(pair[1]) {
			count++
		}
	}

	fmt.Println("overlapping ranges count:", count)
}

func (q *Range) Overlaps(r Range) bool {
	return q.Includes(r.Start) || q.Includes(r.End) || r.Includes(q.Start) || r.Includes(q.End)
}

func (r *Range) Includes(v int) bool {
	return v >= r.Start && v <= r.End
}
