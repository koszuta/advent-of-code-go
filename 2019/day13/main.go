package main

import (
	"Advent-Of-Code-Go/2019/day13/game"
	"flag"
	"log"
)

func main() {
	fileName := flag.String("file", "./input.txt", "path to input file")
	flag.Parse()
	log.Printf("fileName=%s\n", *fileName)

	breakout := game.New(*fileName)
	score := breakout.Play()
	log.Printf("You scored %d points\n", score)

	// Part 1
	// blocks := 0
	// for _, tile := range gameScreen {
	// 	if tile == screen.Block {
	// 		blocks++
	// 	}
	// }
	// log.Println()
	// log.Printf("%d block tiles on the screen\n", blocks)
}
