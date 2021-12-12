package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 1010472

/*
 *   --- Day 8: Seven Segment Search ---
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/8#part2
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	sum := doPart2()
	log.Println("the sum of the output values is", sum)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		inputSegments := strings.Split(parts[0], " ")
		for i, s := range inputSegments {
			inputSegments[i] = alphabetize(s)
		}
		outputSegments := strings.Split(parts[1], " ")
		for i, s := range outputSegments {
			outputSegments[i] = alphabetize(s)
		}

		allSegments := make([]string, 0, len(inputSegments)+len(outputSegments))
		allSegments = append(allSegments, inputSegments...)
		allSegments = append(allSegments, outputSegments...)

		segmentsMap := make(map[string]string)
		segmentsNumberMap := make(map[string]int)
		numberSegmentsMap := make(map[int]map[string]struct{})

		// Find 1, 4, 7, 8; each have unique segments
		for _, segs := range allSegments {
			var num int
			switch len(segs) {
			case 2:
				num = 1
			case 3:
				num = 7
			case 4:
				num = 4
			case 7:
				num = 8
			default:
				continue
			}
			numberSegmentsMap[num] = make(map[string]struct{})
			for _, seg := range segs {
				numberSegmentsMap[num][string(seg)] = struct{}{}
			}
			segmentsNumberMap[segs] = num
		}

		// Find 6; 6 segment digit which doesn't have the "c" segment
		for _, segs := range allSegments {
			if len(segs) == 6 {
				diff := difference(segs, numberSegmentsMap[1])
				if len(diff) == 1 {
					num := 6
					numberSegmentsMap[num] = make(map[string]struct{})
					for _, seg := range segs {
						numberSegmentsMap[num][string(seg)] = struct{}{}
					}
					segmentsNumberMap[segs] = num
					segmentsMap["c"] = diff[0]
					break
				}
			}
		}

		// Find 5; 5 segment digit which doesn't have the "e" segment
		for _, segs := range allSegments {
			if len(segs) == 5 {
				diff := difference(segs, numberSegmentsMap[6])
				if len(diff) == 1 {
					num := 5
					numberSegmentsMap[num] = make(map[string]struct{})
					for _, seg := range segs {
						numberSegmentsMap[num][string(seg)] = struct{}{}
					}
					segmentsNumberMap[segs] = num
					segmentsMap["e"] = diff[0]
					break
				}
			}
		}

		// Find 9; 6 segment digit which doesn't have the "e" segment
		for _, segs := range allSegments {
			if len(segs) == 6 && !strings.Contains(segs, segmentsMap["e"]) {
				num := 9
				numberSegmentsMap[num] = make(map[string]struct{})
				for _, seg := range segs {
					numberSegmentsMap[num][string(seg)] = struct{}{}
				}
				segmentsNumberMap[segs] = num
				break
			}
		}

		// Find 3; 5 segment digit which doesn't have the "b" segment
		for _, segs := range allSegments {
			if len(segs) == 5 {
				diff := difference(segs, numberSegmentsMap[9])
				if len(diff) == 1 {
					if _, found := segmentsNumberMap[segs]; !found {
						num := 3
						numberSegmentsMap[num] = make(map[string]struct{})
						for _, seg := range segs {
							numberSegmentsMap[num][string(seg)] = struct{}{}
						}
						segmentsNumberMap[segs] = num
						segmentsMap["b"] = diff[0]
						break
					}
				}
			}
		}

		// Find 2; last 5 segment digit
		for _, segs := range allSegments {
			if len(segs) == 5 {
				if _, found := segmentsNumberMap[segs]; !found {
					num := 2
					numberSegmentsMap[num] = make(map[string]struct{})
					for _, seg := range segs {
						numberSegmentsMap[num][string(seg)] = struct{}{}
					}
					segmentsNumberMap[segs] = num
					break
				}
			}
		}

		// Find 0; last 6 segment digit
		for _, segs := range allSegments {
			if len(segs) == 6 {
				if _, found := segmentsNumberMap[segs]; !found {
					num := 0
					numberSegmentsMap[num] = make(map[string]struct{})
					for _, seg := range segs {
						numberSegmentsMap[num][string(seg)] = struct{}{}
					}
					segmentsNumberMap[segs] = num
					break
				}
			}
		}

		var outputNum string
		for _, segs := range outputSegments {
			outputNum += strconv.Itoa(segmentsNumberMap[segs])
		}
		num, _ := strconv.Atoi(outputNum)
		sum += num
	}

	return sum
}

func difference(s1 string, s2 map[string]struct{}) []string {
	diff := make([]string, 0)
	for s := range s2 {
		if !strings.Contains(s1, s) {
			diff = append(diff, s)
		}
	}
	return diff
}

func alphabetize(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}
