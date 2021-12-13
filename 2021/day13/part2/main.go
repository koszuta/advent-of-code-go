package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const expectedResult = `
███  ████ ████   ██ █  █ ███  ████ ████ 
█  █    █ █       █ █  █ █  █ █       █ 
█  █   █  ███     █ ████ █  █ ███    █  
███   █   █       █ █  █ ███  █     █   
█    █    █    █  █ █  █ █ █  █    █    
█    ████ █     ██  █  █ █  █ █    ████ `

/*
 *   --- Day 13: Transparent Origami ---
 *            --- Part Two ---
 *
 *   https://adventofcode.com/2021/day/13#part2
 */

var dots [][]bool

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	code := doPart2()
	fmt.Println(code)
}

func doPart2() string {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// I don't feel like finding the actual size beforehand; this is sufficient
	size := 1337
	dots = make([][]bool, size)
	for i := range dots {
		dots[i] = make([]bool, size)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		dots[y][x] = true
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "fold along x=") {
			xFold, _ := strconv.Atoi(strings.Split(line, "fold along x=")[1])
			dots = foldLeft(xFold)

		} else if strings.Contains(line, "fold along y=") {
			yFold, _ := strconv.Atoi(strings.Split(line, "fold along y=")[1])
			dots = foldUp(yFold)
		}
	}

	code := "\n"
	for j, row := range dots {
		for _, dot := range row {
			if dot {
				code += "\u2588"
			} else {
				code += " "
			}
		}
		if j < len(dots)-1 {
			code += "\n"
		}
	}
	return code
}

func foldLeft(x int) [][]bool {
	newDots := make([][]bool, len(dots))
	for i := range newDots {
		newDots[i] = make([]bool, x)
	}
	for k := 0; k < len(dots); k++ {
		for i, j := x+1, x-1; i < len(dots[0]) && j >= 0; i, j = i+1, j-1 {
			if dots[k][i] || dots[k][j] {
				newDots[k][j] = true
			}
		}
	}
	return newDots
}

func foldUp(y int) [][]bool {
	newDots := make([][]bool, y)
	for i := range newDots {
		newDots[i] = make([]bool, len(dots[0]))
	}
	for j, k := y+1, y-1; j < len(dots) && k >= 0; j, k = j+1, k-1 {
		for i := 0; i < len(dots[0]); i++ {
			if dots[j][i] || dots[k][i] {
				newDots[k][i] = true
			}
		}
	}
	return newDots
}
