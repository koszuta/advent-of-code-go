package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var (
	rocks map[Vec2]struct{}
)

const (
	// Input = "><<><>><<<>><><<<>>>><<<>><<<><<<<>>>><<>>><><>>>><<<<>><<<<><<<<>>>><<>>><>>>><<<<>>>><<<<><<<>>><<<<>>><<<><<<>><<<>><<><>><<>><<<>><<<>>>><<<><<<><<<<><><<<<>>><<<>><<><<>><<<<><<><>>><<<><<<>>><<<><>>><<<<>><<<>><<<<>><<>><<<><>>><<<>>>><<<<>>>><<>><<>>>><>><<>>>><<<>>>><<<<><<<><<>>><<<<>>><<>>><<<<>><><<<>><<<>>><<><><<>><<<><<<<>>>><<<>>><<>><<<>>>><><>>><>>><<<>>><<<>>><<>><<<>><>>>><><<>>>><><>>>><<<>>><<<<><>>>><>><<<<>>>><<<>><<<><<><<<>><><>>>><<<<>><>>><<>><<>>><<<><<<>>>><<>>>><><<>>><>>>><<>>><><<><<<<>><<>>>><<>>>><>><<<<><>>>><>>><>>>><<<>><<<><<<<><<<>>><>>>><<<<>><>>>><<<<>>><><<<<><>>><<<<>><<><<<>>>><<<<>><<<><>>>><><<<<>>>><<<>>><<<><<<>>>><><<<<>>><<><<<<>>><>><<<>>>><>>><<><<>>>><>><<<>><>><<>><>>>><<<><>>><<<>>>><<<<>>>><<<<>>>><<<><<>><><>>><<>><>>>><<<<><<<<>>><<<<>>><>><>><>><<<>>>><><>>>><<><<<>>>><<<<>>>><<<>>>><<>><>>>><>><><<<>>>><<><<><<<<>>>><>>>><<<>><>><<<<><<>><>>>><<<<>>>><<<<>>><><<><>><<>>><<<>>><<<>><<>>><<<<>>><<>><<><<<<>>>><<<>>><>><>>><>>><<<<>><<<>>><>>>><<<><<<><<<>>><<<<>>>><<<>>><<<>>><<<>>><<<<>>>><<<>>><<<>>>><<>>><<<<><<<>><<<>>>><<<<>>>><>>><<<<>><<>><<<<><>><<>>>><<>>><<<>>>><<><>>><<>><>>>><<<<>>><<<<><<<<><>>>><>><<>><>>><<<>>>><<<<>>><>>>><><<<<>>><<<<><<<><<>><<<<>>><<<<>><<>>>><<<>>>><<<<>><<><<<><<<>>><<<<>>><>><<<<><<<><<>>>><<<><<<<>><<<><<<<>>>><<<<><<>>><<<<>>><<<>>>><><<<><>>>><<<><<<<>>><<>>>><<<<>>>><<>>><>><<<>><>>>><<<<>>>><<<<>>>><<>>>><<>>><<<><<><>>>><>>>><<>>><<<>>>><<<><<<<><>>><<<>>>><<<<>><<<>>><<>>>><<<<><<<>><<>>>><><<<>>><<<><<<>>>><<>>>><<<>>>><<<<>>>><>>>><<<>><>><<<>><>>>><<>>><>><><<<>>>><<<<>>>><>>><<<<><<<>><<>><<<<>>><<>><<>>><<<>>><><<>><<<<>><<>>>><<<<><<<>>><>><<<>>>><>><<<<>><<<>>>><<<<>><><<>><<<><<<>><<<>>>><<<<>><<<<>>><<<><<<<>>>><<<>><<<<><<<>>><<<><<>><<>>><<<<>><>>>><<><<<>>><<>>><<<>>>><<<><<>><><>>><>>><<<>><<<>>><<>>><<<<><><<<<>><<<<>>><<<>>><<<<>><<<><<>>>><>><<<>><<<<>><>>>><<<><<>>>><<<>>><>><<<<>>>><<<<>>>><<<>>>><<<><><<<<>>>><<<<><<<>>><<<>><<>>><<>>><>>>><<<>>><<<>>>><<>><<>><<>>><<<>>>><<><<<><<<<>><<<<><>><<<<>><>>>><<>><<<<>>><>>><>>><<<>>><<<<>>><<<>>><<<<>>><<<>>><><>><<<<>>>><>>>><><<<<><<>>><<<>>>><<<>>><<<<>>><<<>>><<>>><<<><<<<>><<>>><<<>>><<<<>>><>>>><<<>>><<<>>><>>>><<>>><<<<>>><<<>><<<><<<<><<>><<<><<<>><<<<>>>><><<<>>><<<<><><<<<>><<<<>><<<>><<<<>>>><<>><<<>>><><>><<<>>>><<<<><<<<>><<><<><<<<><<><<<><>><<>><>>><<<<>>><<<>>>><>>>><<><>>>><><<<<>>>><<<><>>><><<<<>><<>><<>>><<>>>><<>>><>>>><<<<>>><<>>>><<<><<<<>>><<><<>><<<>><>>>><<<>>>><<>>><<<><>><<<><<<<><<>><>><<>>><<<>><<<<>>><<<<><<<>>>><<<<>><<<>><<<<><>>><<<><<<<>>><>>>><<>><<<>>>><><<>>><<><<>>>><<>>><<<<>>>><<<<><<<<>><<>><>>><>>><<><<>>>><>>>><<<>><<<><<<<>>>><<>>>><><<<><<>><<>>>><<>>><<<<><<<<>>>><>>><<<<>><<<>><>><<<<>>>><>>>><<<>>>><<>>>><><<>>>><<<>>>><>><>>><<<><<<>>><>><<<<>><>><<<>>>><<<<>>><>>><>><<<><<<><<<>>><>>><<>>>><<<<>>><<<>><<<>><<>><<<<>>>><<>>>><<<><>>><<><<<<>>>><<<><<<<><<<>>>><<<><<><>>><<<>>>><>>>><<<<>>><><<<>>>><>><<><<<><<<>>><<>><<<<>>>><<<><<<<>>>><>>><<<>>><><<<<>><>>><<<><<<>>>><<><<<<>>><<<>>>><<<<>><<>>><><<><<>><>>>><<<<>>>><><><<<<>>>><<<>>>><>>><><<<<>><<<><<<<>>>><<>>><<<>>><<><>>><<<<>><<>>>><>>><<>>>><<<<>><<<>>>><<<><<>>>><<>>>><<>><<><>>>><<>>><<<<>>>><<<<>><<<<><<>>>><<<>>><<><<><<<><<>>>><><<><<>>>><<>>>><<<>>><<<>>><<>>><<>><><><<<>>><<>>>><<<>><<<>><<>>><<<>><>>>><<>>><<<>>><>>><<<<><<>>>><<>>><<>>>><<>><>>><<><<<><<<><<<<>><<<<>>>><<<<>>>><<<<>>>><<>><>>>><<<<>>>><<>>>><<<<>><<<<>><<><<<<>>><<<<>>><<<><>><>><<>><<><<>>><<>>>><<>><>>><>>><<>>>><<<><>>>><<>>><<>>><<<<>>><<<>>>><>>>><>>>><<>>>><<<>><><<<<>>>><<<<>>>><<><<<>>><<<<><<>><<>><><<<>>><>>>><<<>>>><<<<>><>>>><<>><<><<<<>><>><<<<>>><<><><><>><<>>><<<>>>><<<>>><<<>>>><<<>><<><<>><<<>>>><><>>>><<<<>>><<<<><>>>><>><<>>><<>>>><<<>><><<<<>><<<<>>><<<<>><>><>>><<<<><<<<>>>><<<<><<<>><<<>>>><>>>><<<><<<<><<<<><<>>><<>><<<<>><<<><<>>><<<>><>>><>>><><<><<<>><>><>>><<>>><<<<>><>>><<<><><<<<>><<>>>><<<>><<>><<<>>><<<><>>><<<>>><<<<><>>><<>><<<><<<<><<<<>><<>>><>><>><<><><<<>>><<>>><<>>>><<>>><<<><<<>><<<<>>><<<><<>>><<>>><>>><>>>><<><<<<>>>><>>>><>>><<<<>>><<<<>>>><<<<>><<<<>>>><>>>><<>>>><<>>>><<<>>><><<<>>>><><><<<>>><<>><>><<<<>><<>>><<<><<<>><<<><<<<>>>><<<>><<>><<>>>><>>>><<<>>>><<<<>><>>><>>>><<<<>>><<>>><<<>><<<<><<<>>><<<<>><>><<<>><<>>><<<<>>>><<<<>>>><>>>><<<<>><<>><<<>>><<<<><<>>><<<<>><<<<>>><>><><<<<>>><<<><<<>>>><<>><<>><>>><>>>><<<<><<<<><<>>><>>>><<>>>><><<>><>>>><><>>>><<><<<<><<>>>><<<<><<<<>>><<>>>><<>>>><<<<>>><<>><>>><>><<><>>>><<<>><<<>>><<<>>><>>>><<<><>>><<>>>><<>>>><<<<>>>><<<>>>><<<>><>>>><>>>><<>>>><<>>>><>>>><<>>>><<<<><>>><<<><><>>><<<<><><>><<>>><<<<><<<>>><<<><<><<><<<<><<><<>>><<<><><<>>>><<<<>>>><<<<>>>><<<>><<>><>><<>>>><<<><<<<><<<<>>>><>><<<<>>>><>>><<<<>><><<><<<<><<>>><><>>>><<<>>><<<>>><><<<<>>>><<<<>>>><<<<><>>>><<<>><<>><>><<<<>><<<<>>><<>>>><<>><<<>>><<>><>>>><<<>>>><<>>>><<>>>><><><<<<><<<>>>><<>><<<<>><<>>>><<<>>>><>>>><<<>><<>>><>>>><<<>>>><<><<>>>><<>><<<>>>><>>><<<<>>><<>>>><<><<>>>><<>>><>>>><<<><<><>>><<<<><<<<>><<<<>>><<<<><<><<>><<>>><<<<>>>><<>>><<<<>>><<<>><<<>>>><<><><<<<>><<<<>>><<>>><>>><<>>><>>><<<>>>><>>>><<>><<<<>>>><<>><<><<>>><<><<<<><><>>><<><<>>>><<<><<<<>>><<<<>>>><>><>>><>><<<<>><>>><<<<>>><><<><<<<>>>><>><<<>>>><<>><>>>><<>>>><<<>>>><<>>>><><<<>><<>>>><<<>>>><>>><<<<>>>><>>><<>><<><><<><>>><<>><<<>>>><<>><><<><<<>><<<>><<<<><<>><<>>><<<<><<<<><<<>>>><<<<><<<><<>>>><<<<><<<><<<<><>><<>>><<<>>>><>>><<<<>>><<>>><<<<>><<<<>><<<>><<><><<<<><<>>>><<<<>>>><<<<>><<<>><<<<>><<<>>>><<>>>><<>>>><<<>><<><<<>>>><>><<<>><<>>>><>><<<>>>><<<>><<<>>>><>><>><<<>><<<><<<><<<<>><>>>><<<<>>>><<<<><<<>>>><<<<>><>>>><<>>><<<>>><<><<<>>>><><<<>><<>>>><<<<>><>><<<<>>>><<>><<<<><<<>>>><>>>><<<><><><>>><<<>>><>>>><<>>>><>><>>><<<><<>>>><<<<>><>>>><<<<><<>>>><>>>><<<>>>><<<<>>><<<<>>>><<<<><<<<>><<<<>>><<<<>><><<><<><>>><<<<>>><<>><<>><<>><<<><<><<>>>><<>><>>><<><<>>>><<<>>><>>>><<>><<<>><>>><><<<>>><<<<><>><<<>>><<<<>>>><<><><<>>><<>><<>>>><<>><>>>><><<<<><>>><<<>>><<<>><<>>>><>>>><>>>><<<><>>>><<>>><<><<><<><<<<>>>><<<>>>><<<><>>>><<<>>>><<<>>><<<><<<<>>><>>>><<<<>>>><<>>>><>>><<>><><>><<<>><<<><<<<>>>><<<<>><<<>>>><>>>><<<><<<>><<>>><>><<>>>><<>>><<<<>>>><<>><<<<>>>><<<<><<>>><<<><<<<>>><<<<><<<<>>>><<<>><<<<><<<>>><<<>>><<<<><<<<><><<<<><>>><<<<>>><<<<>><<<>><<<<>>><<<<>>>><<>>>><>>><<<>><<>>>><>>><<<>><<<><<<>>>><>>><>><<><<<<>><<<<>>><<>>>><<<><<>><<<>>><<<>><<<<><<<<>><>><<<<><>><>>>><><<<<>>><><<<<>>>><><<<<>><>><<<>><<<<><<>><<<><<<><>><<<><>>>><<<><<<>><<<>>>><<>><<<><>><>><>><<>><<<<>><>>><>><<<>>>><>><<<<>>>><>><<<<>>>><<>>>><<>><<>>>><<><<<>><<><<<<>>><<<<><<>>><<>>>><<<<><<>><<<><<<>>><<<>>><<<>><<<<><<<>><>><<<<>>>><<<<>><<<>><><<<>>>><>><<<>>>><<>><<><<<<><<<<>><<<<><>><<<>>>><<>>><>><<<>>><<<<>>><>>><<>>>><>><<<<>><<<<>><><<<>><<>><>><<>>><<<>><<<<>>>><<<>>><<<>>><><<<><<>>><<<<>><<><<<<><<>>>><<<>>><<>><>><>>>><<<<>>><<><>>><>>><<<<>><<<<>>>><<<<>>>><<>>>><>>>><<<>><<<><<<>>>><><<><<<<><<<<>><>>><<<>>><><<>><<<<>><<<<>>>><>>>><<<<>>>><<<<><><>>><<<>><><<<>>>><>><<<>>>><<><<<>>><<<<><>>><<<<>>>><<<<>>><<<><<<>><<<><<>>>><<><<<<>>><<<<>>><><>>>><>>><<<>>>><<<<>>>><<>>><>>><<<><<<><>>><<>><<<<>>><>>>><<<>><<<><<<><<>>><<<<>>><<<>>>><>>><<<>>><<<<>>>><<>><<<<>>><<<>>><<<<>><<<<><<<<>><<<<><>>><<<><>>>><<<>>><<>>>><>>>><><<>>>><<>><<<<>>><<<<>>><<<>><<<>>><<<<>><<<<><<<>>>><<<><>>><<<><<<<><<>><<<>>><>>><<<<>>><<>><<>>><<<>>><<<><<<>><<<<>>>><<<>><<<<>><<<><<>>>><<<<><<<<>>><<<<><<>>>><<<><<<<><<<<>>><>>><<<<><<>><<<>><<<>><<>><<<><>><<>><>>>><><<<<><<<>>>><<<><<<>>>><>><<<><<>><>><<<<>>>><<<>>><>>><<<<><<>>><<>>><<<>>><<<>>><<>><<<<>><>>>><<<><<<<>><>>><<<<>><<<>><<<>>>><<<>>>><><<<>>><>>>><>><<>><<><><<<><>>><<>>>><>>><<<<>>><<<<>><<<><<<>><<>><><<<<>><<<<><>>><<<<>>>><>>>><>>>><<><<<>>>><<<<>><<<<><<<<>>><<<<>>>><<<><<>>><<><<<<>><<>>>><>><<<<><<><<<<>>>><<>>><<<<><<<>>><<<<><<><<<<>><>>><>><<<<>>><<>><<<><<><<>>><<<>>>><<<<>>>><><<<><<<<>>>><<<<>>><<<>>>><>><<<<>>>><<<<><>>><><<<>><>>>><<<<>><>><><<><<<><<<<>>>><>>>><>><<><>>><<<>>>><<<>><>>>><<<<>><<>>>><><<<>><<<>><<>>>><<>>><>><><<<<>><>><>>><>><<<>>>><<<<>>><<>>><<<<>>><<<>>><<>>>><<><<<<>>>><<<<><>><<>>>><>>><>><><<<<>><<>>><<<><<<<>>>><<><<>><<<>>><<><><<<>>><>>>><<<>><<<<>>>><<>>><<>>><<>><><<<<>><<>><<>>><<><<>><<>><<>>>><<<<>>><<><<>>><<<<><<<<>>><<<<>>><<<<>>>><<<><>><>><<<>>>><<<<>>>><<<>><<>>>><<<<>>><<<>>><<>>>><<><<<<><<<>><<<<>>><<<><><<<>>>><>>>><<>>>><>><>>><<<<>>><<<<>>>><>>><>>>><>><<<>>><<>>><><<<<>>>><><<<<><<<<>>>><<<>>>><<>>><<<<>>><>>><<>>><<>><<<>>><<<<>><<<<><<<>><<<>>><>>>><<<<><>>><>>><<<<>>>><<<<>>>><<<><<>><<<<>>><<<>>><<>>><<<<><>>><>>>><<<<>>><<>><<<><<<>>>><<<<><>>>><<<><>>>><<<<>><><<><<>>><>>>><>><<>>>><<<<>><<>>><<<><<<>>><<<<>><>>><>>><>>><<<>><<<>>><<>>>><<<>><<<<>>>><><<<><>>>><<>>>><<<>>>><<<<>>><<><>>>><<>>>><<>>><<<>>>><>>>><<><<<<>>><><<>>><<<>><<<<>>>><<>>>><<><<>>>><<<<>>><>><<>><<>>>><<<>>><<>>>><<>>><<>>><><<><>>><<>>><>>>><><>><<><<>><<><>><<<>><<>>><<<>><<>>>><><<><<<>>>><<<<>>><<><<>>>><<<>>>><>>>><<>>><<>><><><>>>><<<<>>><<<>>>><<<><<>>><>>>><<>>><<>><>>><>>><<<>>>><<>>>><<<>>><<<<><<>>><<>>><<<>>>><<>>>><<<><<<<>><<<>>>><<<>><<<>>><<<>>>><<>><<<><<<<>>>><<>><>><<<>><<>>>><<<>>>><>>>><>>><><<>>>><<<>><<<<>>>><<<<>><<>>><<<>><>>><<<<>><<>>>><><<<<>>>><><><<<<>><>>><<>>>><<<><<>><<<<>>><<<<>>><>>><<<<>><<<><>><<><<>>>><>>>><<<<><<>>>><<>>>><><<><<>>><<<>><>><<<>>><<<>>>><<<<>>><<<<>><>>><>><<<<><<>>><>><<>>>><<><<<<>>><<<<>>><<<>>><>>>><<>><<<<>>>><>><<<>>><><<>>><<<<>><>><><>>><<>>>><<<><<<>>>><<<>>>><><<<<>>><<<<><<>><<<<>>>><<<>>>><<<>>><<<>><<<>><<<>>><<<>>>><<<<>>>><>>>><<<>>><>>>><<>>><>>><>>><>><<<<><>>><>>><<>>><<<<>><><<>>><<>><<<<>><<<<><<>><<<<>><<>>>><<<<>><>>><<<<><<<><<<>>><>>>><<><<<>>>><<<>>><<>>><>>><<>><<>>><<<<>>>><<>>><<<>><<<>>><>>>><<<<>>>><<<<>>><<>>><<>>><<<>>><<<<><<<<>><<<<><<<<><><<<>>><<<>>><>>><><<>><<<<>>>><<<<>>><<>>><<<><>>><<<>>>><<>>><<>>><<<>>>><<>><<>>>><>>>><<<>>><<<><>>>><<<>><>><>>>><><<<>>>><<<<>>>><>>>><<><<<<><>>><>>><<<<>>><<>>>><<<<>><<>><<><<<<>>><>>><<<<><>>>><<<<>>>><<>><>>><<<>><><<>>>><<<><<<<>><<><<>>><>>><<<<>>>><<>><<<<>><<<<>><<"
	Input = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	Width = 7
)

