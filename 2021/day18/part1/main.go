package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

const expectedResult = 0

/*
 *   --- Day 18:  ---
 *      --- Part One ---
 *
 *   https://adventofcode.com/2021/day/18
 */

type SFN struct {
	left, right       int
	leftSFN, rightSFN *SFN
	depth             int
}

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	x := doPart1()
	log.Println(x)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sum *SFN
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		log.Println(line)
		sfn := parseSnailfishNumber(line)
		log.Println(sfn)
		if sum == nil {
			sum = sfn
		} else {
			sum = sum.add(sfn)
		}
		log.Println(sum)
		log.Println()
	}

	sfn0 := SFN{left: 9, right: 1}
	log.Println(sfn0.magnitude())
	sfn1 := SFN{left: 1, right: 2}
	sfn2 := SFN{left: 3, right: 4}
	sfn3 := SFN{leftSFN: &sfn2, right: 5}
	sfn4 := SFN{leftSFN: &sfn1, rightSFN: &sfn3}
	log.Println(sfn4.magnitude())

	return 0
}

func parseSnailfishNumber(line string) *SFN {
	sfn := SFN{}

	return &sfn
}

func (sfn *SFN) add(otherSFN *SFN) *SFN {
	newSFN := SFN{leftSFN: sfn, rightSFN: otherSFN}
	return newSFN.reduce()
}

func (sfn *SFN) reduce() *SFN {
	for {
		if sfn, didExplode := sfn.explode(); !didExplode {
			if sfn, didSplit := sfn.split(); !didSplit {
				return sfn
			}
		}
	}
}

func (sfn *SFN) explode() (*SFN, bool) {
	if sfn.rightSFN == nil && sfn.leftSFN != nil && sfn.leftSFN.isBasePair() && sfn.leftSFN.depth > 4 {

		return sfn, true
	}
	if sfn.leftSFN == nil && sfn.rightSFN != nil && sfn.rightSFN.isBasePair() && sfn.rightSFN.depth > 4 {

		return sfn, true
	}
	return sfn, false
}

func (sfn *SFN) split() (*SFN, bool) {

	return sfn, true
}

func (sfn *SFN) isBasePair() bool {
	return sfn.leftSFN == nil && sfn.rightSFN == nil
}

func (sfn *SFN) magnitude() int {
	var leftMag int
	if sfn.leftSFN != nil {
		leftMag = sfn.leftSFN.magnitude()
	} else {
		leftMag = sfn.left
	}
	var rightMag int
	if sfn.rightSFN != nil {
		rightMag = sfn.rightSFN.magnitude()
	} else {
		rightMag = sfn.right
	}
	return 3*leftMag + 2*rightMag
}
