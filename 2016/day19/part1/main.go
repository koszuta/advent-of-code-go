package main

import "log"

/*
 *   --- Day 19: An Elephant Named Joseph ---
 *               --- Part One ---
 *
 *   https://adventofcode.com/2016/day/19
 */

const nElves = 3017957

func main() {
	elfGifts := make([]int, nElves)
	for i := range elfGifts {
		elfGifts[i] = 1
	}
	for currElf, nElvesWithGifts := 0, nElves; nElvesWithGifts > 1; currElf = (currElf + 1) % nElves {
		if elfGifts[currElf] != 0 {
			var nextElf int
			for nextElf = (currElf + 1) % nElves; elfGifts[nextElf] == 0; nextElf = (nextElf + 1) % nElves {
			}
			elfGifts[currElf] += elfGifts[nextElf]
			elfGifts[nextElf] = 0
			nElvesWithGifts--
		}
	}
	for elf, nGifts := range elfGifts {
		if nGifts != 0 {
			log.Printf("Elf %d gets all the presents\n", elf+1)
			break
		}
	}
}
