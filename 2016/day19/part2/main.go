package main

import (
	"log"
	"time"
)

/*
 *   --- Day 19: An Elephant Named Joseph ---
 *               --- Part One ---
 *
 *   https://adventofcode.com/2016/day/19
 */

const nElves = 3017957

type elf struct {
	n          int
	prev, next *elf
}

func main() {
	log.Println("go...")
	defer (func(start time.Time) {
		log.Println("...took", time.Since(start))
	})(time.Now())

	// Keep track of the current elf and the one across the circle
	var curr, across *elf

	// Init circular buffer of elves
	{
		elf1 := elf{n: 1}
		prev := &elf1
		for i := 1; i < nElves; i++ {
			curr := elf{n: i + 1, prev: prev}
			if i == nElves/2 {
				across = &curr
			}
			prev.next = &curr
			prev = &curr
		}
		prev.next = &elf1
		elf1.prev = prev
		curr = &elf1
	}

	for nElvesWithGifts := nElves; curr.next != curr; nElvesWithGifts-- {
		// Remove the elf across the circle
		across.next.prev, across.prev.next = across.prev, across.next
		// If the number of elves in the circle was even, the across elf pointer moves one to the left
		// If it's odd, the across elf pointer moves two to the left
		if nElvesWithGifts%2 == 0 {
			across = across.next
		} else {
			across = across.next.next
		}
		// Current elf pointer moves one to the left
		curr = curr.next
	}
	log.Printf("Elf %d gets all the presents\n", curr.n)
}
