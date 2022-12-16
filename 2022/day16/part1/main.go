package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	maxFlow     int
	pressures   map[string]int
	connections map[string][]string
	lines       []string

	inputRegex = regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z]{2}(, [A-Z]{2})*)`)
)

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

	pressures = make(map[string]int, len(lines))
	connections = make(map[string][]string, len(lines))

	for _, line := range lines {
		matches := inputRegex.FindAllStringSubmatch(line, -1)[0]

		valve := matches[1]
		pressure, _ := strconv.Atoi(matches[2])

		if pressure > 0 {
			pressures[valve] = pressure
		}
		connections[valve] = strings.Split(matches[3], ", ")
	}

	remainingPressures := make([]int, 0, len(pressures))
	for _, pressure := range pressures {
		remainingPressures = append(remainingPressures, pressure)
	}
	sort.Slice(remainingPressures, func(i, j int) bool {
		return remainingPressures[i] > remainingPressures[j]
	})

	doFindMaxFlow(0, 30, "AA", remainingPressures, map[string]struct{}{}, map[string]struct{}{})

	log.Println("most pressure released:", maxFlow)
}

func doFindMaxFlow(totalPressure, minute int, valve string, remainingPressures []int, opened, visited map[string]struct{}) {
	if minute == 0 || len(remainingPressures) == 0 || PotentialPressure(minute, remainingPressures)+totalPressure < maxFlow {
		if totalPressure > maxFlow {
			maxFlow = totalPressure
		}
		return
	}

	minute--

	// Maybe turn on the valve
	if pressure, pressurized := pressures[valve]; pressurized {
		if _, open := opened[valve]; !open {
			opened[valve] = struct{}{}
			doFindMaxFlow(totalPressure+pressure*minute, minute, valve, Remove(pressure, remainingPressures), opened, map[string]struct{}{valve: {}})
			delete(opened, valve)
		}
	}

	// Move to another valve
	for _, nextValve := range connections[valve] {
		if _, found := visited[nextValve]; !found {
			visited[nextValve] = struct{}{}
			doFindMaxFlow(totalPressure, minute, nextValve, remainingPressures, opened, visited)
			delete(visited, nextValve)
		}
	}
}

func PotentialPressure(t int, remainingPressures []int) (potentialPressure int) {
	for i, t := 0, t; i < len(remainingPressures) && t > 0; i, t = i+1, t-2 {
		potentialPressure += remainingPressures[i] * t
	}
	return
}

func Remove(target int, s []int) []int {
	newS := make([]int, len(s))
	copy(newS, s)
	for i, v := range newS {
		if v == target {
			return append(newS[:i], newS[i+1:]...)
		}
	}
	return s
}
