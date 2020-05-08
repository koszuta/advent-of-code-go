package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const liters = 150
var combinations = 0
var numContainers []int

func countCombinations(remaining int, count int, containers []int) {
	if remaining == 0 {
		combinations++
		numContainers = append(numContainers, count)
		return
	}
	for len(containers) > 0 {
		size := containers[0]
		containers = append(containers[:0], containers[1:]...)
		if remaining >= size {
			newContainers := make([]int, len(containers))
			copy(newContainers, containers)
			countCombinations(remaining-size, count+1, newContainers)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	containers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		container, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		containers = append(containers, container)
	}
	fmt.Printf("containers=%v\n", containers)

	numContainers = make([]int, 0)
	countCombinations(liters, 0, containers)
	fmt.Printf("combinations=%d\n", combinations)

	leastContainers := -1
	for _, c := range numContainers {
		if leastContainers < 0 || c < leastContainers {
			leastContainers = c
		}
	}
	numCombos := 0
	for _, c := range numContainers {
		if c == leastContainers {
			numCombos++
		}
	}
	fmt.Printf("numCombos=%d\n", numCombos)
}
