package day20

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type coord2D struct {
	x, y int
}

type pixels map[coord2D]struct{}

var (
	algorithm string
	image     pixels
)

func Do(iterations int) int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	algorithm = scanner.Text()

	scanner.Scan() // skip blank line

	image = make(map[coord2D]struct{})
	for j := 0; scanner.Scan(); j++ {
		for i, r := range scanner.Text() {
			if r == '#' {
				image[coord2D{i, j}] = struct{}{}
			}
		}
	}

	for i := 0; i < iterations; i++ {
		image = image.enhance(i%2 == 0)
	}
	return len(image)
}

func (img *pixels) enhance(on bool) pixels {
	outputPixels := make(map[coord2D]struct{})
	setOutputPixel := func(c coord2D) {
		if (algorithm[img.getIndex(c, on)] == '#') != on {
			outputPixels[c] = struct{}{}
		}
	}

	q := 1
	for c := range *img {
		for j := -q; j <= q; j++ {
			for i := -q; i <= q; i++ {
				setOutputPixel(coord2D{c.x + i, c.y + j})
			}
		}
	}
	return outputPixels
}

func (img *pixels) getIndex(c coord2D, on bool) int {
	getBit := func(c coord2D) byte {
		if _, found := (*img)[c]; found == on {
			return '1'
		}
		return '0'
	}

	binary := make([]byte, 9)
	var q int
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			binary[q] = getBit(coord2D{c.x + i, c.y + j})
			q++
		}
	}

	i, _ := strconv.ParseInt(string(binary), 2, 64)
	return int(i)
}

func (img *pixels) prettyPrint() {
	minX, maxX := math.MaxInt64, 0
	for c := range *img {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
	}
	minY, maxY := math.MaxInt64, 0
	for c := range *img {
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	bytes := make([][]byte, maxY-minY+1)
	for j := range bytes {
		bytes[j] = make([]byte, maxX-minX+1)
		for i := range bytes[j] {
			bytes[j][i] = ' '
		}
	}

	for c := range *img {
		bytes[c.y-minY][c.x-minX] = '#'
	}

	for _, row := range bytes {
		fmt.Println(string(row))
	}
	fmt.Println()
}
