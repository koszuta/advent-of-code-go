package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseDimensions(line string) (int, int, int, error) {
	var l, w, h int
	var err error
	dimensions := strings.Split(line, "x")
	l, err = strconv.Atoi(dimensions[0])
	if err != nil {
		return 0, 0, 0, err
	}
	w, err = strconv.Atoi(dimensions[1])
	if err != nil {
		return 0, 0, 0, err
	}
	h, err = strconv.Atoi(dimensions[2])
	if err != nil {
		return 0, 0, 0, err
	}
	return l, w, h, err
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	totalPaper := 0
	totalRibbon := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l, w, h, err := parseDimensions(line)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		// fmt.Printf("%dx%dx%d\n", l, w, h)
		side1 := l * w
		side2 := l * h
		side3 := w * h
		shortSide := int(math.Min(float64(side1), math.Min(float64(side2), float64(side3))))
		totalPaper += (2 * side1) + (2 * side2) + (2 * side3) + shortSide

		longEdge := int(math.Max(float64(l), math.Max(float64(w), float64(h))))
		vol := l * w * h
		totalRibbon += (2 * l) + (2 * w) + (2 * h) - (2 * longEdge) + vol
	}

	fmt.Printf("totalPaper=%d\n", totalPaper)
	fmt.Printf("totalRibbon=%d\n", totalRibbon)
}
