package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	cycle = 1
	regX  = 1

	targetCycles = map[int]struct{}{
		20:  {},
		60:  {},
		100: {},
		140: {},
		180: {},
		220: {},
	}
)

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	var sum int
	doAddSignalStrength := func() {
		if _, found := targetCycles[cycle]; found {
			log.Printf("cycle: %d,\tregister X: %d,\tsignal strength: %d\n", cycle, regX, cycle*regX)
			sum += cycle * regX
		}
	}

	for _, instr := range lines {
		instr = strings.TrimSpace(instr) // sanitize CRLF

		switch {
		case instr == "noop":
		case strings.HasPrefix(instr, "addx"):
			cycle++
			doAddSignalStrength()

			rest := strings.TrimPrefix(instr, "addx ")
			v, _ := strconv.Atoi(rest)
			regX += v
		}
		cycle++
		doAddSignalStrength()
	}
	log.Println("sum of the six signal strengths:", sum)
}
