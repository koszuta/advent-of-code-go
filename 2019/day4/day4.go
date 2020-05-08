package main

import (
	"fmt"
	"strconv"
	"strings"
)

const maxLength = 6

// 172930-683082
var minPassword = 177777
var maxPassword = 679999

func main() {
	fmt.Printf("Go...\n")

	numValidPasswords := 0
OUT:
	for i := 0; i <= maxPassword-minPassword; i++ {
		password := strconv.Itoa(minPassword + i)

		for j := 1; j < maxLength; j++ {
			// fmt.Printf("%c < %c = %v\n", password[j-1], password[j], password[j-1] > password[j])
			if password[j-1] > password[j] {
				continue OUT
			}
		}

		for j := 1; j <= 9; j++ {
			if strings.Contains(password, strconv.Itoa(10*j+j)) && !strings.Contains(password, strconv.Itoa(100*j+10*j+j)) {
				// fmt.Printf("password=%s\n", password)
				numValidPasswords++
				break
			}
		}
	}
	fmt.Printf("numValidPasswords=%d\n", numValidPasswords)
}
