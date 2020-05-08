package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var happinessMatrix [][]int
var totalHappiness []int

func calculateHappiness(start int, person int, happiness int, neighbors []int) {
	if len(neighbors) == 0 {
		happiness += happinessMatrix[person][start]
		happiness += happinessMatrix[start][person]
		totalHappiness = append(totalHappiness, happiness)
	}

	for i := 0; i < len(neighbors); i++ {
		neighbor := neighbors[i]
		neighbors = append(neighbors[:i], neighbors[i+1:]...)
		if person != -1 {
			happiness += happinessMatrix[person][neighbor]
			happiness += happinessMatrix[neighbor][person]
		} else {
			start = i
		}
		calculateHappiness(start, neighbor, happiness, neighbors)
		if person != -1 {
			happiness -= happinessMatrix[person][neighbor]
			happiness -= happinessMatrix[neighbor][person]
		}
		neighbors = append(neighbors, 0)
		copy(neighbors[i+1:], neighbors[i:])
		neighbors[i] = neighbor
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	numLines := 0
	for scanner.Scan() {
		numLines++
	}
	fmt.Printf("numLines=%d\n", numLines)
	numPeople := 0
	for ; numLines != numPeople*numPeople-numPeople; numPeople++ {}
	numPeople++
	fmt.Printf("numPeople=%d\n", numPeople)

	indices := make(map[string]int, numPeople)
	happinessMatrix = make([][]int, numPeople, numPeople)
	for i := range happinessMatrix {
		happinessMatrix[i] = make([]int, numPeople, numPeople)
	}

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		person := parts[0]
		neighbor := parts[10]
		neighbor = neighbor[:len(neighbor)-1]
		happiness, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err)
		}
		if parts[2] == "lose" {
			happiness = -happiness
		}

		_, ok := indices[person]
		if !ok {
			indices[person] = index
			index++
		}
		_, ok = indices[neighbor]
		if !ok {
			indices[neighbor] = index
			index++
		}
		
		happinessMatrix[indices[person]][indices[neighbor]] = happiness
	}

	for _, happinessRow := range happinessMatrix {
		fmt.Printf("%v\n", happinessRow)
	}

	neighbors := make([]int, numPeople, numPeople)
	for i := 0; i < len(neighbors); neighbors[i], i = i, i+1 {}
	calculateHappiness(-1, -1, 0, neighbors)

	highestHappiness := -1
	for _, happiness := range totalHappiness {
		if happiness > highestHappiness {
			highestHappiness = happiness
		}
	}
	fmt.Printf("highestHappiness=%d\n", highestHappiness)
}
