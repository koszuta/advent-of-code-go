package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var lines []string

type Direction int

const (
	Right = Direction(iota)
	Down
	Left
	Up
)

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

	tiles := make([]string, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			break
		}
		tiles = append(tiles, line)
	}

	s := lines[len(lines)-1]
	s = strings.ReplaceAll(s, "L", ".L.")
	s = strings.ReplaceAll(s, "R", ".R.")
	path := strings.Split(s, ".")

	x, y, dir := strings.IndexRune(tiles[0], '.'), 0, Right
	if x == -1 {
		log.Fatalln("couldn't find starting position")
	}
PATH:
	for _, p := range path {
		switch p {
		case "L":
			dir = (dir + 1) % 4
		case "R":
			dir = (dir + 3) % 4
		default:
			distance, _ := strconv.Atoi(p)
			w, h := len(tiles[y]), len(tiles)
			nextX, nextY := x, y
			var wrap int
			switch dir {
			case Right:
				for step := 0; step < distance; step++ {
					if x == w-1 {
						wrap = strings.LastIndex(tiles[y], " ") + 1
					}

					nextX = (x + wrap + 1) % w
					if tiles[y][nextX] == ' ' {
						log.Fatalln("invalid character at x=", nextX)
					}
					if tiles[y][nextX] == '#' {
						continue PATH
					}

					x = (x + 1) % w
				}
			case Down:
				for step := 0; step < distance; step++ {
					if y == 0 || tiles[y-1][x] == ' ' {
						for v := h - 1; len(tiles[v]) <= x || tiles[v][x] == ' '; v, wrap = v-1, wrap+1 {
						}
						wrap += y
					}

					nextY = (y - wrap - 1 + h) % h
					if tiles[nextY][x] == ' ' {
						log.Fatalln("invalid character at y=", nextY)
					}
					if tiles[nextY][x] == '#' {
						continue PATH
					}

					if wrap > 0 {
						y = nextY
						wrap = 0
					} else {
						y = (y - 1 + h) % h
					}
				}
			case Left:
				for step := 0; step < distance; step++ {
					if x == 0 || tiles[y][x-1] == ' ' {
						wrap = x
					}

					nextX := (x - wrap - 1 + w) % w
					if tiles[y][nextX] == ' ' {
						log.Fatalln("invalid character at x=", nextX)
					}
					if tiles[y][nextX] == '#' {
						continue PATH
					}

					if wrap > 0 {
						x = nextX
						wrap = 0
					} else {
						x = (x - 1 + w) % w
					}
				}
			case Up:
				for step := 0; step < distance; step++ {
					if y == h-1 || len(tiles[y+1]) <= x || tiles[y+1][x] == ' ' {
						for ; len(tiles[wrap]) <= x || tiles[wrap][x] == ' '; wrap++ {
						}
						wrap += h - y - 1
					}

					nextY = (y + wrap + 1) % h
					if tiles[nextY][x] == ' ' {
						log.Fatalln("invalid character at y=", nextY)
					}
					if tiles[nextY][x] == '#' {
						continue PATH
					}

					if wrap > 0 {
						y = nextY
						wrap = 0
					} else {
						y = (y + 1) % h
					}
				}
			}
		}
	}

	log.Println("final password:", (y+1)*1000+(x+1)*4+int(dir))
}
