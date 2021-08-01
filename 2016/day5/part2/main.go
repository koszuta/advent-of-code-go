package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

/*
 *   --- Day 5: How About a Nice Game of Chess? ---
 *                  --- Part One ---
 *
 *   https://adventofcode.com/2016/day/5
 */

const doorID = "cxdnnyjw"

func main() {
	password := make([]byte, 8)
	seenPositions := make(map[int]struct{})

	for i := 0; len(seenPositions) < 8; i++ {
		checksum := md5.Sum([]byte(fmt.Sprintf("%s%d", doorID, i)))
		if checksum[0] == 0 && checksum[1] == 0 && checksum[2] < 16 {
			sum := hex.EncodeToString(checksum[:])
			pos, err := strconv.Atoi(string(sum[5]))
			if err == nil && pos < 8 {
				_, seen := seenPositions[pos]
				if !seen {
					log.Println("i:", i, "checksum:", sum, "position:", string(sum[5]), "character:", string(sum[6]))
					password[pos] = sum[6]
					seenPositions[pos] = struct{}{}
				}
			}
		}
	}

	log.Printf("the password is %s\n", string(password))
}
