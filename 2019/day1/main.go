package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Go...\n")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	fuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		for mass > 8 {
			mass /= 3
			mass -= 2
			fuel += mass
		}
	}

	fmt.Printf("fuel=%d\n", fuel)
}
