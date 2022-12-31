package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var lines []string

// Blueprint 1:
//   Each ore robot costs 4 ore.
//   Each clay robot costs 2 ore.
//   Each obsidian robot costs 3 ore and 14 clay.
//   Each geode robot costs 2 ore and 7 obsidian.

// Blueprint 2:
//   Each ore robot costs 2 ore.
//   Each clay robot costs 3 ore.
//   Each obsidian robot costs 3 ore and 8 clay.
//   Each geode robot costs 3 ore and 12 obsidian.

var inputRegex = regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

type Blueprint struct {
	oreBotOreCost        int
	clayBotOreCost       int
	obsidianBotOreCost   int
	obsidianBotClayCost  int
	geodeBotOreCost      int
	geodeBotObsidianCost int
}

func init() {
	b, _ := os.ReadFile("../input.txt")
	lines = strings.Split(string(b), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line) // sanitize CRLF
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-2]
	}
}

func main() {
	defer func(t time.Time) {
		log.Println("took:", time.Since(t))
	}(time.Now())

	var sum int
	for _, line := range lines {
		log.Println(line)
		matches := inputRegex.FindAllStringSubmatch(line, -1)[0]

		blueprintNo, _ := strconv.Atoi(matches[1])
		oreBotOreCost, _ := strconv.Atoi(matches[2])
		clayBotOreCost, _ := strconv.Atoi(matches[3])
		obsidianBotOreCost, _ := strconv.Atoi(matches[4])
		obsidianBotClayCost, _ := strconv.Atoi(matches[5])
		geodeBotOreCost, _ := strconv.Atoi(matches[6])
		geodeBotObsidianCost, _ := strconv.Atoi(matches[7])

		bp := Blueprint{
			oreBotOreCost:        oreBotOreCost,
			clayBotOreCost:       clayBotOreCost,
			obsidianBotOreCost:   obsidianBotOreCost,
			obsidianBotClayCost:  obsidianBotClayCost,
			geodeBotOreCost:      geodeBotOreCost,
			geodeBotObsidianCost: geodeBotObsidianCost,
		}
		sum += doCrackGeodes(0, 1, 0, 0, 1, 0, 0, 0, bp) * blueprintNo
	}
	log.Println("sum:", sum)
}

func doCrackGeodes(t, nOre, nClay, nObsidian, nGeodes, nOreBots, nClayBots, nObsidianBots, nGeodeBots, maxGeodes int, bp Blueprint) int {
	if t == 24 || (maxGeodes-nGeodes)/nGeodeBots > 24-t {
		if nGeodes > maxGeodes {
			return nGeodes
		}
		return maxGeodes
	}
	if nOre > 0 {
		newNGeodes := doCrackGeodes(t+1, nOre-1, nClay, nObsidian, nGeodes, nOreBots+1, nClayBots, nObsidianBots, nGeodeBots, maxGeodes, bp)
		if newNGeodes > maxGeodes {
			maxGeodes = newNGeodes
		}
	}
	return maxGeodes
}
