package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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
		log.Printf("the first invalid number is %d\n", num)
		break
	}
}
