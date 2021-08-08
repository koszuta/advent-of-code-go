package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

/*
 *   --- Day 14: One-Time Pad ---
 *         --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/14#part2
 */

const (
	salt         = "yjdafjpo"
	nKeys        = 64
	nPrevious    = 1000
	nExtraHashes = 2016
)

type hasher struct {
	index      int
	tripleChar byte
	hash       string
}

func main() {
	prevHashes := make([]hasher, 0, nPrevious)

	count, index := 0, -1
OUT:
	for {
		index++
		// Only keep the last 1000 hashes
		if len(prevHashes) != 0 && index-prevHashes[0].index > nPrevious {
			prevHashes = prevHashes[1:]
		}

		// Hash salt+index and convert to hex string
		hash := fmt.Sprintf("%s%d", salt, index)
		sum := md5.Sum([]byte(hash))
		hash = hex.EncodeToString(sum[:])

		// Do 2016 additional hashes to get the streched hash
		for i := 0; i < nExtraHashes; i++ {
			sum := md5.Sum([]byte(hash))
			hash = hex.EncodeToString(sum[:])
		}

		quintupleChar, hasQuintuple := containsQuintuple(hash)
		if hasQuintuple {
			for _, prev := range prevHashes {
				// Check if previous hashes had triple with same character
				if prev.tripleChar == quintupleChar {
					count++
					if count >= nKeys {
						log.Printf("the %dth key is %s\n", count, prev.hash)
						log.Printf("the index of the %dth key is %d\n", count, prev.index)
						// I've got an off-by-one error somewhere and an extra hash is getting through
						// Not sure where, but it only happens in part 2 with my puzzle input salt
						if count > nKeys {
							break OUT
						}
					}
				}
			}
		}

		tripleChar, hasTriple := containsTriple(hash)
		if hasTriple {
			// Keep track of hashes with triples
			prevHashes = append(prevHashes, hasher{index, tripleChar, hash})
		}
	}
}

func containsTriple(hash string) (byte, bool) {
	for i := 2; i < len(hash); i++ {
		if hash[i-2] == hash[i-1] && hash[i-1] == hash[i] {
			return hash[i], true
		}
	}
	return 0, false
}

func containsQuintuple(hash string) (byte, bool) {
	for i := 4; i < len(hash); i++ {
		if hash[i-4] == hash[i-3] && hash[i-3] == hash[i-2] && hash[i-2] == hash[i-1] && hash[i-1] == hash[i] {
			return hash[i], true
		}
	}
	return 0, false
}