type Vec2 struct {
	x, y int
}

type Dir int

const (
	Down = Dir(iota)
	Left
	Right
)

type Shape int

const (
	HBar = Shape(iota)
	Plus
	RevL
	VBar
	Square
	NShapes
)

type Rock struct {
	shape  Shape
	origin Vec2
}

func main() {
	defer func(t time.Time) {
		log.Println("took:", time.Since(t))
	}(time.Now())

	rocks = make(map[Vec2]struct{})

	highestRock := -1
	for move, shape := 0, HBar; shape < 2022; shape++ {
		rock := Rock{
			shape:  shape % NShapes,
			origin: Vec2{x: 2, y: highestRock + 4},
		}
		// PrintRocks(rock)

		// Move rock until it stops
		for do := true; do || rock.Move(Down); do = false {
			if JetPattern(move) {
				rock.Move(Right)
			} else {
				rock.Move(Left)
			}
			move++
		}

		// Update final positions
		for _, r := range ShapeRocks[rock.shape] {
			pos := rock.origin.Add(r)
			rocks[pos] = struct{}{}
			if pos.y > highestRock {
				highestRock = pos.y
			}
		}
	}
	log.Println("highest rock:", highestRock+1)
}

func (r *Rock) Move(dir Dir) bool {
	newPos := r.origin.Add(Directions[dir])

	for _, edge := range EdgeRocks[r.shape][dir] {
		pos := newPos.Add(edge)
		if pos.y < 0 || pos.x < 0 || pos.x >= Width {
			return false
		}
		if _, found := rocks[pos]; found {
			return false
		}
	}

	r.origin = newPos
	return true
}

