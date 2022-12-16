package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const limit = 4000000

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

	var beacon Vec2
	for y := 0; y <= limit; y++ {
		segments := make([]Vec2, 0)
		for sensor, beaconDistance := range sensorBeaconDistance {
			offset := beaconDistance - Abs(y-sensor.y)
			if offset >= 0 {
				segment := Vec2{x: sensor.x - offset, y: sensor.x + offset}
				segments = MergeSegments(append(segments, segment))
			}
		}
		if len(segments) == 2 {
			beacon.x = (segments[0].y + segments[1].x) / 2
			beacon.y = y
			break
		}
	}
	log.Println("tuning frequency:", beacon.x*4000000+beacon.y)
}

func MergeSegments(segments []Vec2) []Vec2 {
	sort.Slice(segments, func(i, j int) bool {
		if segments[i].x == segments[j].x {
			return segments[i].y < segments[j].y
		}
		return segments[i].x < segments[j].x
	})
	for i := 0; i < len(segments); i++ {
		for j := i + 1; j < len(segments); j++ {
			if segments[i].y >= segments[j].x {
				segments[i].y = Max(segments[i].y, segments[j].y)
				segments = append(segments[:j], segments[j+1:]...)
				j--
			}
		}
	}
	return segments
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

func Max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
