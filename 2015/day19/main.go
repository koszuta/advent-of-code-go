package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var medicine string
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
				if len(replacement) < len(molecule) {
					panic("Replacement cannot be shorter than original molecule")
				}
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
				medicine = line
			}
		}
	}

	for k, v := range moleculeMap {
		fmt.Printf("\"%s\"->%v\n", k, *v)
	}
	fmt.Printf("\nmaxMoleculeLength=%d\n", maxMoleculeLength)
	fmt.Printf("\nmedicine=%s\n", medicine)
	fmt.Println()

	start := time.Now()
	newMolecules := make(map[string]bool, 0)
	{
		for i := 0; i < len(medicine); i++ {
			for j := i + 1; j <= i+maxMoleculeLength && j <= len(medicine); j++ {
				m := medicine[i:j]
				replacements, ok := moleculeMap[m]
				if ok {
					for _, replacement := range *replacements {
						newMolecule := make([]byte, 0)
						newMolecule = append(newMolecule, medicine[:i]...)
						newMolecule = append(newMolecule, replacement...)
						newMolecule = append(newMolecule, medicine[j:]...)

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

	var buildMedicine func(int)
	chars := make([]byte, len(medicine))
	chars[0] = 'e'
	size := 1
	steps, minSteps := 0, -1
	buildMedicine = func(start int) {
		if size == len(medicine) && string(chars) == medicine {
			// fmt.Printf("steps=%d\n", steps)
			if steps < minSteps || minSteps == -1 {
				minSteps = steps
				fmt.Printf("new minSteps=%d\n", minSteps)
			}
			return
		}
		for i := start; (i == 0 || chars[i-1] == medicine[i-1]) && i < size; i++ {
			for j := 0; j < maxMoleculeLength && i+j < size; j++ {
				m := string(chars[i : i+j+1])
				replacements, found := moleculeMap[m]
				if found {
					for _, r := range *replacements {
						replExtra := len(r) - len(m)
						steps++
						size += replExtra
						if size <= len(medicine) {
							// Shift over chars to make room for repl molecule
							if replExtra != 0 {
								copy(chars[i+len(r):], chars[i+len(m):])
							}
							// Copy in repl molecule chars
							copy(chars[i:], r)

							buildMedicine(i)

							// Copy in original molecule chars
							copy(chars[i:], m)
							// Move rest of chars back
							if replExtra != 0 {
								copy(chars[i+len(m):], chars[i+len(r):])
							}
						}
						size -= replExtra
						steps--
					}
				}
			}
		}
	}
	buildMedicine(0)

	fmt.Printf("Part 2 took %v\n", time.Since(start))
	fmt.Printf("minSteps=%d\n", minSteps)
}
