package main

import (
	"advent-of-code-go/2021/day18/sfn"
	"bufio"
	"log"
	"os"
	"time"
)

const expectedResult = 4650

/*
 *   --- Day 18: Snailfish ---
 *       --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/18#part2
 */

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	mag := doPart2()
	log.Println("the largest magnitude of any sum of two different snailfish numbers", mag)
}

func doPart2() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sfns := make([]*sfn.SnailfishNumber, 0)
	for scanner.Scan() {
		sfn := sfn.ParseSnailfishNumber(scanner.Text(), nil)
		sfns = append(sfns, sfn)
	}

	maxMag := 0
	for i := 0; i < len(sfns); i++ {
		for j := i + 1; j < len(sfns); j++ {
			this, that := sfns[i].DeepCopy(), sfns[j].DeepCopy()
			mag := this.Add(that).Magnitude()
			if mag > maxMag {
				maxMag = mag
			}
			this, that = sfns[i].DeepCopy(), sfns[j].DeepCopy()
			mag = that.Add(this).Magnitude()
			if mag > maxMag {
				maxMag = mag
			}
		}
	}
	return maxMag
}
