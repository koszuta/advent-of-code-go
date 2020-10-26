package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var distances [][]int
var paths []int

func findPathLength(from int, totalDist int, destinations []int) {
	if len(destinations) == 0 {
		paths = append(paths, totalDist)
	}

	for i := 0; i < len(destinations); i++ {
		to := destinations[i]
		destinations = append(destinations[:i], destinations[i+1:]...)
		if from >= 0 {
			totalDist += distances[from][to]
		}
		findPathLength(to, totalDist, destinations)
		if from >= 0 {
			totalDist -= distances[from][to]
		}
		destinations = append(destinations, 0)
		copy(destinations[i+1:], destinations[i:])
		destinations[i] = to
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	index := 0
	indices := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts1 := strings.Split(line, " to ")
		parts2 := strings.Split(parts1[1], " = ")

		from := parts1[0]
		to := parts2[0]

		_, ok := indices[from]
		if !ok {
			indices[from] = index
			index++
		}
		_, ok = indices[to]
		if !ok {
			indices[to] = index
			index++
		}
	}
	fmt.Printf("%v\n", indices)

	distances = make([][]int, index, index)
	for i := range distances {
		distances[i] = make([]int, index, index)
	}

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts1 := strings.Split(line, " to ")
		parts2 := strings.Split(parts1[1], " = ")

		from := parts1[0]
		to := parts2[0]
		dist, err := strconv.Atoi(parts2[1])
		if err != nil {
			panic(err)
		}
		
		// fmt.Printf("%s[%d] -> %s[%d] = %d\n", from, indices[from], to, indices[to], dist)
		distances[indices[from]][indices[to]] = dist
		distances[indices[to]][indices[from]] = dist
	}

	for _, dist := range distances {
		fmt.Printf("%v\n", dist)
	}

	destinations := make([]int, index, index)
	for i := 0; i < len(destinations); destinations[i], i = i, i+1 {}
	fmt.Printf("%v\n", destinations)
	findPathLength(-1, 0, destinations)
	// fmt.Printf("paths=%d\n", paths)
	fmt.Printf("len(paths)=%d\n", len(paths))

	shortestPath := -1
	longestPath := -1
	for _, path := range paths {
		if shortestPath < 0 || path < shortestPath {
			shortestPath = path
		}
		if longestPath < 0 || path > longestPath {
			longestPath = path
		}
	}
	fmt.Printf("shortestPath=%d\n", shortestPath)
	fmt.Printf("longestPath=%d\n", longestPath)
}
