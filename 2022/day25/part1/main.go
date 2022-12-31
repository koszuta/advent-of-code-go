package main

import (
	"log"
	"math"
	"os"
	"strings"
)

var (
	digits = map[rune]int{
		'2': 2,
		'1': 1,
		'0': 0,
		'-': -1,
		'=': -2,
	}
)

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
	var sum int
	for _, line := range lines {
		n := SNAFUToBase10(line)
		log.Println(line, " -> ", n)
		sum += n
	}
	log.Println("SNAFU number:", sum)
}

func SNAFUToBase10(n string) (b10 int) {
	p := int(math.Pow(5., float64(len(n)-1)))
	for _, d := range n {
		if v, found := digits[d]; found {
			b10 += p * v
		}
		p /= 5
	}
	return b10
}

// func Base10ToSNAFU(n int) (snafu string) {
// 	p := int64(math.Pow(5., float64(len(n)-1)))
// 	for _, d := range n {
// 		if v, found := digits[d]; found {
// 			n10.Add(&n10, big.NewInt(p*v))
// 		}
// 		p /= 5
// 	}
// 	return
// }
