package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var molecule string
var moleculeMap map[string]*[]string
var newMolecules []string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	moleculeMap = make(map[string]*[]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		if len(parts) == 2 {
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

		} else if line != "" {
			molecule = line
		}
	}

	for k, v := range moleculeMap {
		fmt.Printf("\"%s\"->%v\n", k, *v)
	}
	fmt.Printf("molecule=%s\n\n", molecule)

	newMolecules := make([]string, 0)

	for i := 0; i < len(molecule); i++ {
		for m, r := range moleculeMap {
			j := i + len(m)
			if j <= len(molecule) && m == string(molecule[i:j]) {
				for _, replacement := range *r {
					newMolecule := make([]byte, 0)
					newMolecule = append(newMolecule, molecule[:i]...)
					newMolecule = append(newMolecule, replacement...)
					newMolecule = append(newMolecule, molecule[j:]...)
					
					contains := false
					for _, otherMolecule := range newMolecules {
						if string(newMolecule) == otherMolecule {
							contains = true
							break
						}
					}
					if !contains {
						newMolecules = append(newMolecules, string(newMolecule))
					}
				}
			}
		}
	}

	fmt.Printf("len(newMolecules)=%d\n", len(newMolecules))
}
