package main

import (
	"fmt"
	"time"
)

var packages = []int{1, 2, 3, 7, 11, 13, 17, 19, 23, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}

// var packages = []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}

func main() {

	// groups := 3
	groups := 4
	groupSize := 0
	for _, p := range packages {
		groupSize += p
	}
	fmt.Printf("groups=%d, groupSize=%d, totalSum=%d\n", groups, groupSize, groupSize/groups)
	groupSize /= groups

	minPackages := len(packages) + 1
	qe := uint64(0)

	start := time.Now()
	group1 := make([]int, len(packages))
	sum1, size1 := 0, 0
	var packGroup1 func(int)
	packGroup1 = func(j int) {
		for i := j; i < len(packages); i++ {
			p := packages[i]
			sum1 += p
			group1[size1] = p
			size1++
			if sum1 < groupSize {
				packGroup1(i + 1)
			} else if sum1 == groupSize && size1 < minPackages {
				remaining := make([]int, 0)
			ITER_PACKAGES:
				for _, q := range packages {
					for _, w := range group1[:size1] {
						if q == w {
							// Group 1 contains package
							continue ITER_PACKAGES
						}
					}
					remaining = append(remaining, q)
				}

				/*****************************************/
				sum2, size2 := 0, 0
				var packGroup2 func(int)
				packGroup2 = func(j int) {
					for i := j; i < len(remaining); i++ {
						p := remaining[i]
						sum2 += p
						size2++
						if sum2 < groupSize {
							packGroup2(i + 1)
						} else if sum2 == groupSize {
							prod := uint64(1)
							for k := 0; k < size1; k++ {
								prod *= uint64(group1[k])
							}
							qe = prod
							minPackages = size1
							// fmt.Printf("qe=%d\n", qe)
							return
						}
						size2--
						sum2 -= p
					}
				}
				packGroup2(0)
				/*****************************************/

				// prod := uint64(1)
				// for k := 0; k < size1; k++ {
				// 	prod *= uint64(group1[k])
				// }
				// qe = prod
				// minPackages = size1

				// fmt.Printf("group1=%v remaining=%v\n", group1[:size1], remaining)
			}
			size1--
			sum1 -= p
		}
	}
	packGroup1(0)
	fmt.Printf("Day 24 took %v\n", time.Since(start))
	fmt.Printf("minPackages=%d, qe=%d\n", minPackages, qe)
}
