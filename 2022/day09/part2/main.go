package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const nKnots = 10

var (
	dirDecoder = map[string]Vec2{
		"U": {x: 1, y: 0},
		"D": {x: -1, y: 0},
		"L": {x: 0, y: -1},
		"R": {x: 0, y: 1},
	}
)

type Vec2 struct {
	x, y int
}

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	var knots [nKnots]Vec2

	visited := make(map[Vec2]struct{})
	visited[knots[nKnots-1]] = struct{}{}

	for _, line := range lines {
		line = strings.TrimSpace(line) // sanitize CRLF
		parts := strings.Split(line, " ")

		dir := dirDecoder[parts[0]]
		nSteps, _ := strconv.Atoi(parts[1])

		for i := 0; i < nSteps; i++ {
			knots[0] = knots[0].Add(dir)

			for k := 1; k < len(knots); k++ {
				if int(Distance(knots[k-1], knots[k])) > 1 {
					var move Vec2
					if knots[k-1].x-knots[k].x < 0 {
						move.x = -1
					}
					if knots[k-1].x-knots[k].x > 0 {
						move.x = 1
					}
					if knots[k-1].y-knots[k].y < 0 {
						move.y = -1
					}
					if knots[k-1].y-knots[k].y > 0 {
						move.y = 1
					}
					knots[k] = knots[k].Add(move)
					visited[knots[nKnots-1]] = struct{}{}
				} else {
					break
				}
			}
		}
	}
	log.Println("n positions visited by tail:", len(visited))
}

func Distance(v, w Vec2) float64 {
	d1, d2 := float64(w.x-v.x), float64(w.y-v.y)
	return math.Sqrt(d1*d1 + d2*d2)
}

func (u *Vec2) Add(v Vec2) (w Vec2) {
	w.x = u.x + v.x
	w.y = u.y + v.y
	return
}
