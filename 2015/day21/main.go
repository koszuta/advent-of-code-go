package main

import (
	"fmt"
	"math"
)

type item struct {
	cost, damage, armor int
}

type character struct {
	hp, damage, armor int
}

var weapons = [...]item{item{8, 4, 0}, item{10, 5, 0}, item{25, 6, 0}, item{40, 7, 0}, item{74, 8, 0}}
var armors = [...]item{item{}, item{13, 0, 1}, item{31, 0, 2}, item{53, 0, 3}, item{75, 0, 4}, item{102, 0, 5}}
var rings = [...]item{item{}, item{25, 1, 0}, item{50, 2, 0}, item{100, 3, 0}, item{20, 0, 1}, item{40, 0, 2}, item{80, 0, 3}}

func (c1 *character) fight(c2 character) bool {
	for {
		c2.hp -= int(math.Max(1.0, float64(c1.damage-c2.armor)))
		if c2.hp <= 0 {
			return true
		}
		c1.hp -= int(math.Max(1.0, float64(c2.damage-c1.armor)))
		if c1.hp <= 0 {
			return false
		}
	}
}

func main() {
	minCost, maxCost := 357, 0
	{
		for _, weapon := range weapons {
			for _, armor := range armors {
				for r, ring1 := range rings {
					for _, ring2 := range rings[r:] {
						if ring1 == ring2 {
							ring2 = rings[0]
						}

						cost := weapon.cost + armor.cost + ring1.cost + ring2.cost
						damage := weapon.damage + armor.damage + ring1.damage + ring2.damage
						armor := weapon.armor + armor.armor + ring1.armor + ring2.armor

						player := character{100, damage, armor}
						boss := character{109, 8, 2}

						playerWon := player.fight(boss)
						if playerWon && cost < minCost {
							minCost = cost
						}
						if !playerWon && cost > maxCost {
							maxCost = cost
						}
					}
				}
			}
		}
	}
	fmt.Printf("minCost and win=%d\n", minCost)
	fmt.Printf("maxCost and lose=%d\n", maxCost)
}
