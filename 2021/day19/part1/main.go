package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const expectedResult = 0

/*
 *   --- Day 19: Beacon Scanner ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2021/day/19
 */

var expectedBeacons = []Beacon{
	{-892, 524, 684},
	{-876, 649, 763},
	{-838, 591, 734},
	{-789, 900, -551},
	{-739, -1745, 668},
	{-706, -3180, -659},
	{-697, -3072, -689},
	{-689, 845, -530},
	{-687, -1600, 576},
	{-661, -816, -575},
	{-654, -3158, -753},
	{-635, -1737, 486},
	{-631, -672, 1502},
	{-624, -1620, 1868},
	{-620, -3212, 371},
	{-618, -824, -621},
	{-612, -1695, 1788},
	{-601, -1648, -643},
	{-584, 868, -557},
	{-537, -823, -458},
	{-532, -1715, 1894},
	{-518, -1681, -600},
	{-499, -1607, -770},
	{-485, -357, 347},
	{-470, -3283, 303},
	{-456, -621, 1527},
	{-447, -329, 318},
	{-430, -3130, 366},
	{-413, -627, 1469},
	{-345, -311, 381},
	{-36, -1284, 1171},
	{-27, -1108, -65},
	{7, -33, -71},
	{12, -2351, -103},
	{26, -1119, 1091},
	{346, -2985, 342},
	{366, -3059, 397},
	{377, -2827, 367},
	{390, -675, -793},
	{396, -1931, -563},
	{404, -588, -901},
	{408, -1815, 803},
	{423, -701, 434},
	{432, -2009, 850},
	{443, 580, 662},
	{455, 729, 728},
	{456, -540, 1869},
	{459, -707, 401},
	{465, -695, 1988},
	{474, 580, 667},
	{496, -1584, 1900},
	{497, -1838, -617},
	{527, -524, 1933},
	{528, -643, 409},
	{534, -1912, 768},
	{544, -627, -890},
	{553, 345, -567},
	{564, 392, -477},
	{568, -2007, -577},
	{605, -1665, 1952},
	{612, -1593, 1893},
	{630, 319, -379},
	{686, -3108, -505},
	{776, -3184, -501},
	{846, -3110, -434},
	{1135, -1161, 1235},
	{1243, -1093, 1063},
	{1660, -552, 429},
	{1693, -557, 386},
	{1735, -437, 1738},
	{1749, -1800, 1813},
	{1772, -405, 1572},
	{1776, -675, 371},
	{1779, -442, 1789},
	{1780, -1548, 337},
	{1786, -1538, 337},
	{1847, -1591, 415},
	{1889, -1729, 1762},
	{1994, -1805, 1792},
}

const nRequiredMatches = 12

type Coord3D [3]int
type Beacon Coord3D
type Scanner []Beacon

var (
	beacons                        map[Beacon]struct{}
	directionVectorToBeaconPairMap map[Coord3D][2]Beacon
	beaconToDirectionVectorMap     map[[2]Beacon]Coord3D
)

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	x := doPart1()
	log.Println(x)
}

func doPart1() int {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scan := bufio.NewScanner(file)

	scanners := make([]Scanner, 0)
	var scanner Scanner
	headerRegex := regexp.MustCompile(`--- scanner \d+ ---`)
	for scan.Scan() {
		line := scan.Text()
		if headerRegex.Match([]byte(line)) {
			scanner = make(Scanner, 0)
		} else if line == "" {
			scanners = append(scanners, scanner)
		} else {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			scanner = append(scanner, [3]int{x, y, z})
		}
	}
	scanners = append(scanners, scanner)

	beacons = make(map[Beacon]struct{})
	directionVectorToBeaconPairMap = make(map[Coord3D][2]Beacon)
	beaconToDirectionVectorMap = make(map[[2]Beacon]Coord3D)

	// Add all beacons from the first scanner as a basis
	for _, b := range scanners[0] {
		addBeacon(b)
	}
	scanners = scanners[1:]
	nScanners++

	for len(scanners) > 0 {
	SCANNING:
		for i := 0; i < len(scanners); {
			s := scanners[i]
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if y == x {
						continue
					}
					for z := 0; z < 3; z++ {
						if z == x || z == y {
							continue
						}
						for _, xSign := range []int{1, -1} {
							for _, ySign := range []int{1, -1} {
								for _, zSign := range []int{1, -1} {
									var changeOfBasis, prevCOB Coord3D
									nMatches := 0
								OUTER:
									for i, b := range s {
										if len(s)-i < nRequiredMatches-nMatches {
											break OUTER // not enough beacons left to match
										}
										this := Beacon{b[x] * xSign, b[y] * ySign, b[z] * zSign}
										for _, B := range s {
											that := Beacon{B[x] * xSign, B[y] * ySign, B[z] * zSign}
											if this == that {
												continue
											}
											direction := this.to(that)
											if beaconPair, found := directionVectorToBeaconPairMap[direction]; found {
												nMatches++
												{
													var cob, prev Coord3D
													if newBeaconPair, found := directionVectorToBeaconPairMap[direction]; found {
														cob = newBeaconPair[0].to(newBeaconPair[0])
														if cob != prev {
															log.Panicln("cob:", cob, "prev:", prev)
														}
														prev = cob
													}
													if changeOfBasis == [3]int{0, 0, 0} {
														changeOfBasis = cob
													} else {
														if prevCOB != changeOfBasis {
															log.Println("change of basis:", changeOfBasis, "prevCOB:", prevCOB)
														}
													}
												}
												if nMatches >= nRequiredMatches {
													break OUTER
												}
												continue OUTER
											}
										}
									}
									if nMatches >= nRequiredMatches {
										for _, b := range s {
											addBeacon(Beacon{b[x] * xSign, b[y] * ySign, b[z] * zSign})
										}
										scanners = append(scanners[:i], scanners[i+1:]...)
										nScanners++
										continue SCANNING
									}
								}
							}
						}
					}
				}
			}
			log.Panicln("unable to find orientation")
			i++
		}
	}

	return len(beacons)
}

var nScanners = 0

func addBeacon(b Beacon) {
	if _, found := beacons[b]; !found {
		for B := range beacons {
			for _, beaconPair := range [][2]Beacon{{b, B}, {B, b}} {
				d := beaconPair[0].to(beaconPair[1])
				directionVectorToBeaconPairMap[d] = beaconPair
				beaconToDirectionVectorMap[beaconPair] = d
			}
		}
		beacons[b] = struct{}{}
		fmt.Println("added beacon", b.toString(), "from scanner", nScanners)
	}
}

func (b1 *Beacon) to(b2 Beacon) Coord3D {
	return Coord3D{b2[0] - b1[0], b2[1] - b1[1], b2[2] - b1[2]}
}

func (b *Beacon) toString() string {
	return fmt.Sprintf("%d,%d,%d", b[0], b[1], b[2])
}
