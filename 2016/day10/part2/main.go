package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 *   --- Day 10: Balance Bots ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/10#part2
 */

var (
	instructions map[int]dest
	bots         map[int]*bot
	outputs      map[int]int

	inputRex, handoffRex *regexp.Regexp
)

type bot struct {
	low, high int
}

type dest struct {
	low, high string
}

func init() {
	inputRex = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	handoffRex = regexp.MustCompile(`bot (\d+) gives low to (\w+ \d+) and high to (\w+ \d+)`)

	instructions = make(map[int]dest)
	bots = make(map[int]*bot)
	outputs = make(map[int]int)
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputMatch := inputRex.FindAllStringSubmatch(line, -1)
		handoffMatch := handoffRex.FindAllStringSubmatch(line, -1)
		if inputMatch != nil {
			chip, _ := strconv.Atoi(inputMatch[0][1])
			botNo, _ := strconv.Atoi(inputMatch[0][2])
			giveBotMicrochip(chip, botNo)
		}
		if handoffMatch != nil {
			botNo, _ := strconv.Atoi(handoffMatch[0][1])
			lowDest := handoffMatch[0][2]
			highDest := handoffMatch[0][3]

			_, found := instructions[botNo]
			if found {
				log.Panicln("instruction already found for bot", botNo)
			}
			instructions[botNo] = dest{lowDest, highDest}
		}
	}

	for botNo, chips := range bots {
		if chips.high != -1 {
			executeInstruction(botNo)
			break
		}
	}

	log.Printf("the product of outputs 0, 1, and 2 is %d\n", outputs[0]*outputs[1]*outputs[2])
}

func executeInstruction(botNo int) {
	bot := bots[botNo]
	dest := instructions[botNo]

	lowParts := strings.Split(dest.low, " ")
	lowNo, _ := strconv.Atoi(lowParts[1])
	if lowParts[0] == "bot" {
		giveBotMicrochip(bot.low, lowNo)
		lowBot := bots[lowNo]
		if lowBot.low != -1 && lowBot.high != -1 {
			executeInstruction(lowNo)
		}
	} else {
		outputs[lowNo] = bot.low
	}

	highParts := strings.Split(dest.high, " ")
	highNo, _ := strconv.Atoi(highParts[1])
	if highParts[0] == "bot" {
		giveBotMicrochip(bot.high, highNo)
		highBot := bots[highNo]
		if highBot.low != -1 && highBot.high != -1 {
			executeInstruction(highNo)
		}
	} else {
		outputs[highNo] = bot.high
	}
}

func giveBotMicrochip(chip, botNo int) {
	thisBot, found := bots[botNo]
	if !found {
		bots[botNo] = &bot{chip, -1}

	} else {
		if chip > thisBot.low {
			thisBot.high = chip
		} else {
			thisBot.high = thisBot.low
			thisBot.low = chip
		}
	}
}
