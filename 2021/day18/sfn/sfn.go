package sfn

import (
	"fmt"
	"strconv"
)

const logging = false

type SnailfishNumber struct {
	leftNum, rightNum   int
	parent, left, right *SnailfishNumber
}

func ParseSnailfishNumber(line string, parent *SnailfishNumber) *SnailfishNumber {
	sfn := SnailfishNumber{parent: parent}
	line = line[1 : len(line)-1]
	openings := 0
	var i int
	for i = range line {
		r := line[i]
		if r == '[' {
			openings++
		} else if r == ']' {
			openings--
		} else if r == ',' && openings == 0 {
			break
		}
	}

	leftStr := line[:i]
	if leftStr[0] == '[' {
		sfn.left = ParseSnailfishNumber(leftStr, &sfn)
	} else {
		left, _ := strconv.Atoi(leftStr)
		sfn.leftNum = left
	}
	rightStr := line[i+1:]
	if rightStr[0] == '[' {
		sfn.right = ParseSnailfishNumber(rightStr, &sfn)
	} else {
		right, _ := strconv.Atoi(rightStr)
		sfn.rightNum = right
	}
	return &sfn
}

func (sfn *SnailfishNumber) DeepCopy() *SnailfishNumber {
	return sfn.doDeepCopy(nil)
}
func (sfn *SnailfishNumber) doDeepCopy(parent *SnailfishNumber) *SnailfishNumber {
	newSFN := SnailfishNumber{
		parent:   parent,
		leftNum:  sfn.leftNum,
		rightNum: sfn.rightNum,
	}
	var newLeft, newRight *SnailfishNumber
	if sfn.left != nil {
		newLeft = sfn.left.doDeepCopy(&newSFN)
	}
	if sfn.right != nil {
		newRight = sfn.right.doDeepCopy(&newSFN)
	}
	newSFN.left = newLeft
	newSFN.right = newRight
	return &newSFN
}

func (sfn *SnailfishNumber) ToString() string {
	var leftStr string
	if sfn.left == nil {
		leftStr = strconv.Itoa(sfn.leftNum)
	} else {
		leftStr = sfn.left.ToString()
	}
	var rightStr string
	if sfn.right == nil {
		rightStr = strconv.Itoa(sfn.rightNum)
	} else {
		rightStr = sfn.right.ToString()
	}
	return fmt.Sprintf("[%s,%s]", leftStr, rightStr)
}

func (sfn *SnailfishNumber) Add(otherSFN *SnailfishNumber) *SnailfishNumber {
	newSFN := SnailfishNumber{left: sfn, right: otherSFN}
	sfn.parent = &newSFN
	otherSFN.parent = &newSFN
	if logging {
		fmt.Println("after addition:", newSFN.ToString())
	}
	newSFN.Reduce()
	return &newSFN
}

func (sfn *SnailfishNumber) Reduce() {
	for {
		if !sfn.Explode(0) {
			if !sfn.Split() {
				if logging {
					fmt.Println()
				}
				return
			} else if logging {
				fmt.Println("after split:   ", sfn.ToString())

			}
		} else if logging {
			fmt.Println("after explode: ", sfn.ToString())
		}
	}
}

func (sfn *SnailfishNumber) Explode(depth int) bool {
	if depth >= 4 {
		node, prev := sfn.parent, sfn
		for node != nil && node.left == prev {
			prev = node
			node = node.parent
		}
		if node == nil {
			node, prev = prev, prev.left
		}
		if node.left != prev {
			if node.left == nil {
				node.leftNum += sfn.leftNum
			} else {
				node.left.AddToRightmost(sfn.leftNum)
			}
		}

		node, prev = sfn.parent, sfn
		for node != nil && node.right == prev {
			prev = node
			node = node.parent
		}
		if node == nil {
			node, prev = prev, prev.right
		}
		if node.right != prev {
			if node.right == nil {
				node.rightNum += sfn.rightNum
			} else {
				node.right.AddToLeftmost(sfn.rightNum)
			}
		}

		if sfn.parent != nil {
			if sfn.parent.left == sfn {
				sfn.parent.leftNum = 0
				sfn.parent.left = nil
			} else {
				sfn.parent.rightNum = 0
				sfn.parent.right = nil
			}
		}

		return true
	}
	if sfn.left != nil && sfn.left.Explode(depth+1) {
		return true
	}
	if sfn.right != nil && sfn.right.Explode(depth+1) {
		return true
	}
	return false
}

func (sfn *SnailfishNumber) AddToRightmost(val int) bool {
	if sfn.right == nil {
		sfn.rightNum += val
		return true
	}
	return sfn.right.AddToRightmost(val) || sfn.left.AddToRightmost(val)
}

func (sfn *SnailfishNumber) AddToLeftmost(val int) bool {
	if sfn.left == nil {
		sfn.leftNum += val
		return true
	}
	return sfn.left.AddToLeftmost(val) || sfn.right.AddToLeftmost(val)
}

func (sfn *SnailfishNumber) Split() (split bool) {
	if sfn.left != nil && sfn.left.Split() {
		return true
	}
	if sfn.leftNum > 9 {
		newPair := SnailfishNumber{parent: sfn, leftNum: sfn.leftNum / 2, rightNum: (sfn.leftNum + 1) / 2}
		sfn.leftNum = 0
		sfn.left = &newPair
		return true
	}
	if sfn.rightNum > 9 {
		newPair := SnailfishNumber{parent: sfn, leftNum: sfn.rightNum / 2, rightNum: (sfn.rightNum + 1) / 2}
		sfn.rightNum = 0
		sfn.right = &newPair
		return true
	}
	if sfn.right != nil && sfn.right.Split() {
		return true
	}
	return false
}

func (sfn *SnailfishNumber) Magnitude() int {
	var leftMag int
	if sfn.left == nil {
		leftMag = sfn.leftNum
	} else {
		leftMag = sfn.left.Magnitude()
	}
	var rightMag int
	if sfn.right == nil {
		rightMag = sfn.rightNum
	} else {
		rightMag = sfn.right.Magnitude()
	}
	return 3*leftMag + 2*rightMag
}
