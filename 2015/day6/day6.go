package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type op int

const (
	turnOn op = iota
	turnOff
	toggle
)

func indexOf(x, y int) int {
	return 1000*x + y
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lights [1000000]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var op op
		if strings.Contains(line, "turn on") {
			line = line[8:]
			op = turnOn
			
		} else if strings.Contains(line, "turn off") {
			line = line[9:]
			op = turnOff
			
		} else if strings.Contains(line, "toggle") {
			line = line[7:]
			op = toggle
		}

		parts := strings.Split(line, " ")
		from := strings.Split(parts[0], ",")
		to := strings.Split(parts[2], ",")

		var err error
		var fromX, fromY, toX, toY int
		fromX, err = strconv.Atoi(from[0])
		if err != nil {
			panic(err)
		}
		fromY, err = strconv.Atoi(from[1])
		if err != nil {
			panic(err)
		}
		toX, err = strconv.Atoi(to[0])
		if err != nil {
			panic(err)
		}
		toY, err = strconv.Atoi(to[1])
		if err != nil {
			panic(err)
		}

		for x := fromX; x <= toX; x++ {
			for y := fromY; y <= toY; y++ {
				switch op {
				case turnOn:
					lights[indexOf(x, y)]++
				case turnOff:
					if lights[indexOf(x, y)] > 0 {
						lights[indexOf(x, y)]--
					}
				case toggle:
					lights[indexOf(x, y)] += 2
				}
			}
		}
	}

	overallBrightness := 0
	for _, brightness := range lights {
		overallBrightness += brightness
	}
	fmt.Printf("overallBrightness=%d\n", overallBrightness)
}
