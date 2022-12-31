package main

import (
	"log"
	"os"
	"strings"
)

var (
	w, h      int
	blizzards map[Vec2][]rune
)

type Vec2 struct {
	x, y int
}

var lines []string

func init() {
	b, _ := os.ReadFile("../input.txt")
	lines = strings.Split(string(b), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line) // sanitize CRLF
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-2]
	}
}

func main() {
	lines = lines[1 : len(lines)-1]

	w, h = len(lines[0])-2, len(lines)
	blizzards = make(map[Vec2][]rune, w*h)

	for y, line := range lines {
		line = strings.Trim(line, "#")

		for x, r := range line {
			if r != '.' {
				pos := Vec2{x: x, y: y}
				if b, found := blizzards[pos]; found {
					blizzards[pos] = append(b, r)
				} else {
					blizzards[pos] = make([]rune, 1, 4)
					blizzards[pos][0] = r
				}
			}
		}
	}

	// First round of blizzards to
	buf := make(map[Vec2][]rune, w*h)

	for pos, dirs := range blizzards {
		for _, dir := range dirs {
			var newPos Vec2
			switch dir {
			case '<':
				newPos = Vec2{x: pos.x - 1, y: pos.y}
				if newPos.x < 0 {
					newPos.x = w - 1
				}
			case '>':
				newPos = Vec2{x: pos.x + 1, y: pos.y}
				if newPos.x >= w {
					newPos.x = 0
				}
			case '^':
				newPos = Vec2{x: pos.x, y: pos.y - 1}
				if newPos.y < 0 {
					newPos.y = h - 1
				}
			case 'v':
				newPos = Vec2{x: pos.x, y: pos.y + 1}
				if newPos.y >= h {
					newPos.y = 0
				}
			default:
				log.Fatalln("unknown direction", dir)
			}
			var b []rune
			var found bool
			if b, found = buf[newPos]; found {
				b = append(b, dir)
			} else {
				b = make([]rune, 1, 4)
				b[0] = dir
			}
			buf[newPos] = b
		}
	}

	blizzards = buf

	pos := Vec2{x: 0, y: -1}
	for t := 0; ; t++ {
	}
}

func DoWalk(t, step int, pos Vec2) {
	if pos.x == w && pos.y == h {
		return step + 1
	}

	buf := make(map[Vec2][]rune, w*h)

	blizzards = buf
}
