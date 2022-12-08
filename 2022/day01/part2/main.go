package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(b), "\n")

	var sum int
	calorieTotals := make([]int, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line) // sanitize CRLF

		if line == "" {
			calorieTotals = append(calorieTotals, sum)
			sum = 0
		}

		calories, _ := strconv.Atoi(line)
		sum += calories
	}

	sort.Slice(calorieTotals, func(i, j int) bool {
		return calorieTotals[i] > calorieTotals[j]
	})

	fmt.Println("3rd most calories:", calorieTotals[2])
	fmt.Println("2nd most calories:", calorieTotals[1])
	fmt.Println("most calories:", calorieTotals[0])
	fmt.Println("top 3 calories sum:", calorieTotals[0]+calorieTotals[1]+calorieTotals[2])
}
