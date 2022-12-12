package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

var lines []string

func init() {
	b, _ := os.ReadFile("../input.txt")
	lines = strings.Split(string(b), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line) // sanitize CRLF
	}
}

func main() {
	// Initialize the graph and vertices
	movementGraph := dijkstra.NewGraph()
	for i := 0; i < len(lines[0])*len(lines); i++ {
		movementGraph.AddVertex(i)
	}

	// Find the destination and update the height of it and the source square
	var dstID int
	for y, line := range lines {
		for x, height := range line {
			if height == 'E' {
				dstID = ID(x, y)
			}
		}
		lines[y] = strings.ReplaceAll(lines[y], "S", "a")
		lines[y] = strings.ReplaceAll(lines[y], "E", "z")
	}

	// Connect the graph
	for y, line := range lines {
		for x, r := range line {
			srcHeight := byte(r)
			addArc := func(dx, dy int, condish bool) {
				if condish {
					dstHeight := lines[y+dy][x+dx]
					if dstHeight <= srcHeight || dstHeight-srcHeight == 1 {
						movementGraph.AddArc(ID(x, y), ID(x+dx, y+dy), 1)
					}
				}
			}
			addArc(-1, 0, x > 0)               // left
			addArc(+1, 0, x < len(lines[0])-1) // right
			addArc(0, -1, y > 0)               // up
			addArc(0, +1, y < len(lines)-1)    // down
		}
	}

	var distances []int
	{
		distances = make([]int, 0)
		for y, line := range lines {
			for x, height := range line {
				if height == 'a' {
					best, err := movementGraph.Shortest(ID(x, y), dstID)
					if err == nil {
						distances = append(distances, int(best.Distance))
					}
				}
			}
		}
		sort.Slice(distances, func(i, j int) bool {
			return distances[i] < distances[j]
		})
	}
	log.Println("shortest path:", distances[0])
}

func ID(x, y int) int {
	return len(lines[0])*y + x
}
