package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var lines []string

type Vec3 struct {
	x, y, z float64
}

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
	defer func(t time.Time) {
		log.Println("took:", time.Since(t))
	}(time.Now())

	faces := make(map[Vec3]int, len(lines)*6)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		for _, face := range Faces(Vec3{x: float64(x), y: float64(y), z: float64(z)}) {
			faces[face]++
		}
	}

	for f, c := range faces {
		if c > 1 {
			delete(faces, f)
		}
	}
	log.Println("count:", len(faces))
}

func Faces(cube Vec3) [6]Vec3 {
	return [6]Vec3{
		{cube.x - 0.5, cube.y, cube.z},
		{cube.x + 0.5, cube.y, cube.z},
		{cube.x, cube.y - 0.5, cube.z},
		{cube.x, cube.y + 0.5, cube.z},
		{cube.x, cube.y, cube.z - 0.5},
		{cube.x, cube.y, cube.z + 0.5},
	}
}
