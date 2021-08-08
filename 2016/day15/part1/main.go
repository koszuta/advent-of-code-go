package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

/*
 *   --- Day 15: Timing is Everything ---
 *             --- Part One ---
 *
 *   https://adventofcode.com/2016/day/15
 */

type disc struct {
	pos, nPos int
}

var (
	discs []disc
)

func init() {
	nPosRex := regexp.MustCompile(`Disc #\d+ has (\d+) positions; at time=0, it is at position \d+\.`)
	startPosRex := regexp.MustCompile(`Disc #\d+ has \d+ positions; at time=0, it is at position (\d+)\.`)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		nPos, _ := strconv.Atoi(nPosRex.FindAllStringSubmatch(line, -1)[0][1])
		startPos, _ := strconv.Atoi(startPosRex.FindAllStringSubmatch(line, -1)[0][1])

		discs = append(discs, disc{startPos, nPos})
	}
}

func main() {
	time := 0
TICK:
	for {
		t := time
		for _, d := range discs {
			t++
			if (t+d.pos)%d.nPos != 0 {
				time++
				continue TICK
			}
		}
		break
	}
	log.Printf("the first time you can press the button to get a capsule is %d\n", time)
}
