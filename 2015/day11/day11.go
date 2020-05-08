package main

import (
	"fmt"
)

var pw string = "hxbxwxba"

func incrementPassword(pw string) string {
	bytes := []byte(pw)
	for i := len(bytes)-1; i >= 0; i-- {
		if bytes[i] == 'z' {
			bytes[i] = 'a'
		} else {
			bytes[i] = bytes[i]+1
			break
		}
	}
	return string(bytes)
}

func isPasswordValid(pw string) bool {
	containsIOL := false
	for i := 0; i < len(pw); i++ {
		if pw[i] == 'i' || pw[i] == 'o' || pw[i] == 'l' {
			containsIOL = true
			break
		}
	}
	if containsIOL {
		return false
	}

	containsRun := false
	for i := 2; i < len(pw); i++ {
		if pw[i-2]+1 == pw[i-1] && pw[i-1]+1 == pw[i] {
			containsRun = true
			break
		}
	}
	if !containsRun {
		return false
	}

	numPairs := 0
	for i := 1; i < len(pw); i++ {
		if pw[i-1] == pw[i] {
			numPairs++
			i++
		}
	}
	return numPairs > 1
}

func main() {
	fmt.Printf("pw=%s\n", pw)
	for !isPasswordValid(pw) {
		pw = incrementPassword(pw)
	}
	fmt.Printf("next pw=%s\n", pw)

	pw = incrementPassword(pw)
	for !isPasswordValid(pw) {
		pw = incrementPassword(pw)
	}
	fmt.Printf("next pw=%s\n", pw)
}
