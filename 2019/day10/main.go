package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
)

func main() {
	fileName := flag.String("file", "input.txt", "relative path of input file")
	flag.Parse()

	asteroids := make([]string, 0, 0)
	{
		file, err := os.Open(*fileName)
		if err != nil {
			log.Panicln(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			asteroids = append(asteroids, scanner.Text())
		}
	}

	for _, r := range asteroids {
		log.Println(r)
	}

	var maxAsteroid map[float64][]float64
	max, maxX, maxY := 0, 0, 0
	for y, r := range asteroids {
		for x, a := range r {
			if a == '#' {
				los := make(map[float64][]float64)
				for w, rr := range asteroids {
					for v, aa := range rr {
						if (v != x || w != y) && aa == '#' {
							a, b := float64(v-x), float64(w-y)
							θ := math.Atan2(b, a) + (math.Pi / 2)
							dist := math.Sqrt(float64(a*a + b*b))
							los[θ] = append(los[θ], dist)
						}
					}
				}
				// log.Printf("(%d,%d) %v\n", x, y, los)
				asteroidsSeen := len(los)
				if asteroidsSeen > max {
					max = asteroidsSeen
					maxX = x
					maxY = y
					maxAsteroid = los
				}
			}
		}
	}
	log.Printf("%d asteroids seen from (%d, %d)\n", max, maxX, maxY)

	loss := make([]float64, 0, len(maxAsteroid))
	for k, v := range maxAsteroid {
		sort.Float64s(v)
		loss = append(loss, k)
	}
	sort.Slice(loss, func(i, j int) bool {
		q, r := loss[i], loss[j]
		if q < 0 && r >= 0 {
			return false
		}
		if q >= 0 && r < 0 {
			return true
		}
		return q < r
	})

	count := 0
	for {
		for _, los := range loss {
			if len(maxAsteroid[los]) != 0 {
				count++
				if count == http.StatusOK {
					θ, r := los-(math.Pi/2), maxAsteroid[los][0]
					log.Printf("θ=%f r=%f\n", θ, r)
					x, y := maxX+int(r*math.Cos(θ)), maxY+int(r*math.Sin(θ))
					log.Printf("200th asteroid vaporized (%d,%d)\n", x, y)
					log.Printf("ans=%d\n", 100*x+y)
					os.Exit(0)
				}
				maxAsteroid[los] = maxAsteroid[los][1:]
			}
		}
	}
}
