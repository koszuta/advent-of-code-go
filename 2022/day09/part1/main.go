package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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

	head, tail := Vec2{}, Vec2{}

	visited := make(map[Vec2]struct{})
	visited[tail] = struct{}{}

	for _, line := range lines {
		line = strings.TrimSpace(line) // sanitize CRLF
		parts := strings.Split(line, " ")

		dir := dirDecoder[parts[0]]
		nSteps, _ := strconv.Atoi(parts[1])

		for i := 0; i < nSteps; i++ {
			head = head.Add(dir)

			if int(Distance(head, tail)) > 1 {
				switch {
				case dir.x == 0:
					tail.x = head.x
				case dir.y == 0:
					tail.y = head.y
				}
				tail = tail.Add(dir)
				visited[tail] = struct{}{}
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
