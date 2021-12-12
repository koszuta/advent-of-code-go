package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

const expectedResult = 150004

/*
 *   --- Day 12: Passage Pathing ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/12#part2
 */

var (
	connections map[string]map[string]struct{}
	paths       [][]string
)

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	nPaths := doPart2()
	log.Println("the number of paths that visit a single small cave twice is", nPaths)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	connections = make(map[string]map[string]struct{})
	paths = make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		cave1, cave2 := parts[0], parts[1]

		_, found := connections[cave1]
		if !found {
			connections[cave1] = make(map[string]struct{})
		}
		if cave2 != "start" {
			connections[cave1][cave2] = struct{}{}
		}

		_, found = connections[cave2]
		if !found {
			connections[cave2] = make(map[string]struct{})
		}
		if cave1 != "start" {
			connections[cave2][cave1] = struct{}{}
		}
	}

	findPath("start", []string{"start"}[:], 0)
	return len(paths)
}

func findPath(cave string, path []string, depth int) {
	if cave == "end" {
		paths = append(paths, path)
		return
	}
	for nextCave := range connections[cave] {
		isLargeCave := nextCave != strings.ToLower(nextCave)
		if isLargeCave || !contains(path, nextCave) || !smallCaveAlreadyVisitedTwice(path) {
			path = append(path, nextCave)
			findPath(nextCave, path, depth+1)
			path = path[:len(path)-1]
		}
	}
}

func contains(slice []string, S string) bool {
	for _, s := range slice {
		if S == s {
			return true
		}
	}
	return false
}

func smallCaveAlreadyVisitedTwice(path []string) bool {
	cavesVisited := make(map[string]struct{})
	for _, cave := range path {
		isSmallCave := cave == strings.ToLower(cave)
		if isSmallCave {
			_, visited := cavesVisited[cave]
			if visited {
				return true
			}
			cavesVisited[cave] = struct{}{}
		}
	}
	return false
}
