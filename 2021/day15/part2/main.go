package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

const expectedResult = 2979

/*
 *   --- Day 15: Chiton ---
 *      --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/15#part2
 */

type coord2D struct {
	x, y int
}

const extension = 5

var size int

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	totalRisk := doPart2()
	log.Println("the lowest total risk of any path from the top left to the bottom right", totalRisk)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cave := make(map[coord2D]int)

	graph := dijkstra.NewGraph()

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		size = len(line)
		for j, r := range line {
			riskLevel, _ := strconv.Atoi(string(r))
			for q := 0; q < extension; q++ {
				for p := 0; p < extension; p++ {
					risk := riskLevel + p + q
					if risk >= 10 {
						risk -= 9
					}
					pos := coord2D{i + p*size, j + q*size}
					cave[pos] = risk
					graph.AddVertex(pos.index())
				}
			}
		}
	}

	source := coord2D{0, 0}
	target := coord2D{size*extension - 1, size*extension - 1}

	for j := 0; j < size*extension; j++ {
		for i := 0; i < size*extension; i++ {
			centerPos := coord2D{i, j}
			center, err := graph.GetVertex(centerPos.index())
			if center == nil || err != nil {
				log.Panicln(center, err)
			}

			addNeighborArc := func(neighborPos coord2D) {
				if neighborPos.x >= 0 && neighborPos.y >= 0 {
					neighbor, err := graph.GetVertex(neighborPos.index())
					if neighbor == nil || err != nil {
						log.Panicln(neighbor, err)
					}
					err = graph.AddArc(center.ID, neighbor.ID, int64(cave[neighborPos]))
					if err != nil {
						log.Panicln(err)
					}
					err = graph.AddArc(neighbor.ID, center.ID, int64(cave[centerPos]))
					if err != nil {
						log.Panicln(err)
					}
				}
			}
			addNeighborArc(coord2D{centerPos.x - 1, centerPos.y})
			addNeighborArc(coord2D{centerPos.x, centerPos.y - 1})
		}
	}

	bestPath, err := graph.Shortest(source.index(), target.index())
	if err != nil {
		log.Panicln(err)
	}
	return int(bestPath.Distance)
}

func (c *coord2D) index() int {
	return c.y*size*extension + c.x
}
