package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Answer from Part One
var answer = 1038347917

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

	// Find a contiguous set of numbers which sum to the answer from Part One
	sum := 0
	set := make([]int, 0, 0)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		set = append(set, nums[i])

		// Remove numbers from the front of the set until the sum is not greater than the answer
		var j int
		for j = 0; sum > answer; j++ {
			sum -= set[j]
		}
		if j < len(set) {
			set = set[j:]
		} else {
			set = make([]int, 0, 0)
		}

		if sum == answer {
			break
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
