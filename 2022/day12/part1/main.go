package main

import (
	"log"
	"os"
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

	// Find the source and destination and update the heights of those squares
	var srcID, dstID int
	for y, line := range lines {
		for x, r := range line {
			switch int(r) {
			case 'S':
				srcID = ID(x, y)
			case 'E':
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

	best, err := movementGraph.Shortest(srcID, dstID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("shortest path:", best.Distance)
}

func ID(x, y int) int {
	return len(lines[0])*y + x
}
