package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	name           string
	primary        *object
	indirectOrbits int
	satellites     []*object
}

func main() {
	fmt.Printf("Go...\n")

	COM := object{
		"COM",
		nil,
		0,
		make([]*object, 0),
	}

	indirectOrbits := parseSatellites(&COM)
	fmt.Printf("indirectOrbits=%d\n", indirectOrbits)

	YOU, foundYOU := find("YOU", COM)
	SAN, foundSAN := find("SAN", COM)
	fmt.Printf("found YOU and SAN=%v\n", foundYOU && foundSAN)

	primaryOfYOU := YOU.primary
OUT:
	for primaryOfYOU != nil {
		primaryOfSAN := SAN.primary
		for primaryOfSAN != nil {
			if primaryOfYOU == primaryOfSAN {
				break OUT
			}
			primaryOfSAN = primaryOfSAN.primary
		}
		primaryOfYOU = primaryOfYOU.primary
	}
	fmt.Printf("commonPrimary=%v\n", primaryOfYOU)

	orbitalTransfers := (YOU.indirectOrbits - primaryOfYOU.indirectOrbits - 1) + (SAN.indirectOrbits - primaryOfYOU.indirectOrbits - 1)
	fmt.Printf("orbitalTransfers=%d\n", orbitalTransfers)
}

func find(name string, obj object) (object, bool) {
	if obj.name == name {
		return obj, true
	}
	for _, satellite := range obj.satellites {
		sat, found := find(name, *satellite)
		if found {
			return sat, found
		}
	}
	return object{}, false
}

func parseSatellites(obj *object) int {
	indirectOrbits := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ")")
		if names[0] == obj.name {
			// fmt.Printf("%s)))%s\n", names[0], names[1])
			newSatellite := object{
				names[1],
				obj,
				obj.indirectOrbits + 1,
				make([]*object, 0),
			}
			obj.satellites = append(obj.satellites, &newSatellite)
			indirectOrbits += newSatellite.indirectOrbits
		}
	}
	for _, satellite := range obj.satellites {
		indirectOrbits += parseSatellites(satellite)
	}

	return indirectOrbits
}
