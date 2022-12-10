package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	cycle = 1
	regX  = 1
)

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	buf := make([]rune, 40)
	doDrawPixel := func() {
		if cycle-regX >= 0 && cycle-regX < 3 {
			buf[cycle-1] = '█'
		} else {
			buf[cycle-1] = ' '
		}
	}
	maybeDrawRow := func() {
		if cycle > 40 {
			fmt.Println(string(buf))
			cycle = 1
		}
	}

	for _, instr := range lines {
		instr = strings.TrimSpace(instr) // sanitize CRLF

		doDrawPixel()

		switch {
		case instr == "noop":
		case strings.HasPrefix(instr, "addx"):
			cycle++
			maybeDrawRow()

			doDrawPixel()

			rest := strings.TrimPrefix(instr, "addx ")
			v, _ := strconv.Atoi(rest)
			regX += v
		}
		cycle++
		maybeDrawRow()
	}
}