func JetPattern(move int) bool {
	return Input[move%len(Input)] == '>'
}

func (u *Vec2) Add(v Vec2) Vec2 {
	return Vec2{x: v.x + u.x, y: v.y + u.y}
}

var (
	Directions = map[Dir]Vec2{
		Down:  {0, -1},
		Left:  {-1, 0},
		Right: {1, 0},
	}

	EdgeRocks = map[Shape]map[Dir][]Vec2{
		HBar: {
			Down:  []Vec2{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			Left:  []Vec2{{0, 0}},
			Right: []Vec2{{3, 0}},
		},
		Plus: {
			Down:  []Vec2{{0, 1}, {1, 0}, {2, 1}},
			Left:  []Vec2{{1, 0}, {0, 1}, {1, 2}},
			Right: []Vec2{{1, 0}, {2, 1}, {1, 2}},
		},
		RevL: {
			Down:  []Vec2{{0, 0}, {1, 0}, {2, 0}},
			Left:  []Vec2{{0, 0}, {2, 1}, {2, 2}},
			Right: []Vec2{{2, 0}, {2, 1}, {2, 2}},
		},
		VBar: {
			Down:  []Vec2{{0, 0}},
			Left:  []Vec2{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
			Right: []Vec2{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		Square: {
			Down:  []Vec2{{0, 0}, {1, 0}},
			Left:  []Vec2{{0, 0}, {0, 1}},
			Right: []Vec2{{1, 0}, {1, 1}},
		},
	}

	ShapeRocks = map[Shape][]Vec2{
		HBar:   {{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		Plus:   {{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
		RevL:   {{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		VBar:   {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		Square: {{0, 0}, {1, 0}, {0, 1}, {1, 1}},
	}
)

func PrintRocks(moving Rock) {
	var maxY int
	for _, r := range ShapeRocks[moving.shape] {
		pos := moving.origin.Add(r)
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	for r := range rocks {
		if r.y > maxY {
			maxY = r.y
		}
	}

	grid := make([][]byte, maxY+1)
	for y := range grid {
		grid[y] = make([]byte, Width)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	for _, r := range ShapeRocks[moving.shape] {
		pos := moving.origin.Add(r)
		grid[pos.y][pos.x] = '@'
	}

	for r := range rocks {
		grid[r.y][r.x] = '#'
	}

	for y := len(grid) - 1; y >= 0; y-- {
		fmt.Printf("|%s|\n", string(grid[y]))
	}
	fmt.Printf("+%s+\n\n", strings.Repeat("-", Width))
}