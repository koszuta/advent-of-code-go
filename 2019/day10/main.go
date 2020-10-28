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
	for y, row := range asteroids {
		for x, asteroid := range row {
			if asteroid == '#' {
				θs := make(map[float64][]float64)
				for w, row := range asteroids {
					for v, asteroid := range row {
						if (v != x || w != y) && asteroid == '#' {
							a, b := float64(v-x), float64(w-y)
							θ := math.Atan2(b, a) + (math.Pi / 2)
							r := math.Sqrt(float64(a*a + b*b))
							θs[θ] = append(θs[θ], r)
						}
					}
				}
				if len(θs) > max {
					max = len(θs)
					maxX = x
					maxY = y
					maxAsteroid = θs
				}
			}
		}
	}
	log.Printf("%d asteroids seen from (%d, %d)\n", max, maxX, maxY)

	θs := make([]float64, 0, len(maxAsteroid))
	for θ, r := range maxAsteroid {
		sort.Float64s(r)
		θs = append(θs, θ)
	}
	sort.Slice(θs, func(i, j int) bool {
		θ1, θ2 := θs[i], θs[j]
		if θ1 < 0 && θ2 >= 0 {
			return false
		}
		if θ1 >= 0 && θ2 < 0 {
			return true
		}
		return θ1 < θ2
	})

	count := 0
	for {
		for _, θ := range θs {
			if len(maxAsteroid[θ]) != 0 {
				count++
				if count == http.StatusOK {
					θ, r := θ-(math.Pi/2), maxAsteroid[θ][0]
					log.Printf("θ=%f r=%f\n", θ, r)
					x, y := maxX+int(r*math.Cos(θ)), maxY+int(r*math.Sin(θ))
					log.Printf("200th asteroid vaporized is (%d, %d)\n", x, y)
					log.Printf("ans=%d\n", 100*x+y)
					os.Exit(0)
				}
				maxAsteroid[θ] = maxAsteroid[θ][1:]
			}
		}
	}
}
