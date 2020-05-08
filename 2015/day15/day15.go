package main

import (
	"fmt"
)

type ingredient struct {
	name string
	capacity, durability, flavor, texture, calories int
}

const totalTsp = 100

func main() {
	ingredients := [4]ingredient{
		ingredient { "Sprinkles", 5, -1, 0, 0, 5 },
		ingredient { "PeanutButter", -1, 3, 0, 0, 1 },
		ingredient { "Frosting", 0, -1, 4, 0, 6 },
		ingredient { "Sugar", -1, 0, 0, 2, 8 },
	}
	scores := make([]int, 0)

	for i := 0; i <= totalTsp; i++ {
		c1 := i * ingredients[0].capacity
		d1 := i * ingredients[0].durability
		f1 := i * ingredients[0].flavor
		t1 := i * ingredients[0].texture
		C1 := i * ingredients[0].calories
		for j := 0; i + j <= totalTsp; j++ {
			c2 := j * ingredients[1].capacity
			d2 := j * ingredients[1].durability
			f2 := j * ingredients[1].flavor
			t2 := j * ingredients[1].texture
			C2 := j * ingredients[1].calories
			for k := 0; i + j + k <= totalTsp; k++ {
				c3 := k * ingredients[2].capacity
				d3 := k * ingredients[2].durability
				f3 := k * ingredients[2].flavor
				t3 := k * ingredients[2].texture
				C3 := k * ingredients[2].calories

				l := totalTsp - (i + j + k)
				c4 := l * ingredients[3].capacity
				d4 := l * ingredients[3].durability
				f4 := l * ingredients[3].flavor
				t4 := l * ingredients[3].texture
				C4 := l * ingredients[3].calories

				c := c1 + c2 + c3 + c4
				if c < 0 {
					c = 0
				}
				d := d1 + d2 + d3 + d4
				if d < 0 {
					d = 0
				}
				f := f1 + f2 + f3 + f4
				if f < 0 {
					f = 0
				}
				t := t1 + t2 + t3 + t4
				if t < 0 {
					t = 0
				}
				C := C1 + C2 + C3 + C4
				
				if C == 500 {
					score := c * d * f * t
					scores = append(scores, score)
				}
			}
		}
	}

	highestScore := 0
	for _, score := range scores {
		if score > highestScore {
			highestScore = score
		}
	}
	fmt.Printf("highestScore=%d\n", highestScore)
}
