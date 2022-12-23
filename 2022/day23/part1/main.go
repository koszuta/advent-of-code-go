package main

import (
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

type Vec2 struct {
	x, y int
}

var (
	conditions = [4]func(elf Vec2, n, ne, nw, s, sw, se, w, e bool) (Vec2, bool){
		func(elf Vec2, n, ne, nw, s, sw, se, w, e bool) (Vec2, bool) {
			if !n && !ne && !nw {
				return Vec2{x: elf.x, y: elf.y - 1}, true
			}
			return Vec2{}, false
		},
		func(elf Vec2, n, ne, nw, s, sw, se, w, e bool) (Vec2, bool) {
			if !s && !se && !sw {
				return Vec2{x: elf.x, y: elf.y + 1}, true
			}
			return Vec2{}, false
		},
		func(elf Vec2, n, ne, nw, s, sw, se, w, e bool) (Vec2, bool) {
			if !w && !nw && !sw {
				return Vec2{x: elf.x - 1, y: elf.y}, true
			}
			return Vec2{}, false
		},
		func(elf Vec2, n, ne, nw, s, sw, se, w, e bool) (Vec2, bool) {
			if !e && !ne && !se {
				return Vec2{x: elf.x + 1, y: elf.y}, true
			}
			return Vec2{}, false
		},
	}
)

var lines []string

func init() {
	b, _ := os.ReadFile("../input.txt")
	lines = strings.Split(string(b), "\n")
	for i, line := range lines {
		lines[i] = strings.Trim(line, "\r") // sanitize CRLF
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-2]
	}
}

func main() {
	defer func(t time.Time) {
		log.Println("took:", time.Since(t))
	}(time.Now())

	elves := make(map[Vec2]struct{}, len(lines[0])*len(lines)/2)

	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				elves[Vec2{x: x, y: y}] = struct{}{}
			}
		}
	}

	for round := 0; round < 10; round++ {
		buf := make(map[Vec2]struct{}, len(elves))
		proposedMoves := make(map[Vec2][]Vec2, len(elves))

	ELVES:
		for elf := range elves {
			_, n := elves[Vec2{x: elf.x, y: elf.y - 1}]
			_, ne := elves[Vec2{x: elf.x + 1, y: elf.y - 1}]
			_, nw := elves[Vec2{x: elf.x - 1, y: elf.y - 1}]
			_, s := elves[Vec2{x: elf.x, y: elf.y + 1}]
			_, sw := elves[Vec2{x: elf.x - 1, y: elf.y + 1}]
			_, se := elves[Vec2{x: elf.x + 1, y: elf.y + 1}]
			_, w := elves[Vec2{x: elf.x - 1, y: elf.y}]
			_, e := elves[Vec2{x: elf.x + 1, y: elf.y}]

			if !n && !ne && !nw && !s && !sw && !se && !w && !e {
				buf[elf] = struct{}{}
				continue ELVES
			}

			for i := 0; i < 4; i++ {
				if dir, found := conditions[(round+i)%4](elf, n, ne, nw, s, sw, se, w, e); found {
					if _, found := proposedMoves[dir]; found {
						proposedMoves[dir] = append(proposedMoves[dir], elf)
					} else {
						proposedMoves[dir] = make([]Vec2, 1, 8)
						proposedMoves[dir][0] = elf
					}
					continue ELVES
				}
			}

			// Elf couldn't move
			buf[elf] = struct{}{}
		}

		for move, elves := range proposedMoves {
			if len(elves) == 1 {
				buf[move] = struct{}{}
			} else {
				for _, elf := range elves {
					buf[elf] = struct{}{}
				}
			}
		}

		if reflect.DeepEqual(elves, buf) {
			break
		}
		elves = buf
	}

	var minX, minY, maxX, maxY int
	for elf := range elves {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}
	log.Println("n empty tiles:", (maxX-minX+1)*(maxY-minY+1)-len(elves))
}
