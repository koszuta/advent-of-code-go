package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func partOne(nums []int) int {
OUTER:
	for i := 25; i < len(nums); i++ {
		num := nums[i]
		for j := i - 25; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if num == nums[j]+nums[k] {
					continue OUTER
				}
			}
		}
		return num
	}
	return 0
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	// Parse the input as a list of integers
	nums := make([]int, 0, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	// Get the answer from Part One
	answer := partOne(nums)

	// Find a contiguous set of numbers which sum to the answer from Part One
	sum := 0
	set := make([]int, 0, 0)
OUTER:
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			num := nums[j]
			if sum+num > answer {
				// If adding this number would go past the answer,
				// reset everything, and check from the next starting number
				sum = 0
				set = make([]int, 0, 0)
				continue OUTER
			} else {
				// Otherwise, add this number to the set
				sum += num
				set = append(set, num)

				// If we hit the answer exactly, stop checking
				if sum == answer {
					break OUTER
				}
			}
		}
	}

	// The "encryption weakness" is min+max from the set
	min, max := set[0], set[0]
	for _, num := range set {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	log.Printf("the encryption weakness is %d\n", min+max)
}
