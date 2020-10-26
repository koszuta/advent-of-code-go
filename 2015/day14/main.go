package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
	name string
	speed, duration, rest, distance, score int
}

const time = 2503

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reindeers := make([]reindeer, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var speed, duration, rest int
		var err error
		speed, err = strconv.Atoi(line[3])
		if err != nil {
			panic(err)
		}
		duration, err = strconv.Atoi(line[6])
		if err != nil {
			panic(err)
		}
		rest, err = strconv.Atoi(line[13])
		if err != nil {
			panic(err)
		}
		reindeers = append(reindeers, reindeer { line[0], speed, duration, rest, 0, 0 })
	}

	for i := range reindeers {
		reindeer := &reindeers[i]
		s := reindeer.speed
		d := reindeer.duration
		r := reindeer.rest
		t := time / (d + r)
		reindeer.distance = (t * s * d) + (s * int(math.Min(float64(d), float64(time - (t * (d + r))))))
	}
	fmt.Printf("reindeers=%v\n", reindeers)

	furthest := 0
	for _, reindeer := range reindeers {
		if reindeer.distance > furthest {
			furthest = reindeer.distance
		}
	}
	fmt.Printf("furthest=%d\n", furthest)

	for i := range reindeers {
		r := &reindeers[i]
		r.distance = 0
	}
	for t := 0; t < time; t++ {
		// fmt.Printf("time=%d\n", t)
		for i := range reindeers {
			r := &reindeers[i]
			if t % (r.duration + r.rest) < r.duration {
				r.distance += r.speed
				// fmt.Printf("%s traveled a total of %d km\n", r.name, r.distance)
			} else {
				// fmt.Printf("%s is resting at %d km\n", r.name, r.distance)
			}
		}
		lead := 0
		for _, r := range reindeers {
			if r.distance > lead {
				lead = r.distance
			}
		}
		for i := range reindeers {
			r := &reindeers[i]
			if r.distance == lead {
				r.score++
			}
		}
	}

	highestScore := 0
	for _, reindeer := range reindeers {
		if reindeer.score > highestScore {
			highestScore = reindeer.score
		}
	}
	fmt.Printf("highestScore=%d\n", highestScore)
}
