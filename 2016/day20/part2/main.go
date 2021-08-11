package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 *   --- Day 20: Firewall Rules ---
 *          --- Part Two ---
 *
 *   https://adventofcode.com/2016/day/20#part2
 */

const (
	lowestIP  = 0
	highestIP = 4294967295
)

type uiRange struct {
	lo, hi uint
}

var (
	validIPs  []uiRange
	blacklist []uiRange
)

func init() {
	validIPs = append(validIPs, uiRange{lowestIP, highestIP})
	{
		file, err := os.Open("../input.txt")
		if err != nil {
			log.Panicln(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Panicln(err)
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Panicln(err)
			}
			if start > end {
				log.Panicf("invalid blacklist range; start=%d, end=%d\n", start, end)
			}
			blacklist = append(blacklist, uiRange{uint(start), uint(end)})
		}
		if len(blacklist) == 0 {
			log.Panicln("couldn't parse blacklisted IPs")
		}
	}
}

func main() {
	for _, ipRange := range blacklist {
	OUT:
		switch {
		case ipRange.hi < validIPs[0].lo || ipRange.lo > validIPs[len(validIPs)-1].hi:
			// If the blacklisted range is completely outside of the valid ranges, ignore it

		case ipRange.lo == validIPs[0].lo:
			// If the low end of the blacklisted range is equal to the lowest valid IP
			var next int
			for next = 0; next < len(validIPs)-1 && validIPs[next].hi <= ipRange.hi; next++ {
			}
			validIPs = validIPs[next:]
			if validIPs[0].lo <= ipRange.hi {
				validIPs[0].lo = ipRange.hi + 1
			}

		case ipRange.hi == validIPs[len(validIPs)-1].hi:
			// If the high end of the blacklisted range is equal to the highest valid IP
			var next int
			for next = len(validIPs); next >= 0 && validIPs[next-1].lo >= ipRange.lo; next-- {
			}
			validIPs = validIPs[:next]
			if validIPs[len(validIPs)-1].hi >= ipRange.lo {
				validIPs[len(validIPs)-1].hi = ipRange.lo - 1
			}

		default:
			// If the blacklisted range is completely contained in a single valid range,
			// split the valid range in two
			for i, validIP := range validIPs {
				if validIP.lo < ipRange.lo && validIP.hi > ipRange.hi {
					validIPs[i].hi = ipRange.lo - 1
					newRange := uiRange{lo: ipRange.hi + 1, hi: validIP.hi}
					if len(validIPs) < 2 || i == len(validIPs)-1 {
						validIPs = append(validIPs, newRange)

					} else {
						validIPs = append(validIPs[:i+2], validIPs[i+1:]...)
						validIPs[i+1] = newRange
					}
					break OUT
				}
			}

			// Otherwise, delete all valid ranges that fall within the blacklisted range
			// and then update the values in the valid ranges the blacklisted range overlaps
			left, right := 0, len(validIPs)
			for i, validIP := range validIPs {
				if ipRange.lo > validIP.lo {
					left = i
				}
				if ipRange.hi >= validIP.hi {
					right = i
				}
			}
			validIPs = append(validIPs[:left+1], validIPs[right+1:]...)
			if ipRange.lo <= validIPs[left].hi {
				validIPs[left].hi = ipRange.lo - 1
			}
			if left < len(validIPs)-1 && ipRange.hi >= validIPs[left+1].lo {
				validIPs[left+1].lo = ipRange.hi + 1
			}
		}
	}

	nValidIPs := uint(0)
	for _, ipRange := range validIPs {
		nValidIPs += ipRange.hi - ipRange.lo + 1
	}
	log.Printf("the number of IPs allowed by the blacklist is %d\n", nValidIPs)
}
