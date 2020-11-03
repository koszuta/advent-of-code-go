package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var chemEquations map[string]equation

type chemical struct {
	coefficient int
	name        string
}

type equation struct {
	inputs []chemical
	output chemical
}

func (e *equation) canMake(chems map[string]int) (made bool) {
	for _, input := range e.inputs {
		if chems[input.name] < input.coefficient {
			return false
		}
	}
	return true
}

func (e *equation) make(chems map[string]int) (made bool) {
	for _, input := range e.inputs {
		chems[input.name] -= input.coefficient
	}
	chems[e.output.name] += e.output.coefficient
	return true
}

func (e *equation) unmake(chems map[string]int) {
	chems[e.output.name] -= e.output.coefficient
	for _, input := range e.inputs {
		chems[input.name] += input.coefficient
	}
}

func (e *equation) unmakeN(n int, chems map[string]int) {
	chems[e.output.name] -= n * e.output.coefficient
	for _, input := range e.inputs {
		chems[input.name] += n * input.coefficient
	}
}

func parseChemical(parts []string) chemical {
	coefficient, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Panicln(err)
	}
	return chemical{coefficient: coefficient, name: parts[1]}
}

func parseEquation(line string) {
	parts := strings.Split(line, " => ")

	inputs := make([]chemical, 0, 0)
	for _, input := range strings.Split(parts[0], ", ") {
		chem := parseChemical(strings.Split(input, " "))
		inputs = append(inputs, chem)
	}
	output := parseChemical(strings.Split(parts[1], " "))
	e := equation{inputs: inputs, output: output}

	chemEquations[output.name] = e
}

func init() {
	fileName := flag.String("file", "input.txt", "name of the input file")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		log.Panicln(err)
	}

	chemEquations = make(map[string]equation)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseEquation(scanner.Text())
	}

	// log.Printf("chemical equations %v\n", chemEquations)
}

func areChemsLeft(chems map[string]int) bool {
	for name, coefficient := range chems {
		if name != "ORE" && coefficient > 0 {
			return true
		}
	}
	return false
}

func oresForNFuel(fuels int) int {
	chems := make(map[string]int)
	chems["FUEL"] = fuels
	for areChemsLeft(chems) {
		for name, coefficient := range chems {
			if name != "ORE" {
				e := chemEquations[name]
				reactions := int(math.Ceil(float64(coefficient) / float64(e.output.coefficient)))
				e.unmakeN(reactions, chems)
			}
		}
	}
	return chems["ORE"]
}

func main() {
	start := time.Now()

	oresForOneFuel := oresForNFuel(1)
	log.Printf("%d ORE is required to make 1 FUEL\n", oresForOneFuel)
	log.Printf("took %v\n", time.Since(start))

	ores := 1000000000000
	lo, hi := 1, ores
	for lo < hi {
		mid := (lo + hi) / 2
		if oresForNFuel(mid) > ores {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	log.Printf("%d ORE is required to make %d FUEL\n", oresForNFuel(lo-1), lo-1)
	log.Printf("%d ORE is required to make %d FUEL\n", oresForNFuel(lo), lo)
	log.Printf("took %v\n", time.Since(start))
}
