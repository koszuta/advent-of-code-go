package main

import (
	"log"
	"os"
	"regexp"
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

	doFindMaxFlow(0, 30, "AA", map[string]struct{}{}, map[string]struct{}{})

	log.Println("most pressure released:", maxFlow)
}

func doFindMaxFlow(flow, minute int, valve string, opened, visited map[string]struct{}) {
	if minute == 0 || len(opened) == len(pressures) {
		if flow > maxFlow {
			maxFlow = flow
		}
		return
	}

	minute--

	// Maybe turn on the valve
	if pressure, pressurized := pressures[valve]; pressurized {
		if _, open := opened[valve]; !open {
			opened[valve] = struct{}{}
			doFindMaxFlow(flow+pressure*minute, minute, valve, opened, map[string]struct{}{valve: {}})
			delete(opened, valve)
		}
	}

	// Move to another valve
	for _, nextValve := range connections[valve] {
		if _, found := visited[nextValve]; !found {
			visited[nextValve] = struct{}{}
			doFindMaxFlow(flow, minute, nextValve, opened, visited)
			delete(visited, nextValve)
		}
	}
}
