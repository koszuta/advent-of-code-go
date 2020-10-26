package main

import (
	"bufio"
	"fmt"
	"os"
)

const width = 25
const height = 6
const stride = width * height

func main() {
	fmt.Printf("Go...\n")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	input := []byte(scanner.Text())
	if len(input)%stride != 0 {
		fmt.Printf("Invalid input data: %d extra bytes\n", len(input)%stride)
		os.Exit(1)
	}

	for i, digit := range input {
		input[i] = digit - '0'
	}

	var onesByTwos int
	minZeros := -1
	for i := 0; i < len(input); i += stride {
		digitCounts := make([]int, 3)
		for j := i; j < i+stride; j++ {
			digitCounts[input[j]]++
		}
		if minZeros == -1 || digitCounts[0] < minZeros {
			minZeros = digitCounts[0]
			onesByTwos = digitCounts[1] * digitCounts[2]
		}
	}
	fmt.Printf("\nonesByTwos=%d\n\n", onesByTwos)

	for j := 0; j < height; j++ {
	PRINTING:
		for i := 0; i < width; i++ {
			for k := i + width*j; k < len(input); k += stride {
				switch input[k] {
				case 0:
					fmt.Printf("  ")
					continue PRINTING
				case 1:
					fmt.Printf("██")
					continue PRINTING
				}
			}
			fmt.Printf("░░")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
