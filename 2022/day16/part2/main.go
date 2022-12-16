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
	maxFlow int
	lines   []string

	inputRegex = regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z]{2}(, [A-Z]{2})*)`)
)

const (
	ELEPHANT = iota
	YOU
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

	pressures := make(map[string]int, len(lines))
	branches := make(map[string][]string, len(lines))

	for _, line := range lines {
		matches := inputRegex.FindAllStringSubmatch(line, -1)[0]

		valve := matches[1]
		pressure, _ := strconv.Atoi(matches[2])

		if pressure > 0 {
			pressures[valve] = pressure
		}
		branches[valve] = strings.Split(matches[3], ", ")
	}

	remainingPressures := make([]int, 0, len(pressures))
	for _, pressure := range pressures {
		remainingPressures = append(remainingPressures, pressure)
	}
	sort.Slice(remainingPressures, func(i, j int) bool {
		return remainingPressures[i] > remainingPressures[j]
	})

	doFindMaxFlow(0, 26, YOU, [2]string{"AA", "AA"}, remainingPressures, map[string]struct{}{}, [2]map[string]struct{}{{}, {}}, pressures, branches)

	log.Println("most pressure released:", maxFlow)
}

func doFindMaxFlow(totalPressure, minute, turn int, valves [2]string, remainingPressures []int, opened map[string]struct{}, visited [2]map[string]struct{}, pressures map[string]int, branches map[string][]string) {
	if minute == 0 || len(remainingPressures) == 0 || PotentialPressure(minute, remainingPressures)+totalPressure < maxFlow {
		if totalPressure > maxFlow {
			maxFlow = totalPressure
		}
		return
	}

	minute -= turn
	valve := valves[turn]

	// Maybe turn on the valve
	if pressure, pressurized := pressures[valve]; pressurized {
		if _, open := opened[valve]; !open {
			opened[valve] = struct{}{}
			before := visited[turn]
			visited[turn] = map[string]struct{}{valve: {}}
			doFindMaxFlow(totalPressure+pressure*minute, minute, DoTurn(turn), valves, Remove(pressure, remainingPressures), opened, visited, pressures, branches)
			visited[turn] = before
			delete(opened, valve)
		}
	}

	// Move to the next valve
	for _, nextValve := range branches[valve] {
		if _, found := visited[turn][nextValve]; !found {
			visited[turn][nextValve] = struct{}{}
			valves[turn] = nextValve
			doFindMaxFlow(totalPressure, minute, DoTurn(turn), valves, remainingPressures, opened, visited, pressures, branches)
			valves[turn] = valve
			delete(visited[turn], nextValve)
		}
	}
}

func PotentialPressure(t int, remainingPressures []int) (potentialPressure int) {
	for i, t := 0, t; i < len(remainingPressures) && t > 0; i, t = i+2, t-2 {
		potentialPressure += remainingPressures[i] * t
		if i < len(remainingPressures)-1 {
			potentialPressure += remainingPressures[i+1] * t
		}
	}
	return
}

func DoTurn(turn int) int {
	if turn == YOU {
		return ELEPHANT
	}
	return YOU
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
