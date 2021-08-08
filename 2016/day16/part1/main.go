package main

import "log"

/*
 *   --- Day 16: Dragon Checksum ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2016/day/16
 */

const (
	initialState = "10111011111001111"
	lenToFill    = 272
)

func main() {
	log.Println("initial state:", initialState)

	fillData := fill(initialState)
	for len(fillData) < lenToFill {
		fillData = fill(fillData)
	}
	if len(fillData) > lenToFill {
		fillData = fillData[:lenToFill]
	}
	log.Println("fill data:", fillData)

	checksum := sum(fillData)
	for len(checksum)%2 == 0 {
		checksum = sum(checksum)
	}
	log.Println("check sum:", checksum)
}

func sum(data string) string {
	checksum := make([]byte, len(data)/2)
	for i := 1; i < len(data); i += 2 {
		if data[i-1] == data[i] {
			checksum[i/2] = '1'
		} else {
			checksum[i/2] = '0'
		}
	}
	return string(checksum)
}

func fill(state string) string {
	chars := []byte(state)
	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-i-1] = chars[len(chars)-i-1], chars[i]
	}
	for i, c := range chars {
		switch c {
		case '0':
			chars[i] = '1'
		case '1':
			chars[i] = '0'
		default:
			log.Panicln("unknown character", c)
		}
	}
	return state + "0" + string(chars)
}
