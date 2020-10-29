package main

import "log"

type moon struct {
	pos, vel vec3i
}

type vec3i struct {
	x, y, z int
}

var moons = [4]moon{
	moon{
		pos: vec3i{6, 10, 10},
		vel: vec3i{0, 0, 0},
	},
	moon{
		pos: vec3i{-9, 3, 17},
		vel: vec3i{0, 0, 0},
	},
	moon{
		pos: vec3i{9, -4, 14},
		vel: vec3i{0, 0, 0},
	},
	moon{
		pos: vec3i{4, 14, 4},
		vel: vec3i{0, 0, 0},
	},
}

// var moons = [4]moon{
// 	moon{
// 		pos: vec3i{-1, 0, 2},
// 		vel: vec3i{0, 0, 0},
// 	},
// 	moon{
// 		pos: vec3i{2, -10, -7},
// 		vel: vec3i{0, 0, 0},
// 	},
// 	moon{
// 		pos: vec3i{4, -8, 8},
// 		vel: vec3i{0, 0, 0},
// 	},
// 	moon{
// 		pos: vec3i{3, 5, -1},
// 		vel: vec3i{0, 0, 0},
// 	},
// }

// var moons = [4]moon{
// 	moon{
// 		pos: vec3i{-8, -10, 0},
// 		vel: vec3i{0, 0, 0},
// 	},
// 	moon{
// 		pos: vec3i{5, 5, 10},
// 		vel: vec3i{0, 0, 0},
// 	},
// 	moon{
// 		pos: vec3i{2, -7, 3},
// 		vel: vec3i{0, 0, 0},
// 	},
// 	moon{
// 		pos: vec3i{9, -8, -3},
// 		vel: vec3i{0, 0, 0},
// 	},
// }

type hash struct {
	a, b, c, d, e, f, g, h int
}

func intAbs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func primeFactorCounts(n int) map[int]int {
	factors := make(map[int]int)
	p := 2
	for n < -1 || n > 1 {
		if n%p == 0 {
			n /= p
			factors[p] = factors[p] + 1
		} else {
			p++
		}
	}
	return factors
}

func lCM(nums ...int) int {
	commonFactors := make(map[int]int)
	for _, n := range nums {
		factors := primeFactorCounts(n)
		for prime, count := range factors {
			if count > commonFactors[prime] {
				commonFactors[prime] = count
			}
		}
	}

	lcm := 1
	for p, c := range commonFactors {
		for i := 0; i < c; i++ {
			lcm *= p
		}
	}
	return lcm
}

func main() {

	xStates := make(map[hash]struct{})
	yStates := make(map[hash]struct{})
	zStates := make(map[hash]struct{})

	steps := 0
	xCycle, yCycle, zCycle := 0, 0, 0

	// Repeat until all component cycles have been found
	for xCycle == 0 || yCycle == 0 || zCycle == 0 {

		// Check if the x cycle has been found alrady
		if xCycle == 0 {
			// Hash the x components of all moon positions and velocities
			xHash := hash{moons[0].pos.x, moons[0].vel.x, moons[1].pos.x, moons[1].vel.x, moons[2].pos.x, moons[2].vel.x, moons[3].pos.x, moons[3].vel.x}
			_, found := xStates[xHash]
			if !found {
				// Store it if it hasn't been encountered before
				xStates[xHash] = struct{}{}
			} else {
				// Otherwise, the state of the x component after this many steps
				xCycle = steps
			}
		}
		if yCycle == 0 {
			yHash := hash{moons[0].pos.y, moons[0].vel.y, moons[1].pos.y, moons[1].vel.y, moons[2].pos.y, moons[2].vel.y, moons[3].pos.y, moons[3].vel.y}
			_, found := yStates[yHash]
			if !found {
				yStates[yHash] = struct{}{}
			} else {
				yCycle = steps
			}
		}
		if zCycle == 0 {
			zHash := hash{moons[0].pos.z, moons[0].vel.z, moons[1].pos.z, moons[1].vel.z, moons[2].pos.z, moons[2].vel.z, moons[3].pos.z, moons[3].vel.z}
			_, found := zStates[zHash]
			if !found {
				zStates[zHash] = struct{}{}
			} else {
				zCycle = steps
			}
		}

		for i := range moons {
			for j := i + 1; j < len(moons); j++ {
				p1, p2 := moons[i].pos, moons[j].pos
				v1, v2 := &moons[i].vel, &moons[j].vel
				if p1.x < p2.x {
					v1.x++
					v2.x--
				} else if p1.x > p2.x {
					v1.x--
					v2.x++
				}
				if p1.y < p2.y {
					v1.y++
					v2.y--
				} else if p1.y > p2.y {
					v1.y--
					v2.y++
				}
				if p1.z < p2.z {
					v1.z++
					v2.z--
				} else if p1.z > p2.z {
					v1.z--
					v2.z++
				}
			}
		}
		for i := range moons {
			moon := &moons[i]
			moon.pos.x += moon.vel.x
			moon.pos.y += moon.vel.y
			moon.pos.z += moon.vel.z
		}
		steps++
	}

	// Part 1
	totalEnergy := 0
	for _, m := range moons {
		pot := intAbs(m.pos.x) + intAbs(m.pos.y) + intAbs(m.pos.z)
		kin := intAbs(m.vel.x) + intAbs(m.vel.y) + intAbs(m.vel.z)
		totalEnergy += pot * kin
	}
	log.Printf("total energy in the system is %d\n", totalEnergy)

	// Deuxième partie
	log.Printf("x component cycle is %d steps\n", xCycle)
	log.Printf("y component cycle is %d steps\n", yCycle)
	log.Printf("z component cycle is %d steps\n", zCycle)
	log.Printf("moons state repeated in %d steps\n", lCM(xCycle, yCycle, zCycle))
}
