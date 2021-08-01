package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
 *   --- Day 4: Security Through Obscurity ---
 *               --- Part One ---
 *
 *   https://adventofcode.com/2016/day/4
 */

var roomNameRex, sectorIdRex, checksumRex *regexp.Regexp

func init() {
	roomNameRex = regexp.MustCompile(`([a-z\-]+?)\-\d+\[[a-z]{5}\]`)
	sectorIdRex = regexp.MustCompile(`[a-z\-]+\-(\d+)\[[a-z]{5}\]`)
	checksumRex = regexp.MustCompile(`[a-z\-]+\-\d+\[([a-z]{5})\]`)
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	sumOfSectorIDs := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		roomName := roomNameRex.FindStringSubmatch(line)[1]
		sectorID, _ := strconv.Atoi(sectorIdRex.FindStringSubmatch(line)[1])
		checksum := checksumRex.FindStringSubmatch(line)[1]

		actualChecksum := computeCheckSum(roomName)
		if actualChecksum == checksum {
			sumOfSectorIDs += sectorID
		}
	}

	log.Printf("the sum of the sector IDs of the valid room names is %d\n", sumOfSectorIDs)
}

func computeCheckSum(encName string) string {
	charCounts := make(map[string]int)
	for _, c := range strings.ReplaceAll(encName, "-", "") {
		charCounts[string(c)]++
	}

	chars := make([]pair, 0, len(charCounts))
	for k, v := range charCounts {
		chars = append(chars, pair{k, v})
	}
	sort.Slice(chars, func(i, j int) bool {
		if chars[i].v == chars[j].v {
			return chars[i].k < chars[j].k
		}
		return chars[i].v > chars[j].v
	})

	if len(chars) < 5 {
		return ""
	}
	return chars[0].k + chars[1].k + chars[2].k + chars[3].k + chars[4].k
}

type pair struct {
	k string
	v int
}
