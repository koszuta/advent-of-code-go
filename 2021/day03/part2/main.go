package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

const expectedResult = 4406844

/*
 *   --- Day 3: Binary Diagnostic ---
 *           --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/3#part2
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	oxygenGeneratorRating, co2ScrubberRating := doPart2()
	log.Println("oxygen generator rating:", oxygenGeneratorRating)
	log.Println("CO2 scrubber rating:", co2ScrubberRating)
	log.Println("life support rating:", oxygenGeneratorRating*co2ScrubberRating)
}

func doPart2() (int64, int64) {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	nums := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}
	nums2 := make([]string, len(nums))
	copy(nums2, nums)

	doFindRating := func(nums []string, rule func(int, map[int]int, map[int]int) rune) (rating int64) {
		for i := 0; i < len(nums[0]); i++ {
			zeros := make(map[int]int)
			ones := make(map[int]int)
			for _, num := range nums {
				for i, c := range num {
					if c == '0' {
						zeros[i]++
					} else {
						ones[i]++
					}
				}
			}

			bit := rule(i, ones, zeros)

			for j := 0; j < len(nums); {
				if rune(nums[j][i]) != bit {
					nums = append(nums[:j], nums[j+1:]...)
				} else {
					j++
				}
			}
			if len(nums) == 1 {
				break
			}
		}
		rating, _ = strconv.ParseInt(nums[0], 2, 64)
		return
	}

	oxygenGeneratorRating := doFindRating(nums, func(i int, ones, zeros map[int]int) (bit rune) {
		if ones[i] >= zeros[i] {
			return '1'
		} else {
			return '0'
		}
	})

	co2ScrubberRating := doFindRating(nums2, func(i int, ones, zeros map[int]int) (bit rune) {
		if zeros[i] <= ones[i] {
			return '0'
		} else {
			return '1'
		}
	})

	return oxygenGeneratorRating, co2ScrubberRating
}
