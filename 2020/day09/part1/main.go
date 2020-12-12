package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
 *   --- Day 9: Encoding Error ---
 *         --- Part One ---
 *
 *   https://adventofcode.com/2020/day/9
 */

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

	// Find the first number which isn't the sum of two numbers
	// from the 25 numbers immediately preceding it
OUTER:
	for i := 25; i < len(nums); i++ {
		num := nums[i]
		sum := make(map[int]struct{})
		for j := 1; j < 26; j++ {
			num2 := nums[i-j]
			_, sumExists := sum[num-num2]
			if sumExists {
				continue OUTER
			}
			sum[num2] = struct{}{}
		}
		log.Printf("the first invalid number is %d\n", num)
		break
	}
}
