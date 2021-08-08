package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

/*
 *   --- Day 17: Two Steps Forward ---
 *           --- Part One ---
 *
 *   https://adventofcode.com/2016/day/17
 */

const (
	passcode = "mmsxrhfx"
	destX    = 3
	destY    = 3
)

func main() {
	paths := make([]string, 0)
	paths = move(0, 0, passcode, paths)
	if len(paths) == 0 {
		log.Panicln("couldn't find a path to the vault")
	}

	minPath := paths[0]
	for _, path := range paths {
		if len(path) < len(minPath) {
			minPath = path
		}
	}
	log.Printf("the shortest path to reach the vault is %s\n", minPath)
}

func move(x, y int, code string, paths []string) []string {
	if x == destX && y == destY {
		return append(paths, code[len(passcode):])
	}
	if x < 0 || x > destX || y < 0 || y > destY {
		return paths
	}
	checksum := md5.Sum([]byte(code))
	hash := hex.EncodeToString(checksum[:])
	// Up
	if hash[0] > 'a' {
		paths = move(x, y-1, code+"U", paths)
	}
	// Down
	if hash[1] > 'a' {
		paths = move(x, y+1, code+"D", paths)
	}
	// Left
	if hash[2] > 'a' {
		paths = move(x-1, y, code+"L", paths)
	}
	// Right
	if hash[3] > 'a' {
		paths = move(x+1, y, code+"R", paths)
	}
	return paths
}
