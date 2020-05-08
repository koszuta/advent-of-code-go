package main

import (
	"bufio"
	"fmt"
	"os"
)

const steps = 100
var lights, buffer [][]bool

func step() {
	rr := len(lights)-1
	cc := len(lights[0])-1
	for r := range lights {
		for c := range lights[r] {
			if (r == 0 || r == rr) && (c == 0 || c == cc) {
				buffer[r][c] = true
				continue
			}
			onNeighbors := 0
			if r > 0 && lights[r-1][c] {
				onNeighbors++
			}
			if r > 0 && c < cc && lights[r-1][c+1] {
				onNeighbors++
			}
			if c < cc && lights[r][c+1] {
				onNeighbors++
			}
			if r < rr && c < cc && lights[r+1][c+1] {
				onNeighbors++
			}
			if r < rr && lights[r+1][c] {
				onNeighbors++
			}
			if r < rr && c > 0 && lights[r+1][c-1] {
				onNeighbors++
			}
			if c > 0 && lights[r][c-1] {
				onNeighbors++
			}
			if r > 0 && c > 0 && lights[r-1][c-1] {
				onNeighbors++
			}
			if lights[r][c] {
				buffer[r][c] = onNeighbors == 2 || onNeighbors == 3
			} else {
				buffer[r][c] = onNeighbors == 3
			}
		}
	}
	lights, buffer = buffer, lights
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	lights = make([][]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, 0)
		for i := 0; i < len(line); i++ {
			row = append(row, line[i] == '#')
		}
		lights = append(lights, row)
	}
	lights[0][0] = true
	lights[0][len(lights[0])-1] = true
	lights[len(lights)-1][0] = true
	lights[len(lights)-1][len(lights[0])-1] = true

	buffer = make([][]bool, len(lights))
	for i := range buffer {
		buffer[i] = make([]bool, len(lights[0]))
	}

	// for _, row := range lights {
	// 	fmt.Println(row)
	// }
	// fmt.Println()

	for i := 0; i < steps; i++ {
		step();
	}
	
	onLights := 0
	for _, row := range lights {
		for _, col := range row {
			if col {
				onLights++
			}
		}
	}
	fmt.Printf("onLights=%d\n", onLights)
}
