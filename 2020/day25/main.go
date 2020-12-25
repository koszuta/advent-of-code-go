package main

import "log"

/*
 *   --- Day 25: Combo Breaker ---
 *
 *   https://adventofcode.com/2020/day/25
 */

const magicNum = 20201227

const cardPublicKey, doorPublicKey = 15113849, 4206373

// const cardPublicKey, doorPublicKey = 5764801, 17807724 // example

func getLoopSize(publicKey, subjectNo int) (loopSize int) {
	for val := 1; val != publicKey; loopSize++ {
		val = (val * subjectNo) % magicNum
	}
	return
}

func getEncryptionKey(loopSize, subjectNum int) (encryptionKey int) {
	encryptionKey = 1
	for i := 0; i < loopSize; i++ {
		encryptionKey = (encryptionKey * subjectNum) % magicNum
	}
	return
}

func main() {
	cardLoopSize := getLoopSize(cardPublicKey, 7)
	log.Println("card loop size:", cardLoopSize)

	doorLoopSize := getLoopSize(doorPublicKey, 7)
	log.Println("door loop size:", doorLoopSize)

	log.Println("encryption key:", getEncryptionKey(cardLoopSize, doorPublicKey), "(using card)")

	log.Println("encryption key:", getEncryptionKey(doorLoopSize, cardPublicKey), "(using door)")
}
