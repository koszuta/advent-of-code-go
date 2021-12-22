package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const expectedResult = 0

/*
 *   --- Day 18: Snailfish ---
 *       --- Part One ---
 *
 *   https://adventofcode.com/2021/day/18
 */

type SFN struct {
	leftNum, rightNum   int
	parent, left, right *SFN
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
		sfn := parseSnailfishNumber(line, nil)
		if sum == nil {
			sum = sfn
		} else {
			sum = sum.add(sfn)
		}
	}

	fmt.Println("sum:", sum.toString())
	return sum.magnitude()
}

func parseSnailfishNumber(line string, parent *SFN) *SFN {
	sfn := SFN{}
	line = line[1 : len(line)-1]
	openings := 0
	i := 0
OUT:
	for _, r := range line {
		switch r {
		case '[':
			openings++
		case ']':
			openings--
		case ',':
			if openings == 0 {
				break OUT
			}
		}
		i++
	}

	leftStr := line[:i]
	if leftStr[0] == '[' {
		sfn.left = parseSnailfishNumber(leftStr, &sfn)
	} else {
		left, _ := strconv.Atoi(leftStr)
		sfn.leftNum = left
	}
	rightStr := line[i+1:]
	if rightStr[0] == '[' {
		sfn.right = parseSnailfishNumber(rightStr, &sfn)
	} else {
		right, _ := strconv.Atoi(rightStr)
		sfn.rightNum = right
	}
	return &sfn
}

func (sfn *SFN) toString() string {
	var leftStr string
	if sfn.left == nil {
		leftStr = strconv.Itoa(sfn.leftNum)
	} else {
		leftStr = sfn.left.toString()
	}
	var rightStr string
	if sfn.right == nil {
		rightStr = strconv.Itoa(sfn.rightNum)
	} else {
		rightStr = sfn.right.toString()
	}
	return fmt.Sprintf("[%s,%s]", leftStr, rightStr)
}

func (sfn *SFN) add(otherSFN *SFN) *SFN {
	newSFN := SFN{left: sfn, right: otherSFN}
	fmt.Println("after addition:", newSFN.toString())
	newSFN.reduce()
	return &newSFN
}

func (sfn *SFN) reduce() {
	for {
		if !sfn.explode(0) {
			if !sfn.split() {
				fmt.Println()
				return
			} else {
				fmt.Println("after split:   ", sfn.toString())
			}
		} else {
			fmt.Println("after explode: ", sfn.toString())
		}
	}
}

func (sfn *SFN) explode(depth int) bool {
	if depth >= 4 {
		return true
	}
	return false
}

func (sfn *SFN) split() (split bool) {
	if sfn.left != nil && sfn.left.split() {
		return true
	}
	if sfn.leftNum > 9 {
		newPair := SFN{leftNum: sfn.leftNum / 2, rightNum: (sfn.leftNum + 1) / 2}
		sfn.leftNum = 0
		sfn.left = &newPair
		return true
	}
	if sfn.rightNum > 9 {
		newPair := SFN{leftNum: sfn.rightNum / 2, rightNum: (sfn.rightNum + 1) / 2}
		sfn.rightNum = 0
		sfn.right = &newPair
		return true
	}
	if sfn.right != nil && sfn.right.split() {
		return true
	}
	return false
}

func (sfn *SFN) magnitude() int {
	var leftMag int
	if sfn.left == nil {
		leftMag = sfn.leftNum
	} else {
		leftMag = sfn.left.magnitude()
	}
	var rightMag int
	if sfn.right == nil {
		rightMag = sfn.rightNum
	} else {
		rightMag = sfn.right.magnitude()
	}
	return 3*leftMag + 2*rightMag
}
