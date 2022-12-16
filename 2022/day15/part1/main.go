package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const outputY = 2000000

var (
	lines []string

	inputRegex = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
)

type Vec2 struct {
	x, y int
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

	sensorBeaconDistance := make(map[Vec2]int, len(lines))
	for _, line := range lines {
		matches := inputRegex.FindAllSubmatch([]byte(line), -1)[0]
		sensorX, _ := strconv.Atoi(string(matches[1]))
		sensorY, _ := strconv.Atoi(string(matches[2]))
		beaconX, _ := strconv.Atoi(string(matches[3]))
		beaconY, _ := strconv.Atoi(string(matches[4]))

		sensor, beacon := Vec2{sensorX, sensorY}, Vec2{beaconX, beaconY}
		sensorBeaconDistance[sensor] = ManhattanDistance(sensor, beacon)
	}

	minX, maxX := 999999999, -999999999
	for sensor, beaconDistance := range sensorBeaconDistance {
		offset := beaconDistance - Abs(outputY-sensor.y)
		if offset >= 0 {
			if sensor.x-offset < minX {
				minX = sensor.x - offset
			}
			if sensor.x+offset > maxX {
				maxX = sensor.x + offset
			}
		}
	}
	log.Println("n position cannot contain a beacon:", maxX-minX)
}

func ManhattanDistance(from, to Vec2) int {
	return Abs(to.x-from.x) + Abs(to.y-from.y)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
