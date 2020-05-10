package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var medicineMolecule string
var moleculeMap map[string]*[]string
var maxMoleculeLength int = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	moleculeMap = make(map[string]*[]string)
	{
		// Parse input file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " => ")
			if len(parts) == 2 {
				// Lines with => are replacement molecules
				molecule := parts[0]
				replacement := parts[1]
				var replacements []string
				r, ok := moleculeMap[molecule]
				if !ok {
					replacements = make([]string, 0)
				} else {
					replacements = *r
				}
				replacements = append(replacements, replacement)
				moleculeMap[molecule] = &replacements
				if len(molecule) > maxMoleculeLength {
					maxMoleculeLength = len(molecule)
				}

			} else if line != "" {
				// Non-empty line without => is our medicine molecule
				medicineMolecule = line
			}
		}
	}

	for k, v := range moleculeMap {
		fmt.Printf("\"%s\"->%v\n", k, *v)
	}
	fmt.Printf("\nmaxMoleculeLength=%d\n", maxMoleculeLength)
	fmt.Printf("\nmedicineMolecule=%s\n", medicineMolecule)
	fmt.Println()

	start := time.Now()
	newMolecules := make(map[string]bool, 0)
	{
		for i := 0; i < len(medicineMolecule); i++ {
			for j := i + 1; j <= i+maxMoleculeLength && j <= len(medicineMolecule); j++ {
				m := medicineMolecule[i:j]
				replacements, ok := moleculeMap[m]
				if ok {
					for _, replacement := range *replacements {
						newMolecule := make([]byte, 0)
						newMolecule = append(newMolecule, medicineMolecule[:i]...)
						newMolecule = append(newMolecule, replacement...)
						newMolecule = append(newMolecule, medicineMolecule[j:]...)

						newMolecules[string(newMolecule)] = true
					}
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %v\n", elapsed)
	fmt.Printf("len(newMolecules)=%d\n\n", len(newMolecules))

	start = time.Now()
	steps := moleculeReplacement("e", 0, 0)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %v\n", elapsed)
	fmt.Printf("steps=%d\n", steps)
}

func moleculeReplacement(partialMolecule string, start, steps int) int {
	// fmt.Printf("partialMolecule=%s start=%d steps=%d\n", partialMolecule, start, steps)
	if partialMolecule == medicineMolecule {
		return steps
	}
	// fmt.Printf("len(partialMolecule)=%d len(medicineMolecule)=%d\n", len(partialMolecule), len(medicineMolecule))
	for i := start; i < len(partialMolecule); i++ {
		for j := i + 1; j <= i+maxMoleculeLength && j <= len(partialMolecule); j++ {
			m := partialMolecule[i:j]
			// fmt.Printf("%s[%d:%d]=%s\n", partialMolecule, i, j, m)
			r, found := moleculeMap[m]
			if found {
				// fmt.Printf("replacements=%v\n", *r)
				for _, replacement := range *r {
					if (i-1)+(len(partialMolecule)-j)+len(replacement) < len(medicineMolecule) {
						partial := make([]byte, 0)
						partial = append(partial, partialMolecule[:i]...)
						partial = append(partial, replacement...)
						partial = append(partial, partialMolecule[j:]...)
						s := moleculeReplacement(string(partial), i, steps+1)
						if s != -1 {
							return s
						}
					}
				}
			}
		}
	}
	return -1
}
