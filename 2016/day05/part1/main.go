package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

/*
 *   --- Day 5: How About a Nice Game of Chess? ---
 *                  --- Part One ---
 *
 *   https://adventofcode.com/2016/day/5
 */

const doorID = "cxdnnyjw"

func main() {
	password := ""

	for i := 0; len(password) < 8; i++ {
		checksum := md5.Sum([]byte(fmt.Sprintf("%s%d", doorID, i)))
		if checksum[0] == 0 && checksum[1] == 0 && checksum[2] < 16 {
			sum := hex.EncodeToString(checksum[:])
			log.Println("i:", i, "checksum:", sum, "character:", string(sum[5]))
			password += string(sum[5])
		}
	}

	log.Printf("the password is %s\n", password)
}
