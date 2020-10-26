package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
)

const (
	secretKeyFlag string = "secretKey"
	zerosFlag     string = "zeros"
)

var (
	secretKey string
	zeros     int
)

func init() {
	fmt.Printf("Initializing...\n")

	flag.IntVar(&zeros, zerosFlag, 5, "Number of leading zeros required in hash")
	flag.StringVar(&secretKey, secretKeyFlag, "bgvyzdsv", "Secret key string")
	flag.Parse()
	fmt.Printf("%s=%d\n", zerosFlag, zeros)
	fmt.Printf("%s=%s\n", secretKeyFlag, secretKey)
}

func main() {
	fmt.Printf("Go...\n")

OUTER:
	for i := 1; i > 0; i++ {
		input := []byte(secretKey + strconv.Itoa(i))
		hash := md5.New()
		hash.Write(input)
		sum := hash.Sum(nil)
		encoded := hex.EncodeToString(sum[:])
		if len(encoded) >= zeros {
			for j := 0; j < zeros; j++ {
				if encoded[j] != '0' {
					continue OUTER
				}
			}
			fmt.Printf("result=%d\n", i)
			break OUTER
		}
	}
}
