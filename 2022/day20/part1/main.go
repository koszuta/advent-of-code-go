package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	val        int
	prev, next *Node
}

var lines []string

func init() {
	b, _ := os.ReadFile("../input.txt")
	lines = strings.Split(string(b), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line) // sanitize CRLF
	}
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-2]
	}
}

func main() {
	defer func(t time.Time) {
		log.Println("took:", time.Since(t))
	}(time.Now())

	nodes := make([]*Node, 0, len(lines))
	var head, prev, tail *Node
	for _, line := range lines {
		v, _ := strconv.Atoi(line)

		node := Node{
			val:  v,
			prev: prev,
		}
		if prev != nil {
			prev.next = &node
		}
		prev = &node

		if head == nil {
			head = &node
		}
		tail = &node

		nodes = append(nodes, &node)
	}
	head.prev = tail
	tail.next = head

	fmt.Println(head.String())

	for _, this := range nodes {
		target := this
		for j := 0; j < Abs(this.val); j++ {
			if this.val > 0 {
				target = target.next
			} else {
				target = target.prev
			}
		}
		this.prev.next, this.next.prev = this.next, this.prev // remove node from old position
		if this.val > 0 {
			this.next, this.prev, target.next, target.next.prev = target.next, target, this, this
		} else {
			this.next, this.prev, target.prev, target.prev.next = target, target.prev, this, this
		}
		fmt.Println(head.String())
	}
	fmt.Println(head.String())
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (n *Node) String() string {
	result := strconv.Itoa(n.val)
	this := n.next
	for this != nil && this != n {
		result = fmt.Sprintf("%s, %d", result, this.val)
		this = this.next
	}
	if this == n {
		return fmt.Sprintf("%s, (%d)", result, n.val)
	}
	return result
}
